package observability

import (
	"context"

	"github.com/3lvia/libraries-go/pkg/elvia/runtime"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/trace"
)

// NewDefaultTraceProvider creates a new trace provider with the default configuration.
// The default configuration is to use the OTLP gRPC exporter.
// For development environments, the exporter is synchronous. Otherwise, it is batched.
func NewDefaultTraceProvider(ctx context.Context, env runtime.Env, opts ...trace.TracerProviderOption) (*trace.TracerProvider, error) {
	exporter, err := otlptracegrpc.New(ctx)
	if err != nil {
		return nil, err
	}

	var options []trace.TracerProviderOption
	options = append(options, opts...)

	switch {
	case env == runtime.Development:
		options = append(options, trace.WithSyncer(exporter))
	default:
		options = append(options, trace.WithBatcher(exporter))
	}

	traceProvider := trace.NewTracerProvider(options...)
	return traceProvider, nil
}
