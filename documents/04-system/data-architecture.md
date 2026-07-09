# Data Architecture

Version: 1.0

---

# Primary Database

PostgreSQL

Stores

Products

Users

Orders

Preferences

Catalog

Sessions

---

# Vector Search

pgvector

Stores

Product Embeddings

Buying Guides

FAQ

Warranty

---

# Cache

Redis

Stores

Conversation

Temporary Memory

Streaming

---

# Source of Truth

Retail APIs

↓

Retail Connector

↓

Normalized Catalog

↓

AI Retrieval
