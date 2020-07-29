package store

//Config type for configuring store
type Config struct {
	DatabaseURL string `toml:"database_url"`
}

//NewConfig creates default config for store
func NewConfig() *Config {
	return &Config{}
}
