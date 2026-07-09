# Information Architecture

Version: 1.0

---

# Philosophy

SHOPWISE organizes data around information consumed by AI and users, rather than database tables.

The architecture separates:

- Operational Data
- Knowledge
- Memory
- Derived Intelligence

Each category has different ownership, lifecycle, and storage requirements.

---

# Information Layers

Customer

↓

Conversation

↓

Requirements

↓

Knowledge

↓

Reasoning

↓

Commerce

↓

Orders

---

# Information Domains

Customer

Product

Catalog

Inventory

Promotion

Warranty

Conversation

Recommendation

Comparison

Decision

Commerce

Analytics

---

# Source of Truth

Customer → Backend

Products → Retail Provider

Inventory → Retail Provider

Promotions → Retail Provider

Orders → Retail Provider

Knowledge → SHOPWISE

Memory → SHOPWISE

Analytics → SHOPWISE

---

# Data Ownership

Retailers own commerce data.

SHOPWISE owns decision intelligence.

Customers own preferences and memories.