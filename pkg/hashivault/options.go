package hashivault

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/3lvia/libraries-go/pkg/hashivault/auth"
)

type optionsCollector struct {
	client              *http.Client
	vaultAddress        string
	gitHubToken         string
	k8sMountPath        string
	k8sRole             string
	useOIDC             bool
	useOIDCDisableCache bool
	vaultToken          string
	otelTracerName      string
	logger              *log.Logger
}

// Option is a function that can be used to configure this package.
type Option func(*optionsCollector)

// WithClient sets the http client to use when making requests to Vault. This is useful for testing.
func WithClient(client *http.Client) Option {
	return func(o *optionsCollector) {
		o.client = client
	}
}

// WithGitHubToken sets the GitHub token to use when authenticating to Vault.
// deprecated: use WithOIDC instead
func WithGitHubToken(token string) Option {
	return func(o *optionsCollector) {
		o.gitHubToken = token
	}
}

// WithOIDC sets the authentication method to OIDC. This will use the default OIDC configuration.
func WithOIDC() Option {
	return func(o *optionsCollector) {
		o.useOIDC = true
		o.useOIDCDisableCache = false
	}
}

// WithOIDCCached sets the authentication method to OIDC and sets whether to cache the OIDC token.
func WithOIDCCached(useCache bool) Option {
	return func(o *optionsCollector) {
		o.useOIDC = true
		o.useOIDCDisableCache = !useCache
	}
}

// WithVaultToken sets the Vault token to use when authenticating to Vault.
func WithVaultToken(token string) Option {
	return func(o *optionsCollector) {
		o.vaultToken = token
	}
}

// WithVaultAddress sets the address of the Vault server to use when making requests to Vault.
func WithVaultAddress(address string) Option {
	return func(o *optionsCollector) {
		o.vaultAddress = address
	}
}

// WithKubernetes sets the Kubernetes mount path and role to use when authenticating to Vault.
func WithKubernetes(mountPath, role string) Option {
	return func(o *optionsCollector) {
		o.k8sMountPath = mountPath
		o.k8sRole = role
	}
}

// WithOtelTracerName sets the name of the OpenTelemetry tracer to use when creating spans. If no name is set the
// tracer name "go.opentelemetry.io/otel" is used.
func WithOtelTracerName(name string) Option {
	return func(o *optionsCollector) {
		o.otelTracerName = name
	}
}

// WithLogger sets the logger to use when logging. If no logger is set a noop logger is used.
func WithLogger(logger *log.Logger) Option {
	return func(o *optionsCollector) {
		o.logger = logger
	}
}

func (c *optionsCollector) authMethod() auth.Method {
	if c.vaultToken != "" {
		return auth.MethodToken
	}
	if c.k8sMountPath != "" {
		return auth.MethodK8s
	}
	if c.useOIDC {
		return auth.MethodOIDC
	}

	return auth.MethodGitHub
}

func (c *optionsCollector) build() error {
	va := os.Getenv("VAULT_ADDR")
	if va != "" {
		c.vaultAddress = va
	}

	ght := os.Getenv("GITHUB_TOKEN")
	if ght != "" {
		c.gitHubToken = ght
	}

	k8sMP := os.Getenv("MOUNT_PATH")
	if k8sMP != "" {
		c.k8sMountPath = k8sMP
	}

	k8sR := os.Getenv("ROLE")
	if k8sR != "" {
		c.k8sRole = k8sR
	}

	vt := os.Getenv("VAULT_TOKEN")
	if vt != "" {
		c.vaultToken = vt
	}

	if c.vaultAddress == "" {
		return fmt.Errorf("VAULT_ADDR not set")
	}
	if c.vaultToken != "" {
		return nil
	}
	if c.useOIDC {
		return nil
	}
	if c.gitHubToken == "" && c.k8sMountPath == "" {
		return fmt.Errorf("GITHUB_TOKEN or MOUNT_PATH not set")
	}
	return nil
}