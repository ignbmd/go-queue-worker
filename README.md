# Go Queue Worker

Boilerplate for a queue-based Go service with API, worker, and scheduler processes.

## Project layout

```
/cmd
  /api
  /worker
  /scheduler
/internal
  /db
  /queue
  /jobs
  /handlers
/migrations
/docker
  docker-compose.yml
README.md
```

## Setup

1. Copy the provided `.env` and adjust values as needed.
2. Start infrastructure services:

```bash
docker compose -f docker/docker-compose.yml up -d
```

The services load `.env` automatically on startup. If you run the API/worker/scheduler on your host, set `POSTGRES_HOST=localhost` and `REDIS_HOST=localhost`. If you run them inside the Docker Compose network, use `POSTGRES_HOST=postgres` and `REDIS_HOST=redis` to match the service names.

## Running services

Run the commands from the repository root so the `.env` file is discovered automatically.

```bash
go run ./cmd/api
```

```bash
go run ./cmd/worker
```

```bash
go run ./cmd/scheduler
```

## Health check

```bash
curl http://localhost:8080/healthz
```
