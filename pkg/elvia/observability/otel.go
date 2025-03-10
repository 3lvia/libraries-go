package observability

import "os"

const (
	otelEnabled = "OTEL_ENABLED"

	// otelLogLevel is currently not supported by the go SDK.
	// We provide simple support for it here.
	// Follow the issue:
	// https://github.com/open-telemetry/opentelemetry-go/issues/2303
	otelLogLevel = "OTEL_LOG_LEVEL"
)

// IsOTELEnabled returns true if OpenTelemetry is enabled.
func IsOTELEnabled() bool {
	ev := os.Getenv(otelEnabled)

	if ev == "false" {
		return false
	} else if ev == "true" {
		return true
	}

	return true
}

// GetOTELLogLevel returns the OpenTelemetry log level.
func GetOTELLogLevel() string {
	ev := os.Getenv(otelLogLevel)
	return ev
}