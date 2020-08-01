package services

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/StenvL/interest-points-api/models/domain"
	"github.com/StenvL/interest-points-api/models/requests"
	"github.com/StenvL/interest-points-api/models/responses"
	"github.com/StenvL/interest-points-api/store"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// AuthenticationService service to authenticate user.
type AuthenticationService struct {
	store *store.Store
}

// NewAuthenticationService creates an instance of AuthenticationService.
func NewAuthenticationService(store *store.Store) *AuthenticationService {
	return &AuthenticationService{
		store: store,
	}
}

// Authenticate checks user credentials and returns access token if credentials are valid.
func (a *AuthenticationService) Authenticate(r *requests.AuthRequest) (*responses.TokenResponse, error) {
	user, err := a.store.User().GetByLogin(r.Login)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("User with login %s not found", r.Login)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(r.Password))
	if err != nil {
		return nil, errors.New("Wrong login or password")
	}

	tokenData := &domain.AuthToken{
		Name: user.Login,
		StandardClaims: jwt.StandardClaims{
			Subject: strconv.Itoa(int(user.ID)),
			// Hardcodding token's lifetime is not the best way, just wanted to simplify a bit.
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenData)
	tokenString, _ := token.SignedString([]byte("hwMAKnOG4w6FEMDCdgZv"))
	tokenResponse := &responses.TokenResponse{
		AccessToken: tokenString,
	}

	return tokenResponse, nil
}
