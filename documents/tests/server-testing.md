# Server testing guide

**Repository:** [apps/server](https://github.com/quanngynx/shopwise/tree/main/apps/server)

---

## Layout

```
apps/server/internal/
├── platform/testutil/          # NewTestRouter, PerformRequest, AssertJSON
├── users/
│   ├── handler/handler_test.go
│   ├── usecase/service_test.go
│   └── repository/postgres/user_repo_integration_test.go  // integration
└── notification/jobs/welcome_email_test.go
```

---

## Unit tests

### Users usecase

Mocks `UserRepository` and `WelcomeEmailEnqueuer`. Verifies create, duplicate email conflict, not-found paths.

```bash
cd apps/server
pnpm exec node ./scripts/with-go-env.mjs test ./internal/users/usecase/...
```

### Users handler

Uses real `usecase.Service` with stub repository + `middleware.ErrorHandler`. Asserts HTTP status and JSON shape.

```bash
pnpm exec node ./scripts/with-go-env.mjs test ./internal/users/handler/...
```

### Notification jobs

Tests asynq task payload parsing and stub processing.

```bash
pnpm exec node ./scripts/with-go-env.mjs test ./internal/notification/jobs/...
```

---

## Integration tests

Require Docker (testcontainers Postgres 18).

```bash
pnpm test:integration
```

Skips automatically when build tag not set:

```bash
go test ./...   # does not run integration files
```

---

## Notes for Windows

`-race` requires CGO. Use plain `go test` locally; enable race in Linux CI.

---

## Related

- [integration.md](./integration.md)
- [coverage.md](./coverage.md)
