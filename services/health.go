package services

import (
	"github.com/StenvL/interest-points-api/models/responses"
	"github.com/StenvL/interest-points-api/store"
)

// HealthService service to check API server health.
type HealthService struct {
	store *store.Store
}

// NewHealthService creates an instance of HealthService.
func NewHealthService(store *store.Store) *HealthService {
	return &HealthService{
		store: store,
	}
}

// GetHealthStatus checks health status of API server.
func (h *HealthService) GetHealthStatus() []*responses.ServiceStatusResponse {
	storeStatus := &responses.ServiceStatusResponse{
		Name:    "Store",
		Status:  "OK",
		Details: "",
	}

	if err := h.store.Ping(); err != nil {
		storeStatus.Status = "ERROR"
		storeStatus.Details = err.Error()
	}

	return []*responses.ServiceStatusResponse{storeStatus}
}
