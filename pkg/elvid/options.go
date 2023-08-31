package elvid

import (
	"net/http"
	"os"
)

const (
	discoveryEndpoint = "/.well-known/openid-configuration"
)

type Config struct {
	httpClient *http.Client

	cert      string
	address   string
	discovery string

	errorHandler errorHandlerFunc
}

type errorHandlerFunc func(err error)

func newConfig(opts ...Option) Config {
	var c Config
	for _, option := range opts {
		option.apply(&c)
	}

	if c.cert == "" {
		c.cert = os.Getenv("ELVID_CACERT")
	}

	if c.address == "" {
		c.address = os.Getenv("ELVID_BASE_URL")
	}

	if c.discovery == "" {
		c.discovery = os.Getenv("ELVID_DISCOVERY")

		if c.discovery == "" {
			c.discovery = discoveryEndpoint
		}
	}

	return c
}

type Option interface {
	apply(*Config)
}

type optionFunc func(*Config)

func (fn optionFunc) apply(cfg *Config) {
	fn(cfg)
}

// WithClient lets you set your own http.Client that will be used internally.
// Make sure the httpClient is configured to be secure since unsecure clients can
// pose security risks.
func WithClient(httpClient *http.Client) Option {
	return optionFunc(func(c *Config) {
		c.httpClient = httpClient
	})
}

// WithAddress sets the base url to use to access ElvID.
// If this is left empty, it will be loaded from ELVID_BASE_URL env var.
func WithAddress(address string) Option {
	return optionFunc(func(c *Config) {
		c.address = address
	})
}

// WithDiscovery sets the discovery uri to use to access ElvID (appended to address).
// If this is left empty, it will be loaded from ELVID_DISCOVERY env var.
// Otherwise, it will use the standard endpoint for loading openid connect configuration.
// See https://en.wikipedia.org/wiki/Well-known_URI for more information.
func WithDiscovery(discovery string) Option {
	return optionFunc(func(c *Config) {
		c.discovery = discovery
	})
}

// WithErrorHandler sets a callback function that will be called if there are
// any errors in a background routine.
func WithErrorHandler(fn errorHandlerFunc) Option {
	return optionFunc(func(c *Config) {
		c.errorHandler = fn
	})
}