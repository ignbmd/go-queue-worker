package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ignbmd/go-queue-worker/internal/config"
	"github.com/ignbmd/go-queue-worker/internal/db"
	"github.com/ignbmd/go-queue-worker/internal/queue"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	config.LoadEnv()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	pgPool, err := db.NewPostgresPool(ctx, db.FromEnv())
	if err != nil {
		logger.Error("failed to initialize postgres", "error", err)
		os.Exit(1)
	}
	defer pgPool.Close()

	redisClient, err := queue.NewRedisClient(ctx, queue.FromEnv())
	if err != nil {
		logger.Error("failed to initialize redis", "error", err)
		os.Exit(1)
	}
	defer func() {
		if err := redisClient.Close(); err != nil {
			logger.Error("failed to close redis", "error", err)
		}
	}()

	logger.Info("worker started")
	<-ctx.Done()
	logger.Info("worker shutting down")
	time.Sleep(500 * time.Millisecond)
}
