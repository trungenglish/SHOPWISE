# UI Contracts

As per the Declarative UI principle, the AI Runtime outputs structured schemas instead of raw HTML/React components.

## 1. Product Carousel
Used to display multiple product recommendations.

```json
{
  "type": "carousel",
  "items": [
    {
      "productId": "123",
      "name": "Gaming Laptop X",
      "price_vnd": 25000000,
      "image_url": "...",
      "ai_insight": "Best performance for your budget, but slightly heavy."
    }
  ]
}
```

## 2. Comparison Table
Used to highlight differences between 2 or more products.

```json
{
  "type": "comparison_table",
  "products": ["123", "456"],
  "features": [
    {
      "feature_name": "RAM",
      "values": {"123": "16GB", "456": "32GB"}
    },
    {
      "feature_name": "Price",
      "values": {"123": "25,000,000 VND", "456": "30,000,000 VND"}
    }
  ]
}
```

## 3. Checkout Summary
Rendered when the AI prepares a checkout.

```json
{
  "type": "checkout_summary",
  "checkoutId": "chk_999",
  "items": [...],
  "total_vnd": 25000000,
  "requires_confirmation": true
}
```

## 4. Error State
Used for graceful recovery when background tools fail.

```json
{
  "type": "error_state",
  "message": "We couldn't reach the live catalog.",
  "actions": [
    {"label": "Retry", "action": "retry_catalog_search"},
    {"label": "Continue with Cache", "action": "use_cached_data"}
  ]
}
```
