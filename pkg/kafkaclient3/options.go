package kafkaclient3

import (
	"github.com/3lvia/libraries-go/pkg/hashivault"
	"net/http"
)

type optionsCollector struct {
	secrets hashivault.SecretsManager
	client  *http.Client
	key     *apiKey
}

// Option is a function that can be used to configure this package.
type Option func(*optionsCollector)

// WithSecretsManager sets the secrets manager to use when fetching secrets. In future versions, configuration via
// environment variables may be supported, but for now, this is the only way to configure the package. Hence, it is
// currently not really optional, and an error is returned if no secrets manager is provided in StartConsumer/Producer.
func WithSecretsManager(secrets hashivault.SecretsManager) Option {
	return func(o *optionsCollector) {
		o.secrets = secrets
	}
}

// WithHTTPClient sets the HTTP client to use when communicating with the schema registry.
func WithHTTPClient(client *http.Client) Option {
	return func(o *optionsCollector) {
		o.client = client
	}
}

// WithAPIKey sets the API key to use when communicating with the schema registry. If set this will override the
// configuration in Vault.
func WithAPIKey(key, secret string) Option {
	return func(o *optionsCollector) {
		o.key = &apiKey{key, secret}
	}
}
