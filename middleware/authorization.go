package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/StenvL/interest-points-api/models/domain"

	"github.com/StenvL/interest-points-api/utils"
	"github.com/dgrijalva/jwt-go"
)

// Authorization is a middleware method that checks JWT for requests to api server.
func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notAuth := []string{"/api/authenticate", "/api/health"}
		requestPath := r.URL.Path

		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			utils.JSONError(w, "Authorization failed", "Missing auth token", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(tokenHeader, " ")
		if len(tokenParts) != 2 {
			utils.JSONError(w, "Authorization failed", "Invalid/Malformed auth token", http.StatusUnauthorized)
			return
		}

		tokenStr := tokenParts[1]
		tokenModel := &domain.AuthToken{}

		token, err := jwt.ParseWithClaims(tokenStr, tokenModel, func(token *jwt.Token) (interface{}, error) {
			return []byte("hwMAKnOG4w6FEMDCdgZv"), nil
		})

		if err != nil {
			utils.JSONError(w, "Authorization failed", "Malformed authentication token", http.StatusForbidden)
			return
		}

		if !token.Valid {
			utils.JSONError(w, "Authorization failed", "Token is not valid.", http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), "user", tokenModel.Subject)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
