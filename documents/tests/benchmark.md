# Benchmarks

**Repository:** [apps/server](https://github.com/quanngynx/shopwise/tree/main/apps/server)

---

## When to run

- Before optimizing list endpoints or repository queries
- After schema/index changes on `users` table
- Not required for every PR (optional)

---

## Example

Add to `users/repository/postgres` when needed:

```go
func BenchmarkRepositoryList(b *testing.B) {
  // setup db + seed rows
  for b.Loop() {
    _, _ = repo.List(context.Background(), 20, 0)
  }
}
```

Run:

```bash
go test -bench=. -benchmem ./internal/users/repository/postgres/...
```

---

## Baseline expectations (MVP)

No formal SLA. Target P95 API latency under 500ms for list/create on local docker-compose with warm DB pool.

---

## Related

- [coverage.md](./coverage.md)
- [server-testing.md](./server-testing.md)
