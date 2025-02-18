package observability

import (
	"context"

	"github.com/3lvia/libraries-go/pkg/elvia/runtime"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
)

// NewDefaultMetricProvider creates a new metric provider with the default configuration.
// The default configuration is to use the Prometheus exporter.
func NewDefaultMetricProvider(ctx context.Context, env runtime.Env, opts ...metric.Option) (*metric.MeterProvider, error) {
	prom, err := prometheus.New()
	if err != nil {
		return nil, err
	}

	var options []metric.Option
	options = append(options, opts...)
	options = append(options, metric.WithReader(prom))

	provider := metric.NewMeterProvider(options...)

	return provider, nil
}
