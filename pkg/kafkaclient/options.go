package kafkaclient

import (
	"github.com/3lvia/libraries-go/pkg/hashivault"
	"github.com/3lvia/libraries-go/pkg/mschema"
	"net/http"
)

type optionsCollector struct {
	secrets     hashivault.SecretsManager
	client      *http.Client
	key         *apiKey
	creatorFunc EntityCreatorFunc
	format      mschema.Type
	formatSet   bool
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

// WithEntityCreatorFunc sets the function to use when creating entities from messages. If not set, the default
// implementation will be used, which simply returns the message data as is, i.e. the raw byte slice from the Kafka
// message.
func WithEntityCreatorFunc(creator EntityCreatorFunc) Option {
	return func(o *optionsCollector) {
		o.creatorFunc = creator
	}
}

// UseAVRO sets the message format to AVRO. The reason for setting the format at all is so that the package can
// provide a default implementation of the EntityCreatorFunc. If the client provides a custom implementation of
// EntityCreatorFunc, the format is not used.
func UseAVRO() Option {
	return func(o *optionsCollector) {
		o.format = mschema.AVRO
		o.formatSet = true
	}
}

// UseJSON sets the message format to JSON. The reason for setting the format at all is so that the package can
// provide a default implementation of the EntityCreatorFunc. If the client provides a custom implementation of
// EntityCreatorFunc, the format is not used.
func UseJSON() Option {
	return func(o *optionsCollector) {
		o.format = mschema.JSON
		o.formatSet = true
	}
}

// UseProtobuf sets the message format to Protobuf. The reason for setting the format at all is so that the package can
// provide a default implementation of the EntityCreatorFunc. If the client provides a custom implementation of
// EntityCreatorFunc, the format is not used.
func UseProtobuf() Option {
	return func(o *optionsCollector) {
		o.format = mschema.PROTOBUF
		o.formatSet = true
	}
}
