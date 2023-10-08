package kafkaclient

import (
	"context"
	"errors"
	"github.com/3lvia/libraries-go/pkg/mschema"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/linkedin/goavro/v2"
	"net/http"
)

func StartConsumer(ctx context.Context, system, topic, consumerGroup string, opts ...Option) (<-chan *StreamingMessage, func() error, error) {
	collector := &optionsCollector{}
	for _, opt := range opts {
		opt(collector)
	}

	if collector.secrets == nil {
		return nil, nil, errors.New("no secrets manager provided")
	}

	client := collector.client
	if client == nil {
		client = &http.Client{}
	}

	secrets, err := getSecrets(ctx, system, collector.secrets, collector.key)
	if err != nil {
		return nil, nil, err
	}

	registry, err := mschema.New(
		secrets.registryURL,
		mschema.WithClient(client),
		mschema.WithUser(secrets.registryKey, secrets.registrySecret))
	if err != nil {
		return nil, nil, err
	}

	creator := creatorFunc(collector)

	c := config(secrets)

	ch, closer, err := startConsumer(ctx, topic, consumerGroup, creator, registry, c)
	if err != nil {
		return nil, nil, err
	}
	return ch, closer, nil
}

func config(secrets *secretConfigValues) kafka.ConfigMap {
	return map[string]kafka.ConfigValue{
		"security.protocol": "SASL_SSL",
		"sasl.mechanisms":   "PLAIN",
		"bootstrap.servers": secrets.bootstrapServer,
		"sasl.username":     secrets.key,
		"sasl.password":     secrets.secret,
		"acks":              "all",
	}
}

func creatorFunc(collector *optionsCollector) EntityCreatorFunc {
	creator := collector.creatorFunc
	if creator != nil {
		// If the client has provided a custom creator function, we use that, regardless of the format.
		return creator
	}
	if collector.formatSet && collector.format == mschema.AVRO {
		// Since collector.creatorFunc == nil, and the schema is AVRO, we use dynamic typing.
		return defaultAVROCreator
	}

	// The default creator function simply returns the raw byte slice from the Kafka message.
	return func(data []byte, d mschema.Descriptor) (any, error) {
		return data, nil
	}
}

func defaultAVROCreator(value []byte, d mschema.Descriptor) (any, error) {
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
