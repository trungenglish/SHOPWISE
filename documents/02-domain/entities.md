# Entities

Version: 1.0

---

# Product

Represents a sellable consumer electronics item.

Attributes

- Product ID
- Name
- Brand
- Category
- Specifications
- Price
- Images

Relationships

Inventory

Promotion

Warranty

Retail Provider

---

# Customer

Represents a shopper interacting with SHOPWISE.

Attributes

- Customer ID
- Preferences
- Budget
- Purchase History

---

# Shopping Session

Represents one decision-making session.

Attributes

- Session ID
- Conversation History
- Requirements
- Recommendations
- Compared Products

---

# Requirement

Represents structured customer needs extracted from conversation.

Examples

Budget

Brand

CPU

GPU

Weight

Battery

Usage

Priority

---

# Recommendation

Represents an AI-generated recommendation.

Attributes

Product

Reasoning

Confidence

Alternatives

---

# Comparison

Represents comparison between products.

Contains

Compared Products

Metrics

Trade-offs

Summary

---

# Shopping Cart

Represents selected products.

Attributes

Items

Coupons

Total

Retail Provider

---

# Order

Represents a completed purchase.

Attributes

Order ID

Products

Status

Payment

Shipping

---

# Retail Provider

Represents an external retailer.

Examples

Phong Vu

Shopee

CellphoneS

Future implementations provide provider-specific integrations.
