package schemaclient

import (
	"context"
	"errors"
	"net/http"
)

func New(ctx context.Context, system string, opts ...Option) (SchemaRegistry, error) {
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

	r, err := newRegistry(ctx, system, collector.secrets, c)
	if err != nil {
		return nil, err
	}

	return r, nil
}
