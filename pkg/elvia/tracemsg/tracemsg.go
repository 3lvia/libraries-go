package tracemsg

import (
	"context"

	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

// NewHeader creates a new nats.Header with the current context and injects the OpenTelemetry context into it.
func NewHeader(ctx context.Context) nats.Header {
	headers := make(nats.Header)
	propagator := otel.GetTextMapPropagator()
	propagator.Inject(ctx, propagation.HeaderCarrier(headers))

	return headers
}

// Extract extracts the OpenTelemetry context from the nats.Header into a new context.
func Extract(ctx context.Context, header nats.Header) context.Context {
	propagator := otel.GetTextMapPropagator()
	ctx = propagator.Extract(ctx, propagation.HeaderCarrier(header))
	return ctx
}

// Inject injects the OpenTelemetry context from the current context into the nats.Header.
func Inject(ctx context.Context, header nats.Header) nats.Header {
	headers := make(nats.Header)
	propagator := otel.GetTextMapPropagator()
	propagator.Inject(ctx, propagation.HeaderCarrier(headers))

	for k, v := range headers {
		header[k] = v
	}

	return header
}
