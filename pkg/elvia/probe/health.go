package probe

// HealthStatus represents the status of the health of a service.
type HealthStatus string

const (
	// Healthy indicates that the service is healthy.
	Healthy HealthStatus = "healthy"
	// Degraded indicates that the service is degraded.
	Degraded HealthStatus = "degraded"
	// Unhealthy indicates that the service is unhealthy.
	Unhealthy HealthStatus = "unhealthy"
)

// String returns the string representation of the health status.
func (s HealthStatus) String() string { return string(s) }

// HealthReport represents the health of a service.
type HealthReport struct {
	Status HealthStatus `json:"status"`
	Error  string       `json:"error,omitempty"`
}

// HealthSummary represents a summary of the health of the services.
type HealthSummary struct {
	Status   HealthStatus            `json:"status"`
	Services map[string]HealthReport `json:"services,omitempty"`
}

// HealthReportFunc is a function that returns the health of a service.
type HealthReportFunc func() HealthReport

// HealthChecks is a map of health check functions.
type HealthChecks map[string]HealthReportFunc

// HealthReports is a map of health reports.
type HealthReports map[string]HealthReport

// Check returns a summary of the health of the services.
func Check(health HealthReports) HealthSummary {
	status := Healthy

	for _, s := range health {
		if s.Status == Unhealthy {
			status = Unhealthy
			break
		}

		if s.Status == Degraded {
			status = Degraded
		}
	}

	return HealthSummary{
		Status:   status,
		Services: health,
	}
}