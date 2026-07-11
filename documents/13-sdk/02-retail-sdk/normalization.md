# Schema Normalization

Retail APIs are messy and inconsistent. The Retail SDK enforces a **Canonical Schema**.

```text
Retail API
    ↓
Provider Adapter
    ↓
Canonical Schema
    ↓
Tool Observation
    ↓
LLM
```

## Example

**Shopee API:**

```json
{
  "item_name": "Laptop",
  "item_price": "15000000"
}
```

**Phong Vũ API:**

```json
{
  "name": "Laptop",
  "salePrice": 15000000
}
```

**Canonical Model (Go):**

```go
type Product struct {
    Name  string
    Price float64
}
```

The LLM **only** sees the Canonical Model, meaning the AI prompt never needs to have specialized logic for parsing Shopee vs Phong Vũ responses.
