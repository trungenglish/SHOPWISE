# Promotion Provider

Handles active sales, discounts, and bundles.

```go
type PromotionProvider interface{
    Campaigns(ctx context.Context) ([]Campaign, error)
    Coupons(ctx context.Context, userID string) ([]Coupon, error)
    Bundles(ctx context.Context, productID string) ([]Bundle, error)
}
```
