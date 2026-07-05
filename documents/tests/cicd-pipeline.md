# CI/CD pipeline (server)

**Repository:** [apps/server](https://github.com/quanngynx/shopwise/tree/main/apps/server)

---

## Recommended pipeline steps

```yaml
# Example GitHub Actions job (apps/server)
steps:
  - uses: actions/checkout@v4
  - uses: actions/setup-go@v5
    with:
      go-version: "1.25.x"

  - name: Vet
    working-directory: apps/server
    run: go vet ./...

  - name: Unit tests
    working-directory: apps/server
    run: go test -race ./...

  - name: Integration tests
    working-directory: apps/server
    run: go test -tags=integration ./internal/users/repository/postgres/...

  - name: Build binaries
    working-directory: apps/server
    run: |
      go build -o bin/server ./cmd/server
      go build -o bin/worker ./cmd/worker

  - name: OpenAPI drift check
    working-directory: apps/server
    run: |
      go run github.com/swaggo/swag/cmd/swag@v1.16.6 init -g cmd/server/main.go --parseInternal -d ./
      git diff --exit-code docs/
```

---

## Monorepo integration

From repo root, Turbo can invoke `apps/server` scripts:

```bash
pnpm --filter server check-types
pnpm --filter server test:unit
pnpm --filter server build
```

---

## Docker

```bash
cd apps/server
docker compose build
docker compose up -d
curl http://localhost:18080/health
```

---

## Related

- [integration.md](./integration.md)
- [security.md](./security.md)
