package config

import (
	"os"
	"github.com/joho/godotenv"
	"log/slog"
)

// Config holds all configuration for our application
type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	ServerPort string
}

func Load() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		slog.Info("No .env file found")
	}

	cfg := &Config{
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBName:     getEnv("DB_NAME", "devtracker"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}

	return cfg, nil
}

// getEnv is a helper to read an env var with a default fallback
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}