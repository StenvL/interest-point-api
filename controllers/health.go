package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/StenvL/interest-points-api/services"
)

// CheckHealth checks application services health.
func CheckHealth(h *services.HealthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(h.GetHealthStatus())
	}
}
