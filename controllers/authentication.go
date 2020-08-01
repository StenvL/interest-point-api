package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/StenvL/interest-points-api/models/requests"
	"github.com/StenvL/interest-points-api/services"
	"github.com/StenvL/interest-points-api/utils"
)

// Authenticate authenticates user.
func Authenticate(a *services.AuthenticationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var authRequest *requests.AuthRequest
		err := json.NewDecoder(r.Body).Decode(&authRequest)

		if err != nil {
			utils.JSONError(w, "Request body is incorrect", err.Error(), http.StatusBadRequest)
			return
		}

		tokenResponse, err := a.Authenticate(authRequest)

		if err != nil {
			utils.JSONError(w, "Authentication failed", err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tokenResponse)
	}
}
