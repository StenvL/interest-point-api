package store

// Config is a struct for configuring store.
type Config struct {
	DatabaseURL string `toml:"database_url"`
}

// NewConfig creates config with default values.
func NewConfig() *Config {
	return &Config{}
}
