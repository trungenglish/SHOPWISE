# Order Provider

Handles post-purchase workflows.

```go
type OrderProvider interface{
    Status(ctx context.Context, orderID string) (OrderStatus, error)
    Tracking(ctx context.Context, orderID string) (TrackingInfo, error)
    Warranty(ctx context.Context, orderID string) (WarrantyInfo, error)
    Cancel(ctx context.Context, orderID string, reason string) (Order, error)
}
```
