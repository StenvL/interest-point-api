package domain

import "github.com/dgrijalva/jwt-go"

// AuthToken is a struct to store data for access token.
type AuthToken struct {
	UserID uint64
	Login  string
	jwt.StandardClaims
}
