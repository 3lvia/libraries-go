package elvia

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/3lvia/libraries-go/pkg/elvia/runtime"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func Test_Service(t *testing.T) {
	ctx := context.Background()

	t.Run("service should run", func(t *testing.T) {
		t.Parallel()

		otelWasCalled := make([]bool, 3)

		opts := []ServiceOpt{
			WithEnv(runtime.Development),
			WithLoggerLevel(slog.LevelDebug),
			WithOTELAttributes(
				semconv.ServiceVersion("v0.0.0"),
				semconv.DeploymentEnvironment(runtime.Development.String()),
			),
			WithOTELTraceProvider(func(ctx context.Context, env runtime.Env, opts ...trace.TracerProviderOption) (*trace.TracerProvider, error) {
				otelWasCalled[0] = true
				return nil, nil
			}),
			WithOTELLoggerProvider(func(ctx context.Context, env runtime.Env, opts ...log.LoggerProviderOption) (*log.LoggerProvider, error) {
				otelWasCalled[1] = true
				return nil, nil
			}),
			WithOTELMetricProvider(func(ctx context.Context, env runtime.Env, opts ...metric.Option) (*metric.MeterProvider, error) {
				otelWasCalled[2] = true
				return nil, nil
			}),
			WithAPI(DisableAPI),
		}

		svc, err := NewService(ctx, "test", "test", opts...)
		require.NoError(t, err)

		require.True(t, otelWasCalled[0])
		require.True(t, otelWasCalled[1])
		require.True(t, otelWasCalled[2])

		ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
		defer cancel()

		if err := svc.Run(ctx); err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}