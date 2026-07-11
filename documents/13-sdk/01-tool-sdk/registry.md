# Tool Registry

Tools are not hardcoded into the AI's prompt. They are discovered dynamically from the Go Backend via the Registry.

## Backend Registration

```go
registry.Register(
    SearchProductsTool{}
)
```

## Discovery

When the AI starts a conversation, it pings the Backend:

```text
AI Startup
   ↓
Backend Tool Registry
   ↓
[
 "catalog.search",
 "cart.add",
 "checkout.prepare"
]
```

This allows the backend to enable or disable tools dynamically based on feature flags, retailer capabilities, or user roles without touching the Python AI code.
