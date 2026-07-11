# Database Schema Structure

Version: 1.0

This document defines the physical schema structures for the abstract concepts defined in the Information Architecture, mapped to their respective storage engines.

---

## 1. PostgreSQL (Relational Data)

Stores operational data, commerce, and persistent user memory.

### `users`

Stores customer identity and authentication state.

| Column | Type | Constraints | Description |
| :-- | :-- | :-- | :-- |
| `id` | UUID | PRIMARY KEY | Unique identifier for the user |
| `identity` | VARCHAR(255) | UNIQUE, NOT NULL | Primary identity (e.g., email or phone) |
| `created_at` | TIMESTAMP | DEFAULT NOW() | Record creation timestamp |
| `updated_at` | TIMESTAMP | DEFAULT NOW() | Last update timestamp |

### `user_preferences`

Stores preference memory.

| Column | Type | Constraints | Description |
| :-- | :-- | :-- | :-- |
| `id` | UUID | PRIMARY KEY | Unique identifier |
| `user_id` | UUID | FOREIGN KEY (`users.id`) | Reference to user |
| `preferences` | JSONB | NOT NULL | Key-value pairs of preferences (e.g., `{"brand": "Apple", "budget": 1000}`) |
| `updated_at` | TIMESTAMP | DEFAULT NOW() | Last update timestamp |

### `products`

Operational data for products sourced from retail APIs.

| Column | Type | Constraints | Description |
| :-- | :-- | :-- | :-- |
| `id` | UUID | PRIMARY KEY | Internal product identifier |
| `sku` | VARCHAR(100) | UNIQUE, NOT NULL | Retailer SKU |
| `name` | VARCHAR(255) | NOT NULL | Product name |
| `brand` | VARCHAR(100) | INDEX | Product brand |
| `category` | VARCHAR(100) | INDEX | Product category |
| `price` | DECIMAL(10,2) | INDEX | Current price |
| `specifications` | JSONB |  | Detailed specs |
| `metadata` | JSONB |  | Extra metadata (tags, flags) |

### `inventory`

Stock and availability details.

| Column         | Type         | Constraints     | Description                |
| :------------- | :----------- | :-------------- | :------------------------- |
| `product_id`   | UUID         | PRIMARY KEY, FK | Reference to `products.id` |
| `stock`        | INTEGER      | DEFAULT 0       | Quantity available         |
| `warehouse`    | VARCHAR(100) |                 | Warehouse location         |
| `availability` | BOOLEAN      | INDEX           | Is currently available     |

### `promotions`

Discount and campaign structures.

| Column | Type | Constraints | Description |
| :-- | :-- | :-- | :-- |
| `id` | UUID | PRIMARY KEY | Promotion identifier |
| `campaign` | VARCHAR(255) | NOT NULL | Campaign name |
| `coupon_code` | VARCHAR(50) | UNIQUE | Associated coupon code |
| `discount` | JSONB | NOT NULL | Discount details (percentage, flat) |

### `orders`

Commerce transaction data.

| Column | Type | Constraints | Description |
| :-- | :-- | :-- | :-- |
| `id` | UUID | PRIMARY KEY | Order identifier |
| `user_id` | UUID | FOREIGN KEY (`users.id`) | Reference to purchasing user |
| `status` | VARCHAR(50) | NOT NULL | e.g., 'Pending', 'Paid', 'Delivered' |
| `payment` | JSONB |  | Payment details |
| `delivery` | JSONB |  | Shipping and delivery estimates |

---

## 2. pgvector (AI & Embeddings)

Stores knowledge and embeddings for semantic search and retrieval.

### `knowledge_base`

Stores documents like Buying Guides, FAQs, and Policies.

| Column | Type | Constraints | Description |
| :-- | :-- | :-- | :-- |
| `id` | UUID | PRIMARY KEY | Document identifier |
| `type` | VARCHAR(50) | NOT NULL | e.g., 'faq', 'buying_guide', 'warranty' |
| `content` | TEXT | NOT NULL | Raw text content of the knowledge |
| `metadata` | JSONB |  | Filter metadata (brand, category, version) |
| `embedding` | VECTOR(1536) | INDEX (HNSW) | Embedding vector for semantic search |

### `product_insights`

AI-derived insights linked to specific products.

| Column | Type | Constraints | Description |
| :-- | :-- | :-- | :-- |
| `id` | UUID | PRIMARY KEY | Insight identifier |
| `product_id` | UUID | FOREIGN KEY (`products.id`) | Associated product |
| `insight` | TEXT | NOT NULL | Derived insight text |
| `embedding` | VECTOR(1536) | INDEX | Insight embedding |

---

## 3. Redis (Temporary & Streaming State)

Handles caching, fast-access session data, and rate limiting. Represented here via Key Patterns.

### Session Memory

- **Pattern:** `session:{user_id}:{session_id}`
- **Type:** JSON / Hash
- **TTL:** 24 hours
- **Structure:**
  ```json
  {
    "context": "active_conversation",
    "temporary_state": {}
  }
  ```

### Active Cart

- **Pattern:** `cart:{user_id}`
- **Type:** JSON
- **TTL:** 7 days
- **Structure:**
  ```json
  {
    "items": [{ "product_id": "...", "quantity": 1 }],
    "coupons": ["SAVE20"],
    "price_summary": { "subtotal": 100, "total": 80 }
  }
  ```

### Streaming & Rate Limiting

- **Pattern:** `ratelimit:{user_id}:chat`
- **Type:** Counter / Sliding Window
- **TTL:** 60 seconds

---

## 4. Blob Storage (Files)

Handles unstructured assets like media and PDFs.

### Buckets

- `shopwise-images`: Contains product images.
- `shopwise-manuals`: Contains raw PDF manuals and datasheets.
- `shopwise-documents`: Internal architecture and raw retail documents.

### Path Conventions

- **Images:** `s3://shopwise-images/products/{sku}/main.jpg`
- **Manuals:** `s3://shopwise-manuals/brands/{brand}/{sku}-manual.pdf`
