package kafkaclient

import (
	"context"
	"errors"
	"github.com/3lvia/libraries-go/pkg/mschema"
	"net/http"
)

const defaultFormat = mschema.AVRO

func StartConsumer(ctx context.Context, system, topic, application string, opts ...Option) (Starter, AdminClient, func(), error) {
	collector := &optionsCollector{}
	for _, opt := range opts {
		opt(collector)
	}

	emptyCloser := func() {}

	if collector.secrets == nil {
		return nil, nil, emptyCloser, errors.New("no secrets manager provided")
	}

	client := collector.client
	if client == nil {
		client = &http.Client{}
	}

	secrets, err := getSecrets(ctx, system, collector.secrets)
	if err != nil {
		return nil, nil, emptyCloser, err
	}

	registry, err := mschema.New(
		secrets.registryURL,
		mschema.WithClient(client),
		mschema.WithUser(secrets.registryKey, secrets.registrySecret))
	if err != nil {
		return nil, nil, emptyCloser, err
	}

	format := collector.format
	if !collector.formatSet {
		format = defaultFormat
	}

	creator := creatorFunc(format, collector)

	info := kafkaClientInfo{
		broker:   secrets.bootstrapServer,
		userName: secrets.key,
		password: secrets.secret,
		clientID: "libraries-test", // TODO: clientID should be unique for each instance of the consumer
		// should be fetched from Kuberenetes
	}

	kClient, err := fetchClient(topic, application, info)
	if err != nil {
		return nil, nil, emptyCloser, err
	}

	c, err := newConsumer(
		kClient,
		format,
		creator,
		registry,
		collector.offsetSender)
	if err != nil {
		return nil, nil, kClient.Close, err
	}

	admClient, admCloser, err := newAdminClient(application, kClient)
	if err != nil {
		return nil, nil, kClient.Close, err
	}

	starter := func(ctx context.Context) <-chan *StreamingMessage {
		ch := make(chan *StreamingMessage)
		go c.start(ctx, ch)
		return ch
	}

	finalCloser := func() {
		admCloser()
		kClient.Close()
	}

	return starter, admClient, finalCloser, nil
}
