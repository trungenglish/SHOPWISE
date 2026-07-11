# Commerce Provider

The Commerce Provider is responsible for the core transactional loop (cart and checkout).

```go
type CommerceProvider interface{
    CreateCart(ctx context.Context, userID string) (Cart, error)
    AddItem(ctx context.Context, cartID string, item CartItem) (Cart, error)
    RemoveItem(ctx context.Context, cartID string, itemID string) (Cart, error)
    UpdateQuantity(ctx context.Context, cartID string, itemID string, qty int) (Cart, error)
    Checkout(ctx context.Context, cartID string, payment PaymentInfo) (Order, error)
}
```
