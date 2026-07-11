# Payment Checkout Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build an authenticated checkout endpoint that validates and transactionally persists orders with integer monetary values.

**Architecture:** Add a feature-local orders module with domain, usecase, PostgreSQL repository, and Gin handler packages. Wire it through bootstrap and the existing JWT/error middleware, and make its migration module-owned.

**Tech Stack:** Go 1.25.5, Gin 1.10, GORM 1.30, PostgreSQL, google/uuid, standard `testing` and `httptest`.

## Global Constraints

- All implementation remains inside `services/main-backend`.
- Monetary values use positive `int64` minor units; no floating-point arithmetic.
- Customer identity, order ID, timestamp, total, and status are server-controlled.
- The initial status is `PENDING`.
- No UI or frontend files are read or modified.

---

### Task 1: Domain and usecase

**Files:**
- Create: `internal/orders/domain/order.go`
- Create: `internal/orders/usecase/ports.go`
- Create: `internal/orders/usecase/service.go`
- Test: `internal/orders/usecase/service_test.go`

**Interfaces:**
- Produces: `Repository.Create(context.Context, *domain.Order) error`
- Produces: `Service.Create(context.Context, CreateInput) (*domain.Order, error)`

- [ ] Write service tests first for successful total calculation, defaults, validation, duplicate products, overflow, and repository failure.
- [ ] Run `go test ./internal/orders/usecase` and confirm failure because the feature does not exist.
- [ ] Add the domain types, repository port, and minimal service implementation using checked `int64` multiplication/addition.
- [ ] Run `go test ./internal/orders/usecase` and confirm all tests pass.

### Task 2: PostgreSQL persistence

**Files:**
- Create: `internal/orders/repository/postgres/model.go`
- Create: `internal/orders/repository/postgres/repository.go`
- Create: `internal/orders/repository/postgres/migrate.go`
- Test: `internal/orders/repository/postgres/repository_integration_test.go`

**Interfaces:**
- Consumes: `domain.Order` and `usecase.Repository`.
- Produces: `postgres.NewRepository(*gorm.DB) *Repository` and `postgres.Migrate(*gorm.DB) error`.

- [ ] Write an integration test first for persisted order/items and transaction rollback.
- [ ] Run the integration test and confirm failure because the repository is missing.
- [ ] Add feature-owned GORM models and a transaction that creates the order then all items.
- [ ] Run unit compilation and the integration test when Docker is available.

### Task 3: HTTP contract and application wiring

**Files:**
- Create: `internal/orders/handler/dto.go`
- Create: `internal/orders/handler/handler.go`
- Test: `internal/orders/handler/handler_test.go`
- Modify: `internal/bootstrap/bootstrap.go`
- Modify: `internal/platform/testutil/integration/helpers.go`
- Modify: `internal/platform/database/database.go`
- Delete: `internal/platform/database/model/order.go`

**Interfaces:**
- Consumes: `Service.Create`, `middleware.Auth`, and `middleware.UserID`.
- Produces: authenticated `POST /api/v1/checkout` returning an order response with status 201.

- [ ] Write handler tests first for success, unauthorized access, malformed JSON, invalid values, and service failures.
- [ ] Run `go test ./internal/orders/handler` and confirm failure because the handler is missing.
- [ ] Add DTO mapping and the Gin handler, then wire repository/service/handler and the module migration in bootstrap.
- [ ] Remove the obsolete platform order model and its global AutoMigrate entry.
- [ ] Run `gofmt`, targeted order tests, `go vet ./...`, and `go test ./...`; distinguish any pre-existing contract-spec failures.
