# Data Model: AI Shopping Conversation

## Entities

### 1. CustomerSession
Represents an active shopping conversation. Managed by the Go Backend.

- `sessionId` (String, UUID)
- `startedAt` (Timestamp)
- `lastActiveAt` (Timestamp)
- `context` (Object): Short-term memory of user preferences (e.g., extracted budget, brand preference, purpose).
- `messages` (Array): Conversation history (User and AI turns).

### 2. Product
Represents an item in the catalog. Prices MUST be in VND.

- `productId` (String)
- `name` (String)
- `brand` (String)
- `category` (String)
- `price` (Number, VND)
- `specs` (Object): Key-value pairs of technical specifications.
- `inStock` (Boolean)
- `images` (Array of URLs)

### 3. Recommendation
The AI's justification for a product.

- `productId` (String)
- `reasoning` (String): Why this fits the user's needs.
- `tradeoffs` (String): What the user sacrifices (e.g., "Heavier than alternatives").
- `confidence` (Number, 0.0-1.0)

### 4. CheckoutState
Prepared checkout data.

- `checkoutId` (String)
- `items` (Array of Product IDs)
- `totalPrice` (Number, VND)
- `status` (Enum: PENDING, CONFIRMED, REJECTED)
