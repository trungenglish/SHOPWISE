# Provider Interface

The root of any retail integration is the `Provider` interface.

```go
type Provider interface {
    Catalog() CatalogProvider
    Inventory() InventoryProvider
    Promotion() PromotionProvider
    Commerce() CommerceProvider
    Orders() OrderProvider
}
```

A retailer like Shopee implements this interface by returning their specific sub-providers.
