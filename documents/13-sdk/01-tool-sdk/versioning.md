# Tool Versioning

Tools evolve over time. The Tool SDK supports versioning directly in the tool metadata.

```json
{
  "tool": "catalog.search",
  "version": "1.0"
}
```

When breaking changes are introduced to a Retailer's API, the Go Backend can implement version `2.0` of the tool while maintaining `1.0` for legacy AI prompts, ensuring backward compatibility.
