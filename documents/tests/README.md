# Comprehensive Testing Guide

**Version:** 1.1  
**Last Updated:** 2026-06-14  
**Status:** Production Ready  
**Repository:** [apps/server](https://github.com/quanngynx/shopwise/tree/main/apps/server)

---

## Test pyramid

| Layer | Scope | Tool |
| --- | --- | --- |
| Unit | usecase, handler, job handlers | `go test`, httptest, mocks |
| Integration | repository + Postgres | testcontainers, build tag `integration` |
| E2E | Full stack (optional) | docker compose + HTTP client |

---

## Quick commands (`apps/server`)

```bash
pnpm test:unit          # all packages except integration-tagged tests
pnpm test:integration   # Postgres repository tests (requires Docker)
pnpm check-types        # go vet ./...
```

---

## Documentation index

| Guide | Description |
| --- | --- |
| [server-testing.md](./server-testing.md) | Server-specific test layout and examples |
| [integration.md](./integration.md) | testcontainers setup |
| [coverage.md](./coverage.md) | Coverage targets and commands |
| [security.md](./security.md) | API security test checklist |
| [cicd-pipeline.md](./cicd-pipeline.md) | CI steps for Go server |
| [benchmark.md](./benchmark.md) | Optional benchmarks |

---

## Build tags

| Tag | Files | When to run |
| --- | --- | --- |
| `integration` | `*_integration_test.go` | CI nightly or pre-release; needs Docker |

Default `go test ./...` skips integration-tagged files.
