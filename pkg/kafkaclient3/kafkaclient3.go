package kafkaclient3

import (
	"context"
	"errors"
	"github.com/3lvia/libraries-go/pkg/mschema"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"net/http"
)

func StartConsumer(ctx context.Context, system, topic, consumerGroup string, opts ...Option) (<-chan *StreamingMessage, error) {
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

	secrets, err := getSecrets(ctx, system, collector.secrets, collector.key)
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

	// TODO: Make configurable
	format := mschema.AVRO

	c := config(secrets)

	ch, err := startConsumer(ctx, topic, consumerGroup, registry, format, c)
	if err != nil {
		return nil, err
	}
	return ch, nil
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
