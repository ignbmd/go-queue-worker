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

## Running services

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
