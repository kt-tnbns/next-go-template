package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds application configuration loaded from environment.
type Config struct {
	AppPort string
	AppEnv  string
}

// LoadConfig loads configuration from .env file (if present) and environment variables.
func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	return &Config{
		AppPort: port,
		AppEnv:  env,
	}, nil
}

// Port returns the port as int for use with Listen.
func (c *Config) Port() (int, error) {
	return strconv.Atoi(c.AppPort)
}
