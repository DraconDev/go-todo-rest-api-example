package config

import "os"

// Config holds configuration for the application.
type Config struct {
	Port string
}

// LoadConfig loads configuration from environment variables,
// defaulting the port to 8080 if not provided.
func LoadConfig() (*Config, error) {
	cfg := &Config{
		Port: os.Getenv("PORT"),
	}
	if cfg.Port == "" {
		cfg.Port = "3000" // default port
	}
	return cfg, nil
}