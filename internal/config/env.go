package config

import (
	"errors"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from a .env file if present.
// It is safe to call multiple times.
func LoadEnv() {
	paths := envPaths()
	if err := godotenv.Load(paths...); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			slog.Debug("no .env file loaded", "searched", paths)
			return
		}
		slog.Debug("failed to load .env file", "error", err, "searched", paths)
	}
}

func envPaths() []string {
	wd, err := os.Getwd()
	if err != nil {
		return []string{".env"}
	}

	return []string{
		filepath.Join(wd, ".env"),
		filepath.Join(wd, "..", ".env"),
		filepath.Join(wd, "..", "..", ".env"),
	}
}
