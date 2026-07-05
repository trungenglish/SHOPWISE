# Test coverage

**Repository:** [apps/server](https://github.com/quanngynx/shopwise/tree/main/apps/server)

---

## Targets (MVP)

| Package                     | Target                           |
| --------------------------- | -------------------------------- |
| `users/usecase`             | 90%+                             |
| `users/handler`             | 80%+                             |
| `notification/jobs`         | 80%+                             |
| `users/repository/postgres` | Integration tests for CRUD paths |

---

## Commands

```bash
cd apps/server
go test -cover ./internal/users/...
go test -coverprofile=coverage.out ./internal/users/...
go tool cover -html=coverage.out
```

---

## What to prioritize

1. Error paths (validation, conflict, not found)
2. Enqueue port invoked on create
3. Repository duplicate email + delete semantics

---

## Related

- [server-testing.md](./server-testing.md)
