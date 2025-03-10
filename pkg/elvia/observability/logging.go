package observability

import (
	"context"
	"strings"
	"sync"

	"github.com/3lvia/libraries-go/pkg/elvia/runtime"
	"go.opentelemetry.io/contrib/processors/minsev"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/log"
	sdklog "go.opentelemetry.io/otel/sdk/log"
)

var getSeverity = sync.OnceValue(func() log.Severity {
	conv := map[string]log.Severity{
		"":      log.SeverityInfo, // Default to SeverityInfo for unset.
		"debug": log.SeverityDebug,
		"info":  log.SeverityInfo,
		"warn":  log.SeverityWarn,
		"error": log.SeverityError,
	}
	// log.SeverityUndefined for unknown values.
	return conv[strings.ToLower(GetOTELLogLevel())]
})

type EnvSeverity struct{}

func (EnvSeverity) Severity() log.Severity { return getSeverity() }

// NewDefaultLoggerProvider creates a new logger provider with the default configuration.
// The default configuration is to use the OTLP gRPC exporter.
// For development environments, the exporter is synchronous. Otherwise, it is batched.
func NewDefaultLoggerProvider(ctx context.Context, env runtime.Env, opts ...sdklog.LoggerProviderOption) (*sdklog.LoggerProvider, error) {
	logExporter, err := otlploggrpc.New(ctx)
	if err != nil {
		return nil, err
	}

	var options []sdklog.LoggerProviderOption
	options = append(options, opts...)

	switch {
	case env == runtime.Development:
		options = append(options, sdklog.WithProcessor(minsev.NewLogProcessor(sdklog.NewSimpleProcessor(logExporter), EnvSeverity{})))
	default:
		options = append(options, sdklog.WithProcessor(minsev.NewLogProcessor(sdklog.NewBatchProcessor(logExporter), EnvSeverity{})))
	}

	loggerProvider := sdklog.NewLoggerProvider(options...)
	return loggerProvider, nil
}