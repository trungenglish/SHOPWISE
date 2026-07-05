# Integration testing

**Repository:** [apps/server](https://github.com/quanngynx/shopwise/tree/main/apps/server)

---

## Overview

Integration tests use [testcontainers-go](https://golang.testcontainers.org/) to spin up ephemeral Postgres containers. They validate GORM repository SQL against a real database.

---

## Prerequisites

- Docker Desktop or Docker Engine running
- Go 1.25+
- Dependencies: `github.com/testcontainers/testcontainers-go/modules/postgres`

---

## Running tests

From `apps/server`:

```bash
pnpm test:integration
```

Equivalent:

```bash
go test -tags=integration ./internal/users/repository/postgres/...
```

---

## Test file convention

```go
//go:build integration

package postgres_test
```

Files without this tag are included in default `go test ./...`.

---

## CI recommendation

Run integration tests on Linux agents with Docker:

```yaml
- run: go test -tags=integration ./internal/users/repository/postgres/...
```

---

## Troubleshooting

| Issue              | Fix                                     |
| ------------------ | --------------------------------------- |
| Docker not running | Start Docker Desktop                    |
| Port conflicts     | testcontainers uses random ports        |
| Slow first run     | Image pull for `postgres:18.4-bookworm` |

---

## Related

- [server-testing.md](./server-testing.md)
- [cicd-pipeline.md](./cicd-pipeline.md)
