package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config holds application configuration loaded from environment.
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// LoadConfig loads configuration from .env file (if present) and environment variables.
func LoadDatabaseConfig() (*DatabaseConfig, error) {
	_ = godotenv.Load()

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		user = "postgres"
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "postgres"
	}

	database := os.Getenv("DB_DATABASE")
	if database == "" {
		database = "car-rental"
	}

	return &DatabaseConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Database: database,
	}, nil
}
