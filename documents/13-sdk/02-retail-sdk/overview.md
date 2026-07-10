# Retail SDK Overview

## Philosophy

The Retail SDK ensures the SHOPWISE Go Backend does not hardcode integrations like `PhongVuAPI` or `ShopeeAPI`.

Instead, it relies on generic `Retail Provider` adapters.

```text
Retail API
    ↓
Provider Adapter (PhongVu/Shopee)
    ↓
Retail SDK (Generic Interfaces)
    ↓
Commerce Service
    ↓
Tool SDK
    ↓
AI Runtime
```

This Adapter Pattern allows scaling to dozens of external retail partners by simply writing a new interface implementation.
