# Reference Interfaces

This is arguably the most critical part.

## Tool

```go
type Tool interface {
    Metadata() ToolMetadata
    InputSchema() JSONSchema
    OutputSchema() JSONSchema
    Execute(ctx context.Context, input any) (any, error)
}
```

---

Retail

```go
type RetailProvider interface{
    Catalog()
    Inventory()
    Promotion()
    Commerce()
    Orders()
}
```

---

Memory

```go
type MemoryStore interface{
    Read()
    Write()
    Delete()
    Search()
}
```

---

Knowledge

```go
type KnowledgeStore interface{
    Retrieve()
    Embed()
    Update()
}
```

---

UI

```go
type UIRenderer interface{
    Render(schema UISchema)
}
```
