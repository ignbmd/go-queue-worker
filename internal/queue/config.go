package queue

import (
	"os"
	"strconv"
	"time"
)

// FromEnv builds RedisConfig from environment variables.
func FromEnv() RedisConfig {
	return RedisConfig{
		Host:         getenv("REDIS_HOST", "localhost"),
		Port:         getenv("REDIS_PORT", "6379"),
		Password:     os.Getenv("REDIS_PASSWORD"),
		DB:           getenvInt("REDIS_DB", 0),
		DialTimeout:  getenvDuration("REDIS_DIAL_TIMEOUT", 5*time.Second),
		ReadTimeout:  getenvDuration("REDIS_READ_TIMEOUT", 3*time.Second),
		WriteTimeout: getenvDuration("REDIS_WRITE_TIMEOUT", 3*time.Second),
		PoolSize:     getenvInt("REDIS_POOL_SIZE", 10),
		MinIdleConns: getenvInt("REDIS_MIN_IDLE_CONNS", 2),
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
