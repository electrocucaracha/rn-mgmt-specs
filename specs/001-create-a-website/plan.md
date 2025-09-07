# Implementation Plan: Rental Property Management Platform

**Branch**: `001-create-a-website` | **Date**: September 4, 2025 | **Spec**: [spec.md](./spec.md)
**Input**: Feature specification from `/specs/001-create-a-website/spec.md`

## Execution Flow (/plan command scope)
```
1. Load feature spec from Input path
   → If not found: ERROR "No feature spec at {path}"
2. Fill Technical Context (scan for NEEDS CLARIFICATION)
   → Detect Project Type from context (web=frontend+backend, mobile=app+api)
   → Set Structure Decision based on project type
3. Evaluate Constitution Check section below
   → If violations exist: Document in Complexity Tracking
   → If no justification possible: ERROR "Simplify approach first"
   → Update Progress Tracking: Initial Constitution Check
4. Execute Phase 0 → research.md
   → If NEEDS CLARIFICATION remain: ERROR "Resolve unknowns"
5. Execute Phase 1 → contracts, data-model.md, quickstart.md, agent-specific template file (e.g., `CLAUDE.md` for Claude Code, `.github/copilot-instructions.md` for GitHub Copilot, or `GEMINI.md` for Gemini CLI).
6. Re-evaluate Constitution Check section
   → If new violations: Refactor design, return to Phase 1
   → Update Progress Tracking: Post-Design Constitution Check
7. Plan Phase 2 → Describe task generation approach (DO NOT create tasks.md)
8. STOP - Ready for /tasks command
```

**IMPORTANT**: The /plan command STOPS at step 7. Phases 2-4 are executed by other commands:
- Phase 2: /tasks command creates tasks.md
- Phase 3-4: Implementation execution (manual or via tools)

## Summary
Rental Property Management Platform that enables real estate investors to store comprehensive property data, automatically calculate investment metrics (NOI, Cap Rate, CoC Return, etc.), and collaborate on property evaluation decisions. Technical approach: Vite frontend with vanilla HTML/CSS/JavaScript, Go backend with PostgreSQL database and minimal dependencies.

## Technical Context
**Language/Version**: Go 1.21+ (backend), Vanilla JavaScript ES2022 (frontend)  
**Primary Dependencies**: Vite (build tool), Go Fiber (web framework), GORM (ORM)  
**Storage**: PostgreSQL 14+  
**Testing**: Go testing package, Testify, Playwright (E2E)  
**Target Platform**: Linux/macOS/Windows server, Modern browsers (Chrome 90+, Firefox 88+, Safari 14+)
**Project Type**: web (frontend + backend)  
**Performance Goals**: <500ms API response time, 100+ concurrent users  
**Constraints**: Minimal dependencies, vanilla JS preferred, responsive design  
**Scale/Scope**: Small to medium real estate investors (10-1000 properties per user)

## Constitution Check
*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

**Simplicity**:
- Projects: 2 (backend API, frontend client)
- Using framework directly? (Go Fiber directly, no wrapper classes)
- Single data model? (Yes, PostgreSQL schema maps to Go structs)
- Avoiding patterns? (Direct ORM usage, no Repository pattern initially)

**Architecture**:
- EVERY feature as library? (Backend services as packages, frontend modules)
- Libraries listed: 
  - property-service (CRUD operations)
  - calculation-service (financial metrics)
  - user-service (authentication)
  - comment-service (collaboration)
- CLI per library: (Go services expose CLI commands with --help/--version/--format)
- Library docs: llms.txt format planned? (Yes, auto-generated from code comments)

**Testing (NON-NEGOTIABLE)**:
- RED-GREEN-Refactor cycle enforced? (Yes, tests written first)
- Git commits show tests before implementation? (Required)
- Order: Contract→Integration→E2E→Unit strictly followed? (Yes)
- Real dependencies used? (PostgreSQL testcontainers, real database)
- Integration tests for: new libraries, contract changes, shared schemas? (Yes)
- FORBIDDEN: Implementation before test, skipping RED phase (Enforced)

**Observability**:
- Structured logging included? (Yes, JSON logging via slog)
- Frontend logs → backend? (Yes, via POST /api/logs endpoint)
- Error context sufficient? (Request IDs, stack traces)

**Versioning**:
- Version number assigned? (v1.0.0)
- BUILD increments on every change? (Yes, automated)
- Breaking changes handled? (API versioning, migration scripts)

## Project Structure

### Documentation (this feature)
```
specs/001-create-a-website/
├── plan.md              # This file (/plan command output)
├── research.md          # Phase 0 output (/plan command)
├── data-model.md        # Phase 1 output (/plan command)
├── quickstart.md        # Phase 1 output (/plan command)
├── contracts/           # Phase 1 output (/plan command)
└── tasks.md             # Phase 2 output (/tasks command - NOT created by /plan)
```

### Source Code (repository root)
```
# Option 2: Web application (frontend + backend detected)
backend/
├── cmd/
│   └── server/          # Main application entry point
├── internal/
│   ├── models/          # Database models (GORM)
│   ├── services/        # Business logic libraries
│   ├── handlers/        # HTTP handlers
│   └── middleware/      # Authentication, logging
├── pkg/
│   ├── database/        # Database connection
│   └── utils/           # Shared utilities
├── migrations/          # Database schema migrations
└── tests/
    ├── contract/        # API contract tests
    ├── integration/     # Service integration tests
    └── unit/           # Unit tests

frontend/
├── src/
│   ├── components/      # Vanilla JS components
│   ├── pages/          # Page controllers
│   ├── services/       # API client services
│   └── utils/          # Frontend utilities
├── public/             # Static assets
├── tests/
│   ├── e2e/           # Playwright tests
│   └── unit/          # Frontend unit tests
└── vite.config.js     # Build configuration
```

**Structure Decision**: Option 2 (Web application) - Frontend + Backend detected from requirements
api/
└── [same as backend above]

ios/ or android/
└── [platform-specific structure]
```

**Structure Decision**: [DEFAULT to Option 1 unless Technical Context indicates web/mobile app]

## Phase 0: Outline & Research
1. **Extract unknowns from Technical Context** above:
   - For each NEEDS CLARIFICATION → research task
   - For each dependency → best practices task
   - For each integration → patterns task

2. **Generate and dispatch research agents**:
   ```
   For each unknown in Technical Context:
     Task: "Research {unknown} for {feature context}"
   For each technology choice:
     Task: "Find best practices for {tech} in {domain}"
   ```

3. **Consolidate findings** in `research.md` using format:
   - Decision: [what was chosen]
   - Rationale: [why chosen]
   - Alternatives considered: [what else evaluated]

**Output**: research.md with all NEEDS CLARIFICATION resolved

## Phase 1: Design & Contracts
*Prerequisites: research.md complete*

1. **Extract entities from feature spec** → `data-model.md`:
   - Entity name, fields, relationships
   - Validation rules from requirements
   - State transitions if applicable

2. **Generate API contracts** from functional requirements:
   - For each user action → endpoint
   - Use standard REST/GraphQL patterns
   - Output OpenAPI/GraphQL schema to `/contracts/`

3. **Generate contract tests** from contracts:
   - One test file per endpoint
   - Assert request/response schemas
   - Tests must fail (no implementation yet)

4. **Extract test scenarios** from user stories:
   - Each story → integration test scenario
   - Quickstart test = story validation steps

5. **Update agent file incrementally** (O(1) operation):
   - Run `/scripts/update-agent-context.sh [claude|gemini|copilot]` for your AI assistant
   - If exists: Add only NEW tech from current plan
   - Preserve manual additions between markers
   - Update recent changes (keep last 3)
   - Keep under 150 lines for token efficiency
   - Output to repository root

**Output**: data-model.md, /contracts/*, failing tests, quickstart.md, agent-specific file

## Phase 2: Task Planning Approach
*This section describes what the /tasks command will do - DO NOT execute during /plan*

**Task Generation Strategy**:
- Load `/templates/tasks-template.md` as base structure
- Generate contract test tasks from OpenAPI specification (api.yaml)
- Generate model creation tasks from data-model.md entities
- Generate service layer tasks for business logic (property calculations, user management)
- Generate handler tasks for API endpoints
- Generate frontend component tasks for UI
- Generate integration test tasks from quickstart scenarios

**Ordering Strategy**:
- **Phase 1**: Database and models (foundation)
  - Database migrations and schema setup [P]
  - GORM model definitions [P]
  - Basic CRUD repository patterns [P]

- **Phase 2**: Backend services and API
  - Authentication service and JWT middleware
  - Property service with calculation logic
  - Comment service for collaboration
  - User service for profile management
  - API handlers implementing OpenAPI specification

- **Phase 3**: Frontend components and pages
  - Authentication pages (login, register) [P]
  - Property management pages (list, create, edit, detail) [P]
  - Collaboration features (comments, sharing) [P]
  - Investment analysis pages (metrics, comparisons) [P]

- **Phase 4**: Integration and testing
  - Contract tests for all API endpoints
  - Integration tests for service interactions
  - End-to-end tests using Playwright
  - Performance validation tests

**TDD Ordering**: Tests before implementation for every component
**Dependency Ordering**: Models → Services → Handlers → Frontend
**Mark [P] for parallel execution**: Independent components that don't depend on each other

## Implementation Progress

### ✅ Completed Phases

**Phase 1: Project Foundation (COMPLETED)**
- Project structure with backend/ and frontend/ directories
- Go backend with Fiber framework, GORM ORM, PostgreSQL driver
- Vite frontend with minimal dependencies
- Environment configuration and linting setup

**Phase 2: Database and Models (COMPLETED)**
- Complete PostgreSQL database schema design
- GORM models with proper relationships and validations
- Database migration system
- Environment variables configuration

**Phase 3.1: Setup (COMPLETED)**
- Project initialization with all required dependencies
- Database connection and migration setup
- TDD test structure (currently in RED phase)

**Phase 3.2: Backend Implementation (COMPLETED)**
- Complete business models with GORM annotations
- Financial calculation engine implementation
- Service layer architecture
- Test-driven development structure

**Phase 3.3: Frontend Implementation (COMPLETED)**
- API client implementation
- Responsive UI components
- Authentication flow structure
- Modern JavaScript architecture

**Phase 3.5: Containerization and Deployment (COMPLETED)**
- Production Dockerfiles for backend and frontend
- Development Docker setup with hot reload
- Docker Compose configurations for dev and prod
- Nginx reverse proxy configuration
- Comprehensive deployment documentation
- Makefile with Docker/Podman support

### 🚧 Current Phase: Integration and Testing

**Phase 3.4: API Implementation (IN PROGRESS)**
- Currently in TDD RED phase (tests failing as expected)
- Need to implement route handlers to move to GREEN phase
- Authentication middleware pending
- API endpoints pending implementation

### 📋 Next Steps

**Immediate (Next 1-2 sessions):**
1. Test Docker builds once network connectivity allows
2. Implement API route handlers to satisfy TDD tests
3. Add authentication middleware
4. Test full application workflow

**Near Term (Next 3-5 sessions):**
1. Complete integration testing with containerized environment
2. Set up CI/CD pipeline
3. Deploy to cloud platform for validation
4. Implement monitoring and logging

**Estimated Output**: 35-40 numbered, ordered tasks covering:
- 8 database/model tasks
- 12 backend service/API tasks  
- 10 frontend component tasks
- 8 testing and integration tasks

**IMPORTANT**: This phase is executed by the /tasks command, NOT by /plan

## Phase 3+: Future Implementation
*These phases are beyond the scope of the /plan command*

**Phase 3**: Task execution (/tasks command creates tasks.md)  
**Phase 4**: Implementation (execute tasks.md following constitutional principles)  
**Phase 5**: Validation (run tests, execute quickstart.md, performance validation)

## Complexity Tracking
*Fill ONLY if Constitution Check has violations that must be justified*

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| [e.g., 4th project] | [current need] | [why 3 projects insufficient] |
| [e.g., Repository pattern] | [specific problem] | [why direct DB access insufficient] |


## Progress Tracking
*This checklist is updated during execution flow*

**Phase Status**:
- [x] Phase 0: Research complete (/plan command)
- [x] Phase 1: Design complete (/plan command)
- [x] Phase 2: Task planning complete (/plan command - describe approach only)
- [x] Phase 3: Tasks generated (/tasks command)
- [ ] Phase 4: Implementation complete
- [ ] Phase 5: Validation passed

**Gate Status**:
- [x] Initial Constitution Check: PASS
- [x] Post-Design Constitution Check: PASS  
- [x] All NEEDS CLARIFICATION resolved
- [x] Complexity deviations documented (none required)

---
*Based on Constitution v2.1.1 - See `/memory/constitution.md`*