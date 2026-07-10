# Permission Model

The Tool SDK enforces a role-based access control (RBAC) model. LLMs cannot bypass this because execution happens in the authenticated Backend.

## Hierarchy

```text
Public
  ↓
Authenticated
  ↓
Customer
  ↓
Admin
  ↓
System
```

## Examples

*   `catalog.search` -> **Public** (Anyone can search)
*   `checkout.submit` -> **Customer** (Must be logged in and checking out their own cart)
*   `promotion.sync` -> **Admin** (Internal staff updating campaigns)

If the AI attempts to call a tool the current user does not have permission for, the backend returns a clean observation indicating the failure so the AI can gracefully inform the user.
