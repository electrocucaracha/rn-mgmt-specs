# Rental Property Management Platform - Makefile

# Detect container tool
COMPOSE := $(shell which docker-compose 2>/dev/null || which podman-compose 2>/dev/null || echo "podman compose")

.PHONY: help dev prod test clean logs stop build

# Default target
help:
	@echo "Rental Property Management Platform"
	@echo ""
	@echo "Available commands:"
	@echo "  make dev     - Start development environment with hot reload"
	@echo "  make prod    - Start production environment"
	@echo "  make test    - Run all tests"
	@echo "  make build   - Build container images"
	@echo "  make logs    - View logs from running containers"
	@echo "  make stop    - Stop all running containers"
	@echo "  make clean   - Clean up containers and volumes"
	@echo "  make backend - Run backend tests only"
	@echo "  make frontend- Run frontend tests only"
	@echo ""
	@echo "Using container tool: $(COMPOSE)"

# Development environment
dev:
	@echo "Starting development environment..."
	$(COMPOSE) -f docker-compose.dev.yml up -d
	@echo "Development environment started!"
	@echo "Frontend: http://localhost:5173"
	@echo "Backend:  http://localhost:8080"
	@echo "Database: localhost:5432"

# Production environment
prod:
	@echo "Starting production environment..."
	$(COMPOSE) up -d
	@echo "Production environment started!"
	@echo "Application: http://localhost"
	@echo "API:         http://localhost:8080"

# Build container images
build:
	@echo "Building container images..."
	$(COMPOSE) build
	$(COMPOSE) -f docker-compose.dev.yml build

# Run all tests
test: backend frontend

# Backend tests
backend:
	@echo "Running backend tests..."
	cd backend && go test ./... -v

# Frontend tests  
frontend:
	@echo "Running frontend tests..."
	cd frontend && npm test

# View logs
logs:
	@echo "Viewing logs (Ctrl+C to exit)..."
	$(COMPOSE) logs -f

# View development logs
logs-dev:
	@echo "Viewing development logs (Ctrl+C to exit)..."
	$(COMPOSE) -f docker-compose.dev.yml logs -f

# Stop all containers
stop:
	@echo "Stopping all containers..."
	$(COMPOSE) down
	$(COMPOSE) -f docker-compose.dev.yml down

# Clean up everything
clean: stop
	@echo "Cleaning up containers, images, and volumes..."
	$(COMPOSE) down -v --rmi all
	$(COMPOSE) -f docker-compose.dev.yml down -v --rmi all
	podman system prune -f || docker system prune -f || true

# Database operations
db-reset:
	@echo "Resetting database..."
	$(COMPOSE) down postgres
	$(COMPOSE) up -d postgres

# Quick development setup
quick-dev: build dev
	@echo "Quick development setup complete!"

# Quick production setup  
quick-prod: build prod
	@echo "Quick production setup complete!"

# Install local dependencies (for local development)
install:
	@echo "Installing backend dependencies..."
	cd backend && go mod download
	@echo "Installing frontend dependencies..."
	cd frontend && npm install

# Format code
fmt:
	@echo "Formatting Go code..."
	cd backend && go fmt ./...
	@echo "Formatting frontend code..."
	cd frontend && npm run lint:fix

# Security scan
security:
	@echo "Running security scans..."
	cd backend && go mod audit
	cd frontend && npm audit
