package backend

// HealthStatus represents the API health response.
type HealthStatus struct {
	Status string `json:"status"`
}

// Health returns the current API health status.
func Health() HealthStatus {
	return HealthStatus{Status: "ok"}
}
