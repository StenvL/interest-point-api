package apiserver

import (
	"github.com/StenvL/interest-points-api/store"
)

//Config type for configuring API service
type Config struct {
	BindAddr string `toml:"bind_addr"`
	Store    *store.Config
}

//NewConfig creates default config
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		Store:    store.NewConfig(),
	}
}
