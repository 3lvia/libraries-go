package kafkaclient

import (
	"context"
	"errors"
	"github.com/3lvia/libraries-go/pkg/mschema"
	"github.com/linkedin/goavro/v2"
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

	registry, err := mschema.New(
		secrets.registryURL,
		mschema.WithClient(client),
		mschema.WithUser(secrets.registryKey, secrets.registrySecret))
	if err != nil {
		return nil, err
	}

	format := collector.format
	if !collector.formatSet {
		format = mschema.AVRO
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
		registry)
	if err != nil {
		return nil, err
	}

	ch := make(chan *StreamingMessage)
	go c.start(ctx, ch)

	return ch, nil
}

func creatorFunc(format mschema.Type, collector *optionsCollector) EntityCreatorFunc {
	creator := collector.creatorFunc
	if creator != nil {
		return creator
	}
	if format == mschema.AVRO {
		// Since collector.creatorFunc == nil, and the schema is AVRO, we use dynamic typing.
		return func(value []byte, d mschema.Descriptor) (any, error) {
			codec, err := goavro.NewCodec(d.Schema())
			if err != nil {
				return nil, err
			}
			obj, _, err := codec.NativeFromBinary(value)
			if err != nil {
				return nil, err
			}
			return obj, nil
		}
	}

	return func(data []byte, d mschema.Descriptor) (any, error) {
		return data, nil
	}
}
