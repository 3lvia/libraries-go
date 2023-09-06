/*
Package hashivault provides a Vault client for the Hashicorp Vault secrets management solution.

AUTHENTICATION
Four modes of authentication against Vault are supported(here listed according to precedence):
1. Vault tokens (for people), usually in debugging situations where the other methods are not available
2. Kubernetes authentication for pods
3. Azure AD SSO authentication (OICD) for people
4. GitHub authentication for people

The package can be configured via the options pattern, i.e. by sending a number of options to the New function.
However, environment variables can also be used to configure this package. Configuration via environment variables
takes precedence over configuration via the options pattern. The following environment variables are supported:
 1. GITHUB_TOKEN. If this variable is set, the client will authenticate against Vault using the GitHub auth method.
    This takes precedence over the other methods, except pre-authentication, se 4) below.
 2. MOUNT_PATH and ROLE. If these variables are set, the client will authenticate using the Kubernetes auth method.
 3. VAULT_ADDR. This variable must be set to the address of the Vault server.
 4. VAULT_TOKEN. If this variable is set, the client will be pre-authenticated, and will use the supplied token for
    all requests to Vault. This takes precedence over the other methods.

OPTIONS
The following options are supported:
 1. WithClient. This option can be used to set the http client to use when making requests to Vault. This is useful
    for testing, but should normally not be used in production code.
 2. WithOIDC. This option can be used to set the authentication method to OIDC.
 3. WithVaultToken. This option can be used to set the Vault token to use when authenticating to Vault.
 4. WithVaultAddress. This option can be used to set the address of the Vault server to use when making requests to
    Vault.
 5. WithGitHubToken. This option can be used to set the GitHub token to use when authenticating to Vault.
 6. WithKubernetes. This option can be used to set the Kubernetes mount path and role to use when authenticating to
    Vault.
 7. WithOtelTracerName. This option can be used to set the name of the OpenTelemetry tracer to use when creating
    spans. If no name is set the tracer name "go.opentelemetry.io/otel" is used.
 8. WithLogger. This option can be used to set the logger to use when logging. If no logger is set a noop logger is
    used.

RECOMMENDED SETUP (ELVIA)
In the context of developing and running services in Elvia, the recommended approach is to use OICD authentication
against Azure AD while developing, while Kubernetes authentication is used when running in the Kubernetes cluster.
The latter is configured via environment variables, while the former is configured via the options pattern. Thus, the
recommended approach is to use the following code in the main function:
```
v, errChan, err := hashivault.New(hashivault.WithOIDC(), hashivault.WithVaultAddress("https://vault.dev-elvia.io"))

	if err != nil {
	  log.Fatal(err)
	}

```
This ensures that the client will authenticate against Vault using OICD when running on the development machine,
while it will use Kubernetes authentication when running in the Kubernetes cluster (because the environment
variables MOUNT_PATH and ROLE will be set).

RENEWAL OF TOKEN
The client will periodically renew the authentication token. The token is renewed when it has less than 30 seconds
left to live. The token is renewed in a separate goroutine, so the client will not block while waiting for the token
to be renewed.

INSTRUMENTATION
The package uses the OpenTelemetry SDK for Go for tracing as well as *log.Logger for simple logging. It is up to the
client to configure the tracer and logger. The logger is set with the option WithLogger. If not set, a noop logger is
used by default. Tracing is configured via otel.SetTracerProvider

GENERAL USAGE AND ABSTRACTIONS
The main abstraction of this package is the SecretsManager interface. A new instance of SecretsManager is created
with the New function. The returned SecretsManager is safe to use concurrently. Normally, only a single instance of
SecretsManager should be created in an application, presumably in the main function as the application is being
initialised.

The secrets that are returned by the SecretsManager are represented as functions that return a map[string]any. The
point is that the returned function will always return the latest version of the secret. Therefore, clients should
save a reference to the function rather than saving the actual secrets, and invoke the func just-in-time as the
secret is needed. The returned function is safe to use concurrently.

The SecretsManager interface also provides a method for setting the default Google credentials for the current
process.

The token refresh functionality runs in a separate goroutine, and also a new goroutine will be started for each
fetched secret that is renewable and has a lease duration. In order to communicate errors from these goroutines, the
New function returns a channel of errors in addition to the SecretsManager. Clients should start a goroutine that
reads from this channel and handles errors as appropriate.

The following example shows how to use the SecretsManager:
```
import (

	"context"
	"github.com/3lvia/hashivault-go/pkg/hashivault"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
	"log"
	"os"

)

	func main() {
		ctx := context.Background()

		// Use stout logger, replace with more sophisticated logger in production.
		l := log.New(os.Stdout, "", log.LstdFlags)

		// Use stdout exporter, replace with more sophisticated exporter in production.
		exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
		if err != nil {
			log.Fatal(err)
		}

		tp := trace.NewTracerProvider(
			trace.WithSyncer(exporter),
			trace.WithSampler(trace.AlwaysSample()),
		)
		otel.SetTracerProvider(tp)

		v, errChan, err := hashivault.New(
			ctx,
			hashivault.WithOIDC(),
			hashivault.WithVaultAddress("https://vault.dev-elvia.io"),
			hashivault.WithLogger(l),
		)
		if err != nil {
			log.Fatal(err)
		}

		go func(ec <-chan error) {
			for err := range ec {
				log.Println(err)
			}
		}(errChan)

		secret, err := v.GetSecret(ctx, "kunde/kv/data/appinsights/kunde")
		if err != nil {
			log.Fatal(err)
		}

		mapOfSecrets := secret()
		_ = mapOfSecrets
	}

```
*/
package hashivault
