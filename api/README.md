# API

Go Fiber backend for the car rental project. Uses Clean Architecture with config loaded from `.env`.

## Project structure

```
api/
├── cmd/
│   └── app/
│       └── main.go              # Entry point: config, middleware, routes
├── config/
│   └── config.go                # Config struct + LoadConfig() from .env
├── internal/
│   ├── domain/
│   │   ├── entity/              # Domain entities
│   │   └── repository/          # Repository interfaces
│   ├── usecase/                 # Application business logic
│   ├── handler/                 # HTTP handlers and route registration
│   └── infrastructure/
│       ├── persistence/         # Repository implementations (DB adapters)
│       └── middleware/          # CORS, logger, etc.
├── pkg/
│   ├── response/                # Shared JSON response helpers
│   └── utils/                   # General utilities
├── assets/
│   ├── logs/                    # Log files
│   └── static/                  # Static assets
├── docs/                        # Swagger documentation
├── .env.example                 # Example environment variables
├── .air.toml                    # Live-reload config (air)
├── go.mod
└── go.sum
```

## Run

Copy `.env.example` to `.env` and adjust if needed:

```bash
cp .env.example .env
go run ./cmd/app
```

Server listens on `http://localhost:APP_PORT` (default `3000`).

With live-reload:

```bash
air
```

## Configuration

| Variable  | Description       | Default     |
|----------|-------------------|-------------|
| APP_PORT | HTTP server port  | 3000        |
| APP_ENV  | Environment name  | development |

## API endpoints

| Method | Path    | Description   |
|--------|---------|---------------|
| GET    | /health | Health check  |
| GET    | /docs   | Swagger UI    |
