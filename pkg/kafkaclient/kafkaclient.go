package kafkaclient

import (
	"context"
	"errors"
	"github.com/3lvia/libraries-go/pkg/mschema"
	"net/http"
)

func StartConsumer(ctx context.Context, system, topic, application string, opts ...Option) (<-chan *StreamingMessage, error) {
	collector := &optionsCollector{}
	for _, opt := range opts {
		opt(collector)
	}

	if collector.secrets == nil {
		return nil, errors.New("no secrets manager provided")
	}

	client := collector.client
	if client == nil {
		client = &http.Client{}
	}

	secrets, err := getSecrets(ctx, system, collector.secrets)
	if err != nil {
		return nil, err
	}

	creator := collector.creatorFunc
	if creator == nil {
		creator = func(data []byte, schemaID int) (any, error) {
			return data, nil
		}
	}

	registry, err := mschema.New(
		secrets.registryURL,
		mschema.WithClient(client),
		mschema.WithUser(secrets.registryKey, secrets.registrySecret))
	if err != nil {
		return nil, err
	}

	c, err := newConsumer(
		system,
		topic,
		application,
		secrets.bootstrapServer,
		secrets.key,
		secrets.secret,
		creator,
		registry)
	if err != nil {
		return nil, err
	}

	ch := make(chan *StreamingMessage)
	go c.start(ctx, ch)

	return ch, nil
}
