package db

import (
	"os"
	"strconv"
	"time"
)

// FromEnv builds PostgresConfig from environment variables.
func FromEnv() PostgresConfig {
	return PostgresConfig{
		Host:              getenv("POSTGRES_HOST", "localhost"),
		Port:              getenv("POSTGRES_PORT", "5432"),
		User:              getenv("POSTGRES_USER", "postgres"),
		Password:          getenv("POSTGRES_PASSWORD", "postgres"),
		Database:          getenv("POSTGRES_DB", "postgres"),
		SSLMode:           getenv("POSTGRES_SSLMODE", "disable"),
		MaxConns:          int32(getenvInt("POSTGRES_MAX_CONNS", 10)),
		MinConns:          int32(getenvInt("POSTGRES_MIN_CONNS", 2)),
		MaxConnIdleTime:   getenvDuration("POSTGRES_MAX_CONN_IDLE_TIME", 5*time.Minute),
		MaxConnLifetime:   getenvDuration("POSTGRES_MAX_CONN_LIFETIME", 60*time.Minute),
		HealthCheckPeriod: getenvDuration("POSTGRES_HEALTH_CHECK_PERIOD", 1*time.Minute),
	}
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getenvInt(key string, fallback int) int {
	if value := os.Getenv(key); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil {
			return parsed
		}
	}
	return fallback
}

func getenvDuration(key string, fallback time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if parsed, err := time.ParseDuration(value); err == nil {
			return parsed
		}
	}
	return fallback
}
