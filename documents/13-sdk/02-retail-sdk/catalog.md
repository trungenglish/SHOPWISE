# Catalog Provider

Handles product discovery and information retrieval.

```go
type CatalogProvider interface {
    Search(ctx context.Context, query string) (SearchResult, error)
    Get(ctx context.Context, productID string) (Product, error)
    Compare(ctx context.Context, productIDs []string) (Comparison, error)
    Categories(ctx context.Context) ([]Category, error)
    Brands(ctx context.Context) ([]Brand, error)
}
```
