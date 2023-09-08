package kafkaclient

import (
	"context"
	"errors"
	"github.com/3lvia/libraries-go/pkg/mschema"
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

	registryKey, registrySecret, registryURL, err := getSecrets(ctx, system, collector.secrets)
	if err != nil {
		return nil, err
	}

	if err := mschema.Init(mschema.WithClient(c), mschema.WithUser(registryKey, registrySecret), mschema.WithSchemaRegistryURL(registryURL)); err != nil {
		return nil, err
	}

	schema, err := mschema.Get(topic)
	if err != nil {
		return nil, err
	}
	_ = schema

	ch := make(chan *StreamingMessage)

	return ch, nil
}
