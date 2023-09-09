package mschema

import "net/http"

type optionsCollector struct {
	user           string
	password       string
	otelTracerName string
	client         *http.Client
}

// Option is a function that can be passed to Init to configure the mschema package.
type Option func(*optionsCollector)

// WithUser sets the user and password to use when connecting to the schema registry.
func WithUser(user, password string) Option {
	return func(collector *optionsCollector) {
		collector.user = user
		collector.password = password
	}
}

// WithClient sets the HTTP client to use when connecting to the schema registry. If not set, the default HTTP client
// will be used.
func WithClient(client *http.Client) Option {
	return func(collector *optionsCollector) {
		collector.client = client
	}
}

// WithOtelTracerName sets the name of the OpenTelemetry tracer to use when creating spans. If no name is set the
// tracer name "go.opentelemetry.io/otel" is used.
func WithOtelTracerName(name string) Option {
	return func(o *optionsCollector) {
		o.otelTracerName = name
	}
}
