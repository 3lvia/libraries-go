package auth

import (
	"log"
	"net/http"
)

type optionsCollector struct {
	client *http.Client

	gitHubToken string

	k8sServicePath string
	k8sRole        string

	l              *log.Logger
	otelTracerName string

	disableLocalCache bool
}

// Option is a function that can be used to configure this package.
type Option func(*optionsCollector)

// WithClient sets the http client to use for requests
func WithClient(client *http.Client) Option {
	return func(o *optionsCollector) {
		o.client = client
	}
}

// WithGitHubToken sets the GitHub token to use for authentication
// Deprecated: Use WithToken or WithOICD instead
func WithGitHubToken(token string) Option {
	return func(o *optionsCollector) {
		o.gitHubToken = token
	}
}

// WithK8s sets the Kubernetes service path and role to use for authentication
func WithK8s(servicePath, role string) Option {
	return func(o *optionsCollector) {
		o.k8sServicePath = servicePath
		o.k8sRole = role
	}
}

// WithLogger sets the logger to use for logging.
func WithLogger(l *log.Logger) Option {
	return func(o *optionsCollector) {
		o.l = l
	}
}

// WithOtelTracerName sets the name of the OpenTelemetry tracer to use when creating spans. If no name is set the
// tracer name "go.opentelemetry.io/otel" is used.
func WithOtelTracerName(name string) Option {
	return func(o *optionsCollector) {
		o.otelTracerName = name
	}
}

// DisableLocalCache disables the local cache of secrets. Is only applicable when using the OICD authentication method
// on the local development machine.
func DisableLocalCache() Option {
	return func(o *optionsCollector) {
		o.disableLocalCache = true
	}
}
