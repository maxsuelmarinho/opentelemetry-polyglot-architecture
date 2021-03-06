package handler

import (
	"encoding/json"
	"net/http"
)

type HealthCheckHandler interface {
	HealthCheck(w http.ResponseWriter, r *http.Request)
}

func NewHealthCheckHandler() HealthCheckHandler {
	return &healthCheckHandler{}
}

type healthCheckHandler struct {
}

func (h *healthCheckHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	payload := map[string]string{
		"status": "UP",
	}
	response, _ := json.Marshal(payload)
	w.Write(response)
}
