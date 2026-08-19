package backend

type HealthStatus struct {
	Status string `json:"status"`
}

func Health() HealthStatus {
	return HealthStatus{Status: "ok"}
}