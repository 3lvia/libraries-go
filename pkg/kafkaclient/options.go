package kafkaclient

import (
	"github.com/3lvia/libraries-go/pkg/hashivault"
	"net/http"
)

type optionsCollector struct {
	secrets     hashivault.SecretsManager
	creatorFunc EntityCreatorFunc
	client      *http.Client
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

// WithEntityCreatorFunc sets the function to use when creating entities from messages. If not set, the default
// implementation will be used, which simply returns the message data as is, i.e. the raw byte slice from the Kafka
// message.
func WithEntityCreatorFunc(creator EntityCreatorFunc) Option {
	return func(o *optionsCollector) {
		o.creatorFunc = creator
	}
}

// WithHTTPClient sets the HTTP client to use when communicating with the schema registry.
func WithHTTPClient(client *http.Client) Option {
	return func(o *optionsCollector) {
		o.client = client
	}
}
