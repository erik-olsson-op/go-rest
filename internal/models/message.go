package models

type HealthCheck struct {
	Status string `json:"status"`
}

func NewHealthCheck(status string) *HealthCheck {
	return &HealthCheck{Status: status}
}

type ErrMessage struct {
	Message any `json:"message"`
}

func NewErrMessage(msg any) *ErrMessage {
	return &ErrMessage{Message: msg}
}
