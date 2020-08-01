package domain

import "github.com/dgrijalva/jwt-go"

// AuthToken is a struct to store data for access token.
type AuthToken struct {
	Name string `json:"name,omitempty"`
	jwt.StandardClaims
}
