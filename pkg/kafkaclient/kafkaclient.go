package kafkaclient

import (
	"context"
	"errors"
	"github.com/3lvia/libraries-go/pkg/kafkaclient/schemaclient"
	"net/http"
)

func Init(ctx context.Context, system, topic string, opts ...Option) error {
	collector := &optionsCollector{}
	for _, opt := range opts {
		opt(collector)
	}

	if collector.secrets == nil {
		return errors.New("no secrets manager provided")
	}

	c := collector.client
	if c == nil {
		c = &http.Client{}
	}

	sClient, err := schemaclient.New(ctx, system, topic, schemaclient.WithSecretsManager(collector.secrets), schemaclient.WithHTTPClient(c))
	if err != nil {
		return err
	}
	_ = sClient

	return nil
}
