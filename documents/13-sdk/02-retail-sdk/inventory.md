# Inventory Provider

Handles stock levels, warehouse locations, and delivery estimates.

```go
type InventoryProvider interface{
    Availability(ctx context.Context, productID string) (StockLevel, error)
    Warehouse(ctx context.Context, locationID string) (WarehouseInfo, error)
    DeliveryEstimate(ctx context.Context, productID string, zip string) (DeliveryEstimate, error)
}
```
