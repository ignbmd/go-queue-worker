package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type HealthHandler struct {
	DB    *pgxpool.Pool
	Redis *redis.Client
}

type healthResponse struct {
	Status string            `json:"status"`
	Checks map[string]string `json:"checks"`
}

func (h HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	checks := map[string]string{}
	status := "ok"

	if h.DB != nil {
		if err := h.DB.Ping(ctx); err != nil {
			checks["postgres"] = "unhealthy"
			status = "degraded"
		} else {
			checks["postgres"] = "ok"
		}
	}

	if h.Redis != nil {
		if err := h.Redis.Ping(ctx).Err(); err != nil {
			checks["redis"] = "unhealthy"
			status = "degraded"
		} else {
			checks["redis"] = "ok"
		}
	}

	response := healthResponse{Status: status, Checks: checks}
	w.Header().Set("Content-Type", "application/json")
	if status != "ok" {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
	_ = json.NewEncoder(w).Encode(response)
}
