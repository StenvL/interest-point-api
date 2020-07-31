package apiserver

import (
	"net/http"

	"github.com/StenvL/interest-points-api/store"
)

//Config type for configuring API service
type Config struct {
	BindAddr           string   `toml:"bind_addr"`
	CorsAllowedOrigins []string `toml:"cors_allowed_origins"`
	CorsAllowedMethods []string `toml:"cors_allowed_methods"`
	CorsAllowedHeaders []string `toml:"cors_allowed_headers"`
	Store              *store.Config
}

//NewConfig creates default config
func NewConfig() *Config {
	return &Config{
		BindAddr:           ":8080",
		CorsAllowedOrigins: []string{"*"},
		CorsAllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		CorsAllowedHeaders: []string{"*"},
		Store:              store.NewConfig(),
	}
}
