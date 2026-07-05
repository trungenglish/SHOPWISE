# Security testing checklist

**Repository:** [apps/server](https://github.com/quanngynx/shopwise/tree/main/apps/server)

---

## API security checks

- [ ] Errors return `application/problem+json` (RFC 7807), not stack traces in release mode
- [ ] Swagger UI disabled when `GIN_MODE=release`
- [ ] CORS restricted to `ALLOWED_ORIGINS`
- [ ] Request IDs propagated via `X-Request-ID`
- [ ] Input validation on user create/update (email, name required)
- [ ] No sensitive fields in OpenAPI (`swaggerignore` on internal fields)
- [ ] Container runs as non-root (distroless image)
- [ ] Server container read-only root filesystem in compose

---

## Future (identity module)

- [ ] JWT validation middleware on protected routes
- [ ] Rate limiting at edge or middleware
- [ ] RBAC per administration module

---

## Manual smoke tests

```bash
# CORS preflight
curl -X OPTIONS http://localhost:18080/api/v1/users -H "Origin: http://localhost:3001" -v

# Validation error shape
curl -X POST http://localhost:18080/api/v1/users -H "Content-Type: application/json" -d '{}'
```

---

## Related

- [server-testing.md](./server-testing.md)
- [cicd-pipeline.md](./cicd-pipeline.md)
