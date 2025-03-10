package observability

import "os"

const (
	OTELEnabled              = "OTEL_ENABLED"
	OTELExporterOTLPEndpoint = "OTEL_EXPORTER_OTLP_ENDPOINT"

	// OTELLogLevel is currently not supported by the go SDK.
	// We provide simple support for it here.
	// Follow the issue:
	// https://github.com/open-telemetry/opentelemetry-go/issues/2303
	OTELLogLevel = "OTEL_LOG_LEVEL"

	OTELExporterOTLPEndpointDefault = "localhost:4317"
)

// IsOTELEnabled returns true if OpenTelemetry is enabled.
func IsOTELEnabled() bool {
	ev := os.Getenv(OTELEnabled)

	if ev == "false" {
		return false
	} else if ev == "true" {
		return true
	}

	return true
}

// GetOTELExporterOTLPEndpoint returns the OpenTelemetry OTLP exporter endpoint.
func GetOTELExporterOTLPEndpoint() string {
	ev := os.Getenv(OTELExporterOTLPEndpoint)

	if ev == "" {
		return OTELExporterOTLPEndpointDefault
	}

	return ev
}

// GetOTELLogLevel returns the OpenTelemetry log level.
func GetOTELLogLevel() string {
	ev := os.Getenv(OTELLogLevel)
	return ev
}