package config

type Config struct {
	BindAddr string `json:"bind_addr"`
	LogLevel string `json:"log_level"`
	DatabaseURL string `json:"database_url"`
	SessionKey string `json:"session_key"`
}

// NewConfig() returns default config
func NewConfig() Config {
	return Config{
		BindAddr: ":7070",
		LogLevel: "debug",
		DatabaseURL: DATABASE_URL,
	}
}

const DATABASE_URL = "postgres://postgres:postgres@0.0.0.0:5432/postgres"