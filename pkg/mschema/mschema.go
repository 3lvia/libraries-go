package mschema

import (
	"net/http"
	"sync"
)

var (
	schemaRegistryURL string
	user              string
	password          string
)

var schemasByID = map[int]Descriptor{}
var mux = &sync.Mutex{}

var currentHTTPClient = &http.Client{}

func Init(opts ...Option) error {
	collector := &optionsCollector{}
	for _, opt := range opts {
		opt(collector)
	}

	schemaRegistryURL = collector.schemaRegistryURL
	user = collector.user
	password = collector.password

	currentHTTPClient = collector.client
	if currentHTTPClient == nil {
		currentHTTPClient = &http.Client{}
	}

	return nil
}
