package config

import (
	"os"
)

type Config struct {
	Port               string
	Auth0Domain        string
	Auth0Audience      string
	CORSAllowedOrigins string
}

func Load() *Config {
	return &Config{
		Port:               getEnv("PORT", "8080"),
		Auth0Domain:        getEnv("AUTH0_DOMAIN", ""),
		Auth0Audience:      getEnv("AUTH0_AUDIENCE", ""),
		CORSAllowedOrigins: getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:3000"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
