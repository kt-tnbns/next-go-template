# Car Rental API - common commands (like package.json scripts)
# Usage: make <target>

.PHONY: build run dev test tidy deps swagger docker-up docker-down clean

# Build the binary
build:
	go build -o ./cmd/app/tmp/main ./cmd/app

# Run the app (no hot reload)
run: build
	./cmd/app/tmp/main

# Dev with hot reload (Air)
dev:
	cd cmd/app && air

# Run tests
test:
	go test ./...

# Download/update dependencies
deps:
	go mod download
	go mod tidy

# Generate Swagger docs (requires: go install github.com/swaggo/swag/cmd/swag@latest)
gen-docs:
	swag init -g cmd/app/main.go -o docs

# Start dependencies (e.g. Postgres) via Docker
docker-up:
	docker compose up -d

# Stop Docker services
docker-down:
	docker compose down

# Clean build artifacts
clean:
	rm -f ./cmd/app/tmp/main
	rm -f ./cmd/app/tmp/build-errors.log
