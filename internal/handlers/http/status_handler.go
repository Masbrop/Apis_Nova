package http

import (
	"context"
	"net/http"
	"time"

	"apis_nova/internal/domain/status"
)

type StatusHandler struct {
	service *status.Service
	timeout time.Duration
}

func NewStatusHandler(service *status.Service) *StatusHandler {
	return &StatusHandler{
		service: service,
		timeout: 2 * time.Second,
	}
}

func (h *StatusHandler) GetHealth(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), h.timeout)
	defer cancel()

	snapshot, err := h.service.Check(ctx)
	if err != nil {
		writeJSON(w, http.StatusServiceUnavailable, snapshot)
		return
	}

	writeJSON(w, http.StatusOK, snapshot)
}
