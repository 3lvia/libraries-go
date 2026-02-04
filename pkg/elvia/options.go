package elvia

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/3lvia/libraries-go/pkg/elvia/api"
	"github.com/3lvia/libraries-go/pkg/elvia/observability"
	"github.com/3lvia/libraries-go/pkg/elvia/probe"
	"github.com/3lvia/libraries-go/pkg/elvia/runtime"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.39.0"
)

const (
	// DefaultAPIAddr is the default address of the API.
	DefaultAPIAddr = ":8080"
	// DisableAPI is a special value to disable the API.
	DisableAPI = ""
)

// NewTraceProvider is a function to create a new trace provider.
type NewTraceProvider func(ctx context.Context, env runtime.Env, opts ...trace.TracerProviderOption) (*trace.TracerProvider, error)

// NewLoggerProvider is a function to create a new logger provider.
type NewLoggerProvider func(ctx context.Context, env runtime.Env, opts ...log.LoggerProviderOption) (*log.LoggerProvider, error)

// NewMetricProvider is a function to create a new metric provider.
type NewMetricProvider func(ctx context.Context, env runtime.Env, opts ...metric.Option) (*metric.MeterProvider, error)

// NewApiEngine is a function to create a new API engine.
type NewApiEngine func(env runtime.Env) *gin.Engine

// ConfigureApiEndpoint is a function to configure API endpoints.
type ConfigureApiEndpoint func(engine *gin.Engine)

// ConfigureApiHealthEndpoint is a function to configure the health endpoint.
type ConfigureApiHealthEndpoint func(engine *gin.Engine, fn func() probe.HealthReports)

// ServiceOpt is a function to configure the service.
type ServiceOpt func(*config)

type config struct {
	env runtime.Env

	loggerLevel slog.Level

	otelEnabled           bool
	otelAttributes        []attribute.KeyValue
	otelPropagator        propagation.TextMapPropagator
	otelNewTraceProvider  NewTraceProvider
	otelNewLoggerProvider NewLoggerProvider
	otelNewMetricProvider NewMetricProvider

	withApiAddr           string
	withHTTPServer        *http.Server
	withApiEngine         NewApiEngine
	withApiEndpoints      []ConfigureApiEndpoint
	withApiHealthEndpoint ConfigureApiHealthEndpoint
}

func defaultConfig(name string) config {
	return config{
		env:         runtime.Production,
		loggerLevel: slog.LevelWarn,
		otelEnabled: true,
		otelAttributes: []attribute.KeyValue{
			semconv.ServiceName(name),
		},
		otelPropagator: propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
		otelNewTraceProvider:  observability.NewDefaultTraceProvider,
		otelNewLoggerProvider: observability.NewDefaultLoggerProvider,
		otelNewMetricProvider: observability.NewDefaultMetricProvider,
		withApiAddr:           DefaultAPIAddr,
		withHTTPServer:        nil,
		withApiEngine:         api.NewDefaultEngine,
		withApiEndpoints: []ConfigureApiEndpoint{
			api.ConfigureStandardEndpoints,
			api.ConfigureStandardMetricsEndpoint,
		},
		withApiHealthEndpoint: api.ConfigureStandardHealthEndpoint,
	}
}

// WithEnv sets the environment of the service.
// This setting can affect the configuration of multiple components.
func WithEnv(env runtime.Env) ServiceOpt {
	return func(c *config) {
		c.env = env
	}
}

// WithEnvLoggerLevel sets the log level of the service based on the environment.
// The log level is set to debug for development and test environments, and to warn for production.
func WithEnvLoggerLevel(env runtime.Env) ServiceOpt {
	return func(c *config) {
		c.env = env
		logLevel := c.loggerLevel
		switch env {
		case runtime.Development, runtime.Test:
			logLevel = slog.LevelDebug
		}
		c.loggerLevel = logLevel
	}
}

// WithLoggerLevel sets the log level of the service.
// The default log level is warn.
func WithLoggerLevel(level slog.Level) ServiceOpt {
	return func(c *config) {
		c.loggerLevel = level
	}
}

// WithOTELDisabled disables OpenTelemetry for the service.
// OpenTelemetry is enabled by default.
// OTEL_ENABLED environment variable will override this setting.
func WithOTELDisabled() ServiceOpt {
	return func(c *config) {
		c.otelEnabled = false
		if observability.IsOTELEnabled() {
			c.otelEnabled = true
		}
	}
}

// WithOTELAttributes sets the OpenTelemetry attributes of the service.
// You don't have to set the service name attribute as it is set automatically.
// Schema version: 1.39.0
func WithOTELAttributes(attrs ...attribute.KeyValue) ServiceOpt {
	return func(c *config) {
		c.otelAttributes = append(c.otelAttributes, attrs...)
	}
}

// WithOTELPropagator sets the OpenTelemetry propagator of the service.
func WithOTELPropagator(propagation propagation.TextMapPropagator) ServiceOpt {
	return func(c *config) {
		c.otelPropagator = propagation
	}
}

// WithOTELTraceProvider sets the OpenTelemetry trace provider of the service.
func WithOTELTraceProvider(provider NewTraceProvider) ServiceOpt {
	return func(c *config) {
		c.otelNewTraceProvider = provider
	}
}

// WithOTELLoggerProvider sets the OpenTelemetry logger provider of the service.
func WithOTELLoggerProvider(provider NewLoggerProvider) ServiceOpt {
	return func(c *config) {
		c.otelNewLoggerProvider = provider
	}
}

// WithOTELMetricProvider sets the OpenTelemetry metric provider of the service.
func WithOTELMetricProvider(provider NewMetricProvider) ServiceOpt {
	return func(c *config) {
		c.otelNewMetricProvider = provider
	}
}

// WithAPI sets the API address of the service.
// The standard API configures a Gin engine with standard endpoints for health, metrics, and not found.
// See WithAPIEngine and WithAPIEndpoints for more control over the API, and optionally WithHTTPServer to set the HTTP server.
// If the address is empty, the API will be disabled. Use DisableAPI for this purpose.
func WithAPI(addr string) ServiceOpt {
	return func(c *config) {
		c.withApiAddr = addr
	}
}

// WithHTTPServer sets the HTTP server of the service.
func WithHTTPServer(server *http.Server) ServiceOpt {
	return func(c *config) {
		c.withHTTPServer = server
	}
}

// WithAPIEngine sets the API engine of the service.
func WithAPIEngine(engine NewApiEngine) ServiceOpt {
	return func(c *config) {
		c.withApiEngine = engine
	}
}

// WithAPIEndpoints sets the API endpoints of the service.
// The default endpoints are standard endpoints for metrics, and not found.
// To add custom endpoints, use this option, and optionally include the standard endpoints
// found in api.ConfigureStandardEndpoints, api.ConfigureStandardMetricsEndpoint, and api.ConfigureStandardHealthEndpoint.
func WithAPIEndpoints(endpoints ...ConfigureApiEndpoint) ServiceOpt {
	return func(c *config) {
		c.withApiEndpoints = endpoints
	}
}

// WithAPIHealthEndpoint sets the health endpoint of the service.
func WithAPIHealthEndpoint(endpoint ConfigureApiHealthEndpoint) ServiceOpt {
	return func(c *config) {
		c.withApiHealthEndpoint = endpoint
	}
}