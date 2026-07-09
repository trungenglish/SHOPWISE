# Domain Model

Version: 1.0

---

# Overview

SHOPWISE is modeled around the customer's decision-making journey rather than traditional e-commerce operations.

The domain is centered on helping customers make informed purchasing decisions through AI-assisted interactions.

---

# Core Domains

## Product Discovery

Responsible for helping customers discover relevant products based on natural language requirements.

Capabilities

- Search
- Filtering
- Semantic Retrieval

---

## Recommendation

Responsible for identifying products that best satisfy customer goals.

Capabilities

- Candidate Generation
- Ranking
- Personalization

---

## Product Comparison

Responsible for presenting objective comparisons between multiple products.

Capabilities

- Specification Comparison
- Performance Comparison
- Price Comparison
- Trade-off Analysis

---

## Decision Support

Responsible for reducing uncertainty before purchase.

Capabilities

- Explainability
- Confidence Estimation
- Alternative Suggestions

---

## Commerce

Responsible for executing retailer operations.

Capabilities

- Cart
- Coupon
- Checkout
- Inventory

---

## Knowledge

Responsible for answering questions about products and retailer policies.

Capabilities

- FAQ
- Warranty
- Buying Guide
- Promotion

---

## Conversation

Responsible for maintaining conversational context.

Capabilities

- Session
- Context
- Memory
- History

---

# Domain Relationships

Conversation

↓

Requirement Understanding

↓

Product Discovery

↓

Recommendation

↓

Comparison

↓

Decision Support

↓

Commerce

↓

Post-purchase
