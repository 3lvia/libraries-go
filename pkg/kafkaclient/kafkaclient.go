package kafkaclient

import (
	"context"
	"errors"
	"github.com/3lvia/libraries-go/pkg/mschema"
	"net/http"
)

func Start(ctx context.Context, system, topic, groupID, clientID string, opts ...Option) (<-chan *StreamingMessage, error) {
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

	cv, err := getSecrets(ctx, system, collector.secrets)
	if err != nil {
		return nil, err
	}
	_ = cv

	creator := func(data []byte, schema string) interface{} {
		return data
	}

	registry, err := mschema.New(cv.registryURL, mschema.WithClient(c), mschema.WithUser(cv.registryKey, cv.registrySecret))
	if err != nil {
		return nil, err
	}

	r, err := newReader(
		topic,
		cv.bootstrapServer,
		groupID,
		clientID,
		cv.key,
		cv.secret,
		creator,
		registry)
	if err != nil {
		return nil, err
	}

	ch := make(chan *StreamingMessage)
	go r.start(ctx, ch)

	return ch, nil
}
