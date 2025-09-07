# Rental Property Management Platform

A comprehensive rental property management system that enables real estate investors to store detailed property information, analyze investment opportunities, and collaborate on property evaluation decisions.

## 🚀 Quick Start

### Using Makefile (Recommended)

1. **Development environment with hot reload**
   ```bash
   make dev
   ```
   - Frontend: http://localhost:5173
   - Backend: http://localhost:8080
   - Database: localhost:5432

2. **Production environment**
   ```bash
   make prod
   ```
   - Application: http://localhost
   - API: http://localhost:8080

3. **Run tests**
   ```bash
   make test
   ```

4. **View logs**
   ```bash
   make logs        # Production logs
   make logs-dev    # Development logs
   ```

5. **Stop services**
   ```bash
   make stop
   ```

6. **Clean everything**
   ```bash
   make clean
   ```

### Manual Docker Setup

### Prerequisites
- Go 1.21+
- Node.js 18+
- PostgreSQL 14+
- Docker (optional)

### Backend Setup

1. **Navigate to backend directory**
   ```bash
   cd backend
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```

4. **Start PostgreSQL** (if not using Docker)
   ```bash
   # Create database
   createdb rental_property_mgmt
   ```

5. **Run database migrations**
   ```bash
   go run cmd/server/main.go
   # Migrations run automatically on startup
   ```

6. **Start the API server**
   ```bash
   go run cmd/server/main.go
   ```

   The backend will be available at `http://localhost:8080`

### Frontend Setup

1. **Navigate to frontend directory**
   ```bash
   cd frontend
   ```

2. **Install dependencies**
   ```bash
   npm install
   ```

3. **Start development server**
   ```bash
   npm run dev
   ```

   The frontend will be available at `http://localhost:5173`

### Docker Setup (Recommended)

#### Production Setup
1. **Start all services with Docker Compose**
   ```bash
   docker-compose up -d
   ```

   This will start:
   - PostgreSQL database on port 5432
   - Backend API on port 8080
   - Frontend (Nginx) on port 80

2. **View logs**
   ```bash
   docker-compose logs -f
   ```

3. **Stop services**
   ```bash
   docker-compose down
   ```

#### Development Setup with Hot Reload
1. **Start development environment**
   ```bash
   docker-compose -f docker-compose.dev.yml up -d
   ```

   This will start:
   - PostgreSQL database on port 5432
   - Backend with hot reload on port 8080
   - Frontend dev server on port 5173

2. **Watch logs during development**
   ```bash
   docker-compose -f docker-compose.dev.yml logs -f backend-dev frontend-dev
   ```

#### Alternative: Local Development Setup

## 🐳 Docker Implementation

### Complete Containerization
The project now includes full Docker containerization with both production and development setups:

**Production Stack:**
- ✅ **Backend Dockerfile**: Multi-stage build with Go 1.21 Alpine
- ✅ **Frontend Dockerfile**: Multi-stage build with Node.js 18 Alpine + Nginx
- ✅ **docker-compose.yml**: Production-ready orchestration
- ✅ **Health checks**: PostgreSQL health monitoring
- ✅ **Security**: Nginx configuration with security headers
- ✅ **Optimization**: .dockerignore files for faster builds

**Development Stack:**
- ✅ **Backend Dockerfile.dev**: Hot reload with Air
- ✅ **docker-compose.dev.yml**: Development orchestration
- ✅ **Hot Reload**: Both frontend (Vite) and backend (Air) hot reload
- ✅ **Volume Mounts**: Source code mounted for live development

### Container Architecture

```yaml
# Production (docker-compose.yml)
services:
  postgres:     # PostgreSQL 14 with health checks
  backend:      # Go app (multi-stage build)
  frontend:     # Nginx serving built React app

# Development (docker-compose.dev.yml)  
services:
  postgres:     # Same PostgreSQL setup
  backend-dev:  # Go app with Air hot reload
  frontend-dev: # Vite dev server with hot reload
```

### Network Configuration
- **Frontend**: Nginx proxies `/api/*` requests to backend
- **Backend**: Connects to PostgreSQL via service name
- **Database**: Persistent volume for data storage
- **CORS**: Properly configured for development/production

## 🧪 Testing

### Running Tests

The project follows Test-Driven Development (TDD). Tests must be written first and fail before implementation.

**Backend Tests:**
```bash
cd backend

# Run all tests
go test ./...

# Run contract tests
go test ./tests/contract/ -v

# Run integration tests
go test ./tests/integration/ -v

# Run with coverage
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

**Frontend Tests:**
```bash
cd frontend

# Run unit tests
npm test

# Run end-to-end tests
npm run test:e2e
```

### Test Categories

1. **Contract Tests** - Test API endpoints against OpenAPI specification
2. **Integration Tests** - Test complete user workflows with real database
3. **Unit Tests** - Test individual components and functions

## 📊 Features Implemented

### Core Functionality ✅
- [x] User registration and authentication
- [x] Property CRUD operations
- [x] Financial metrics calculation
  - [x] Cap Rate calculation
  - [x] Cash-on-Cash Return
  - [x] Net Operating Income (NOI)
  - [x] Monthly Mortgage Payment
  - [x] Cash to Close
  - [x] Rent-to-Value Ratio (RTV)
  - [x] Gross Rent Multiplier (GRM)

### Models Implemented ✅
- [x] User model with GORM annotations
- [x] Property model with JSON fields for flexible data
- [x] PropertyValuation for third-party data
- [x] FinancialMetrics for calculated metrics
- [x] Comment model for collaboration
- [x] BuyingBoxCriteria for investment criteria

### API Structure ✅
- [x] RESTful endpoints
- [x] JSON request/response format
- [x] Error handling with proper HTTP status codes
- [x] Basic CORS configuration

### Frontend ✅
- [x] Vanilla JavaScript ES6 modules
- [x] Responsive HTML/CSS design
- [x] API client service
- [x] Basic property management interface
- [x] Authentication flow

## 🚧 In Progress / TODO

### Backend
- [ ] JWT authentication middleware implementation
- [ ] Property valuation endpoints
- [ ] Comment system endpoints
- [ ] Buying criteria comparison logic
- [ ] Property search and filtering
- [ ] Metric recalculation triggers

### Frontend
- [ ] Property detail pages
- [ ] Financial metrics display
- [ ] Property comparison interface
- [ ] Buying criteria management
- [ ] Comment threads
- [ ] Property search/filter UI
- [ ] Charts and graphs for metrics

### Testing
- [ ] Complete contract test suite for all endpoints
- [ ] Integration tests for all quickstart scenarios
- [ ] Frontend unit tests
- [ ] End-to-end tests with Playwright

## 🏗️ Architecture

### Docker Architecture
```
Production Stack:
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Frontend      │    │    Backend      │    │   PostgreSQL    │
│   (Nginx)       │◄───│   (Go Fiber)    │◄───│   Database      │
│   Port: 80      │    │   Port: 8080    │    │   Port: 5432    │
└─────────────────┘    └─────────────────┘    └─────────────────┘

Development Stack:
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Frontend      │    │    Backend      │    │   PostgreSQL    │
│   (Vite Dev)    │◄───│  (Go + Air)     │◄───│   Database      │
│   Port: 5173    │    │   Port: 8080    │    │   Port: 5432    │
│   Hot Reload    │    │   Hot Reload    │    │   Persistent    │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

### Backend Architecture
```
backend/
├── cmd/server/           # Application entry point
├── internal/
│   ├── models/          # Database models (GORM)
│   ├── services/        # Business logic
│   ├── handlers/        # HTTP request handlers
│   └── middleware/      # Authentication, logging, etc.
├── pkg/database/        # Database connection setup
├── migrations/          # Database schema migrations
└── tests/
    ├── contract/        # API contract tests
    └── integration/     # Integration tests
```

### Frontend Architecture
```
frontend/
├── src/
│   ├── components/      # Reusable UI components
│   ├── pages/          # Page-specific logic
│   ├── services/       # API client and utilities
│   └── utils/          # Helper functions
├── public/             # Static assets
└── tests/              # Frontend tests
```

### Database Schema
- **PostgreSQL** with JSONB for flexible property attributes
- **UUID** primary keys for all entities
- **Proper indexes** for query performance
- **Foreign key constraints** with cascade deletes
- **Validation constraints** at database level

## 📈 Financial Calculations

The system automatically calculates key real estate investment metrics:

### Formulas Implemented

**Monthly Mortgage Payment:**
```
P = L[c(1 + c)^n]/[(1 + c)^n - 1]
Where: P = payment, L = loan amount, c = monthly interest rate, n = number of payments
```

**Net Operating Income (NOI):**
```
NOI = (Monthly Rent × 12) - Annual Operating Expenses
```

**Cap Rate:**
```
Cap Rate = (NOI / Purchase Price) × 100
```

**Cash-on-Cash Return:**
```
Annual Cash Flow = NOI - Annual Debt Service
Cash-on-Cash Return = (Annual Cash Flow / Initial Cash Investment) × 100
```

**Rent-to-Value Ratio:**
```
RTV = (Monthly Rent × 12 / Purchase Price) × 100
```

**Gross Rent Multiplier:**
```
GRM = Purchase Price / (Monthly Rent × 12)
```

## 🤝 Contributing

1. **Follow TDD**: Write failing tests before implementing features
2. **Run tests**: Ensure all tests pass before submitting
3. **Follow conventions**: Use established code patterns
4. **Document changes**: Update README and comments

### Development Workflow

1. Write failing contract/integration tests
2. Implement minimum code to make tests pass
3. Refactor while keeping tests green
4. Add unit tests for edge cases
5. Update documentation

## 📝 API Documentation

Once the server is running, API documentation is available at:
- OpenAPI Spec: `http://localhost:8080/docs`
- Health Check: `http://localhost:8080/api/v1/health`

## 🐛 Troubleshooting

### Common Issues

**Container build failures:**
- Check Docker/Podman is running: `docker --version` or `podman --version`
- Clear build cache: `make clean`
- Check network connectivity for base image pulls
- For Podman: ensure registries are configured in `/etc/containers/registries.conf`

**Database connection errors in containers:**
- Verify PostgreSQL container is healthy: `docker-compose ps` 
- Check container logs: `make logs`
- Ensure environment variables match in docker-compose.yml
- Wait for health check to pass before backend starts

**Frontend/Backend communication issues:**
- Check Nginx proxy configuration in nginx.conf
- Verify CORS_ORIGINS environment variable
- Ensure backend is accessible on port 8080
- Check network connectivity between containers

**Development hot reload not working:**
- Ensure volumes are properly mounted in docker-compose.dev.yml
- Check Air configuration in backend/.air.toml
- Verify file permissions on mounted volumes
- Restart development containers: `make stop && make dev`

**Backend won't start:**
- Check Go version: `go version`
- Verify database connection in .env
- Ensure PostgreSQL is running
- Check port 8080 availability

**Frontend won't start:**
- Check Node.js version: `node --version`
- Run `npm install` to install dependencies
- Check port 5173 availability

**Database connection errors:**
- Verify PostgreSQL is running: `pg_isready`
- Check database credentials in .env
- Ensure database exists: `psql -l`

**Tests failing:**
- Run `go mod tidy` to sync dependencies
- Check that database is accessible for integration tests
- Ensure tests are run from correct directory

## 📄 License

This project is part of the rental property management specifications and is intended for educational and development purposes.
