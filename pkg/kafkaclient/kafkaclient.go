package kafkaclient

import (
	"context"
	"errors"
	"github.com/3lvia/libraries-go/pkg/kafkaclient/schemaclient"
	"net/http"
)

func Start(ctx context.Context, system, topic string, opts ...Option) (<-chan *StreamingMessage, error) {
	collector := &optionsCollector{}
	for _, opt := range opts {
		opt(collector)
	}

	if collector.secrets == nil {
		return nil, errors.New("no secrets manager provided")
	}

	c := collector.client
	if c == nil {
		c = &http.Client{}
	}

	sClient, err := schemaclient.New(ctx, system, schemaclient.WithSecretsManager(collector.secrets), schemaclient.WithHTTPClient(c))
	if err != nil {
		return nil, err
	}

	schema, err := sClient.Get(topic)
	if err != nil {
		return nil, err
	}
	_ = schema

	ch := make(chan *StreamingMessage)

	return ch, nil
}
