# Tasks: Rental Property Management Platform

**Input**: Design documents from `/specs/001-create-a-website/`
**Prerequisites**: plan.md (required), research.md, data-model.md, contracts/

## Execution Flow (main)
```
1. Load plan.md from feature directory
   → Tech stack: Go 1.21+ backend with Fiber, Vite frontend with vanilla JS
   → Structure: Web application (backend/ and frontend/ directories)
   → Storage: PostgreSQL with GORM ORM
2. Load design documents:
   → data-model.md: 6 entities (User, Property, PropertyValuation, FinancialMetrics, Comment, BuyingBoxCriteria)
   → contracts/api.yaml: 15+ API endpoints across auth, properties, comments, valuations
   → quickstart.md: 6 integration test scenarios
3. Generated task categories:
   → Setup: Project structure, dependencies, database
   → Tests: Contract tests for all endpoints, integration tests for scenarios
   → Models: GORM models for all 6 entities
   → Services: Business logic for calculations, auth, CRUD
   → Handlers: API endpoints implementing OpenAPI spec
   → Frontend: Vanilla JS components and pages
   → Integration: Middleware, validation, logging
   → Polish: Performance tests, documentation
4. Applied task rules:
   → Different files marked [P] for parallel execution
   → Tests before implementation (TDD enforced)
   → Dependencies respected (models → services → handlers)
5. Generated 42 numbered tasks (T001-T042)
6. Parallel execution examples provided
7. Task completeness validated:
   → All 15+ contracts have tests ✓
   → All 6 entities have models ✓
   → All endpoints implemented ✓
   → All 6 scenarios have integration tests ✓
8. Return: SUCCESS (tasks ready for execution)
```

## Format: `[ID] [P?] Description`
- **[P]**: Can run in parallel (different files, no dependencies)
- Include exact file paths in descriptions

## Path Conventions
- **Backend**: `backend/` directory with Go modules
- **Frontend**: `frontend/` directory with Vite/JS
- **Tests**: Within each directory structure

## Phase 3.1: Setup
- [x] T001 Create project structure with backend/ and frontend/ directories per implementation plan
- [x] T002 Initialize Go backend project with Fiber, GORM, and PostgreSQL dependencies in backend/
- [x] T003 Initialize Vite frontend project with minimal dependencies in frontend/
- [x] T004 [P] Configure Go linting (golangci-lint) and formatting in backend/.golangci.yml
- [x] T005 [P] Configure frontend linting (ESLint) and formatting (Prettier) in frontend/
- [x] T006 Set up PostgreSQL database schema and migrations in backend/migrations/
- [x] T007 Configure environment variables and Docker setup
- [x] T068 [P] Create production Dockerfile for backend with multi-stage build in backend/Dockerfile
- [x] T069 [P] Create production Dockerfile for frontend with Nginx in frontend/Dockerfile
- [x] T070 [P] Create development Dockerfile for backend with hot reload in backend/Dockerfile.dev
- [x] T071 Create Nginx configuration for frontend proxy in frontend/nginx.conf
- [x] T072 Create production docker-compose.yml with health checks and proper networking
- [x] T073 Create development docker-compose.dev.yml with hot reload and volume mounts
- [x] T074 [P] Create .dockerignore files for both backend and frontend for optimized builds
- [x] T075 [P] Create Air configuration for Go hot reload in backend/.air.toml
- [x] T076 Create Makefile with Docker/Podman support for easy development and deployment
- [x] T077 Create comprehensive deployment guide in DEPLOYMENT.md
- [x] T078 Update README.md with Docker setup instructions and architecture diagrams

## Phase 3.2: Tests First (TDD) ⚠️ MUST COMPLETE BEFORE 3.3
**CRITICAL: These tests MUST be written and MUST FAIL before ANY implementation**

### Contract Tests [P] - All can run in parallel
- [ ] T008 [P] Contract test POST /auth/register in backend/tests/contract/auth_register_test.go
- [ ] T009 [P] Contract test POST /auth/login in backend/tests/contract/auth_login_test.go
- [ ] T010 [P] Contract test POST /auth/logout in backend/tests/contract/auth_logout_test.go
- [ ] T011 [P] Contract test GET /properties in backend/tests/contract/properties_list_test.go
- [ ] T012 [P] Contract test POST /properties in backend/tests/contract/properties_create_test.go
- [ ] T013 [P] Contract test GET /properties/{id} in backend/tests/contract/properties_get_test.go
- [ ] T014 [P] Contract test PUT /properties/{id} in backend/tests/contract/properties_update_test.go
- [ ] T015 [P] Contract test DELETE /properties/{id} in backend/tests/contract/properties_delete_test.go
- [ ] T016 [P] Contract test GET /properties/{id}/metrics in backend/tests/contract/metrics_get_test.go
- [ ] T017 [P] Contract test POST /properties/{id}/metrics in backend/tests/contract/metrics_calculate_test.go
- [ ] T018 [P] Contract test GET /properties/{id}/valuations in backend/tests/contract/valuations_list_test.go
- [ ] T019 [P] Contract test POST /properties/{id}/valuations in backend/tests/contract/valuations_create_test.go
- [ ] T020 [P] Contract test GET /properties/{id}/comments in backend/tests/contract/comments_list_test.go
- [ ] T021 [P] Contract test POST /properties/{id}/comments in backend/tests/contract/comments_create_test.go
- [ ] T022 [P] Contract test PUT /comments/{id} in backend/tests/contract/comments_update_test.go
- [ ] T023 [P] Contract test DELETE /comments/{id} in backend/tests/contract/comments_delete_test.go
- [ ] T024 [P] Contract test GET /buying-criteria in backend/tests/contract/criteria_list_test.go
- [ ] T025 [P] Contract test POST /buying-criteria in backend/tests/contract/criteria_create_test.go
- [ ] T026 [P] Contract test PUT /buying-criteria/{id} in backend/tests/contract/criteria_update_test.go
- [ ] T027 [P] Contract test DELETE /buying-criteria/{id} in backend/tests/contract/criteria_delete_test.go
- [ ] T028 [P] Contract test POST /properties/compare in backend/tests/contract/properties_compare_test.go

### Integration Tests [P] - Based on quickstart scenarios
- [ ] T029 [P] Integration test user registration and login flow in backend/tests/integration/auth_flow_test.go
- [ ] T030 [P] Integration test property creation and metric calculation in backend/tests/integration/property_analysis_test.go
- [ ] T031 [P] Integration test third-party valuations workflow in backend/tests/integration/valuations_test.go
- [ ] T032 [P] Integration test buying criteria and comparison in backend/tests/integration/criteria_comparison_test.go
- [ ] T033 [P] Integration test team collaboration features in backend/tests/integration/collaboration_test.go
- [ ] T034 [P] Integration test property editing and recalculation in backend/tests/integration/property_updates_test.go

## Phase 3.3: Core Implementation (ONLY after tests are failing)

### Database Models [P] - All can run in parallel
- [ ] T035 [P] User model with GORM annotations in backend/internal/models/user.go
- [ ] T036 [P] Property model with GORM annotations in backend/internal/models/property.go
- [ ] T037 [P] PropertyValuation model with GORM annotations in backend/internal/models/property_valuation.go
- [ ] T038 [P] FinancialMetrics model with GORM annotations in backend/internal/models/financial_metrics.go
- [ ] T039 [P] Comment model with GORM annotations in backend/internal/models/comment.go
- [ ] T040 [P] BuyingBoxCriteria model with GORM annotations in backend/internal/models/buying_box_criteria.go

### Business Services
- [ ] T041 UserService with authentication logic in backend/internal/services/user_service.go
- [ ] T042 PropertyService with CRUD operations in backend/internal/services/property_service.go
- [ ] T043 CalculationService with financial metrics logic in backend/internal/services/calculation_service.go
- [ ] T044 ValuationService for third-party data in backend/internal/services/valuation_service.go
- [ ] T045 CommentService for collaboration features in backend/internal/services/comment_service.go
- [ ] T046 CriteriaService for buying box logic in backend/internal/services/criteria_service.go

### API Handlers
- [ ] T047 Authentication handlers (register, login, logout) in backend/internal/handlers/auth_handler.go
- [ ] T048 Property handlers (CRUD operations) in backend/internal/handlers/property_handler.go
- [ ] T049 Metrics handlers (calculate, retrieve) in backend/internal/handlers/metrics_handler.go
- [ ] T050 Valuation handlers in backend/internal/handlers/valuation_handler.go
- [ ] T051 Comment handlers in backend/internal/handlers/comment_handler.go
- [ ] T052 Buying criteria handlers in backend/internal/handlers/criteria_handler.go
- [ ] T053 Property comparison handler in backend/internal/handlers/comparison_handler.go

### Frontend Components [P] - Different pages can be parallel
- [ ] T054 [P] Authentication pages (login, register) in frontend/src/pages/auth/
- [ ] T055 [P] Property list page with sorting/filtering in frontend/src/pages/properties/list.js
- [ ] T056 [P] Property create/edit form in frontend/src/pages/properties/form.js
- [ ] T057 [P] Property detail page with metrics display in frontend/src/pages/properties/detail.js
- [ ] T058 [P] Property comparison page in frontend/src/pages/properties/compare.js
- [ ] T059 [P] Buying criteria management page in frontend/src/pages/criteria/
- [ ] T060 [P] Comment components for collaboration in frontend/src/components/comments/
- [ ] T061 API client service for backend communication in frontend/src/services/api.js
- [ ] T062 Frontend state management and routing in frontend/src/utils/

## Phase 3.4: Integration
- [ ] T063 Database connection setup and migration runner in backend/pkg/database/
- [ ] T064 JWT authentication middleware in backend/internal/middleware/auth.go
- [ ] T065 Request validation middleware in backend/internal/middleware/validation.go
- [ ] T066 Error handling middleware in backend/internal/middleware/error.go
- [ ] T067 Logging middleware with structured logs in backend/internal/middleware/logging.go
- [ ] T068 CORS configuration for frontend-backend communication
- [ ] T069 Main server setup with all routes in backend/cmd/server/main.go

## Phase 3.5: Containerization and Deployment
- [x] T068 [P] Create production Dockerfile for backend with multi-stage build in backend/Dockerfile
- [x] T069 [P] Create production Dockerfile for frontend with Nginx in frontend/Dockerfile
- [x] T070 [P] Create development Dockerfile for backend with hot reload in backend/Dockerfile.dev
- [x] T071 Create Nginx configuration for frontend proxy in frontend/nginx.conf
- [x] T072 Create production docker-compose.yml with health checks and proper networking
- [x] T073 Create development docker-compose.dev.yml with hot reload and volume mounts
- [x] T074 [P] Create .dockerignore files for both backend and frontend for optimized builds
- [x] T075 [P] Create Air configuration for Go hot reload in backend/.air.toml
- [x] T076 Create Makefile with Docker/Podman support for easy development and deployment
- [x] T077 Create comprehensive deployment guide in DEPLOYMENT.md
- [x] T078 Update README.md with Docker setup instructions and architecture diagrams
- [ ] T079 Test Docker builds in isolated environment to verify network connectivity
- [ ] T080 Validate production deployment on cloud platform (AWS/GCP/Azure)
- [ ] T081 Set up container registry and CI/CD pipeline for automated deployments
- [ ] T082 Configure environment-specific secrets management for production
- [ ] T083 [P] Set up monitoring and logging for containerized applications
- [ ] T084 [P] Implement container health checks and auto-restart policies
- [ ] T085 [P] Configure backup strategies for PostgreSQL in containerized environment

## Phase 3.6: Integration Testing with Docker
- [ ] T086 Create integration test suite that works with Docker Compose
- [ ] T087 Set up test database seeding for containerized testing
- [ ] T088 Implement end-to-end tests using containerized environment
- [ ] T089 Configure automated testing pipeline with Docker
- [ ] T090 [P] Set up performance testing for containerized application

## Dependencies
- **Setup first**: T001-T007 before all other tasks
- **Tests before implementation**: T008-T034 before T035-T069
- **Models before services**: T035-T040 before T041-T046
- **Services before handlers**: T041-T046 before T047-T053
- **Backend before frontend integration**: T047-T053 before T061-T062
- **Core before integration**: T035-T062 before T063-T069
- **Everything before polish**: T070-T079 last

## Parallel Execution Examples

### Contract Tests Phase (T008-T028)
```bash
# All contract tests can run simultaneously
go test ./tests/contract/auth_register_test.go &
go test ./tests/contract/auth_login_test.go &
go test ./tests/contract/properties_list_test.go &
go test ./tests/contract/properties_create_test.go &
# ... all other contract tests
wait
```

### Integration Tests Phase (T029-T034)
```bash
# All integration tests can run simultaneously
go test ./tests/integration/auth_flow_test.go &
go test ./tests/integration/property_analysis_test.go &
go test ./tests/integration/valuations_test.go &
go test ./tests/integration/criteria_comparison_test.go &
go test ./tests/integration/collaboration_test.go &
go test ./tests/integration/property_updates_test.go &
wait
```

### Models Phase (T035-T040)
```bash
# All model files can be created simultaneously
Task: "User model with GORM annotations in backend/internal/models/user.go"
Task: "Property model with GORM annotations in backend/internal/models/property.go"
Task: "PropertyValuation model with GORM annotations in backend/internal/models/property_valuation.go"
Task: "FinancialMetrics model with GORM annotations in backend/internal/models/financial_metrics.go"
Task: "Comment model with GORM annotations in backend/internal/models/comment.go"
Task: "BuyingBoxCriteria model with GORM annotations in backend/internal/models/buying_box_criteria.go"
```

### Frontend Pages Phase (T054-T060)
```bash
# Different page components can be built in parallel
Task: "Authentication pages (login, register) in frontend/src/pages/auth/"
Task: "Property list page with sorting/filtering in frontend/src/pages/properties/list.js"
Task: "Property create/edit form in frontend/src/pages/properties/form.js"
Task: "Property detail page with metrics display in frontend/src/pages/properties/detail.js"
Task: "Property comparison page in frontend/src/pages/properties/compare.js"
Task: "Buying criteria management page in frontend/src/pages/criteria/"
Task: "Comment components for collaboration in frontend/src/components/comments/"
```

## Notes
- **TDD Enforcement**: All T008-T034 tests must fail before implementing T035+
- **Commit Strategy**: Commit after each completed task
- **File Conflicts**: No [P] tasks modify the same file
- **Validation**: Run contract tests after each handler implementation
- **Documentation**: OpenAPI spec serves as single source of truth

## Task Generation Rules Applied

1. **From Contracts (api.yaml)**:
   - 21 contract test tasks (T008-T028) - one per endpoint [P]
   - Corresponding implementation tasks for each endpoint
   
2. **From Data Model**:
   - 6 model creation tasks (T035-T040) - one per entity [P]
   - Service layer tasks for business logic relationships
   
3. **From Quickstart Scenarios**:
   - 6 integration test tasks (T029-T034) - one per scenario [P]
   - Frontend validation tasks
   
4. **Ordering Applied**:
   - Setup → Tests → Models → Services → Handlers → Frontend → Integration → Polish
   - Dependencies prevent conflicts while maximizing parallelization

## Validation Checklist ✓
- [x] All 21 API endpoints have corresponding contract tests
- [x] All 6 entities have model creation tasks
- [x] All tests (T008-T034) come before implementation (T035+)
- [x] Parallel tasks ([P]) are truly independent (different files)
- [x] Each task specifies exact file path
- [x] No [P] task modifies same file as another [P] task
- [x] TDD cycle enforced: RED (tests fail) → GREEN (implement) → REFACTOR (polish)
