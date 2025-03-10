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

// HealthReport represents the probe of a service.
type HealthReport struct {
	Status HealthStatus `json:"status"`
	Name   string       `json:"name"`
	Error  string       `json:"error,omitempty"`
}

// HealthSummary represents a summary of the probe of the services.
type HealthSummary struct {
	Status   HealthStatus   `json:"status"`
	Services []HealthReport `json:"services,omitempty"`
}

// HealthCheckFunc is a function that returns the probe of a service.
type HealthCheckFunc func() HealthReport

// HealthChecksFunc is a function that returns the probe of all services.
type HealthChecksFunc func() []HealthReport

// Check returns a summary of the probe of the services.
func Check(health []HealthReport) HealthSummary {
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