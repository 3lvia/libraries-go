package kafkaclient

import (
	"context"
	"errors"
	"github.com/3lvia/libraries-go/pkg/mschema"
	"net/http"
)

const defaultFormat = mschema.AVRO

func Hello() {}

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

	registry, err := mschema.New(
		secrets.registryURL,
		mschema.WithClient(client),
		mschema.WithUser(secrets.registryKey, secrets.registrySecret))
	if err != nil {
		return nil, err
	}

	format := collector.format
	if !collector.formatSet {
		format = defaultFormat
	}

	creator := creatorFunc(format, collector)

	c, err := newConsumer(
		system,
		topic,
		application,
		secrets.bootstrapServer,
		secrets.key,
		secrets.secret,
		format,
		creator,
		registry,
		collector.offsetSender)
	if err != nil {
		return nil, err
	}

	ch := make(chan *StreamingMessage)
	go c.start(ctx, ch)

	return ch, nil
}
