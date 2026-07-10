# Capability Negotiation

Not every retailer supports the same features. Some support guest checkout, others don't. Some have real-time inventory, others have batch syncs.

## Provider Manifest

Each Provider declares its capabilities:

```yaml
provider: phongvu
version: 1.0

capabilities:
  catalog: true
  inventory: true
  promotions: true
  cart: true
  checkout: true
  guest_checkout: false
  realtime_inventory: true
```

## AI Negotiation

Before suggesting an action, the backend negotiates capabilities.

If the user wants to check out, but the Provider (e.g., Shopee) returns `SupportsCheckout: false` (because they only allow checkout via their app), the backend intercepts the `checkout.prepare` tool call and instructs the AI to generate a deep link to the Shopee app instead of showing an in-app checkout panel.
