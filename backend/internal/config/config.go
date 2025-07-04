package config

import (
	"os"
)

type Config struct {
	DatabaseURL string
	JWTSecret   string
	Environment string
	RedisURL    string
}

func Load() *Config {
	return &Config{
		DatabaseURL: getEnv("DB_URL", "postgres://iam:iam@localhost:5432/iam?sslmode=disable"),
		JWTSecret:   getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
		Environment: getEnv("ENV", "development"),
		RedisURL:    getEnv("REDIS_URL", "localhost:6379"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}