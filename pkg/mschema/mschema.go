package mschema

import (
	"net/http"
	"sync"
)

func New(opts ...Option) (Registry, error) {
	collector := &optionsCollector{}
	for _, opt := range opts {
		opt(collector)
	}

	schemaRegistryURL := collector.schemaRegistryURL
	user := collector.user
	password := collector.password

	currentHTTPClient := collector.client
	if currentHTTPClient == nil {
		currentHTTPClient = &http.Client{}
	}

	r := &registry{
		url:         schemaRegistryURL,
		user:        user,
		password:    password,
		client:      currentHTTPClient,
		descriptors: map[string]Descriptor{},
		mux:         &sync.RWMutex{},
	}

	return r, nil
}
