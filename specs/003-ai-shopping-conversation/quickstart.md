# Quickstart: AI Shopping Conversation MVP

This guide explains how to validate the feature end-to-end once implemented.

## Prerequisites
- Node.js & pnpm (for React frontend)
- Go 1.24 (for Retail Backend)
- Python 3.x with `uv` (for AI Runtime)

## Running the Stack
1. **Frontend**: `cd apps/web && pnpm dev` (Runs on port 5173)
2. **Go Backend**: `cd services/main-backend && go run main.go` (Runs on port 8080)
3. **AI Runtime**: `cd services/ai-runtime && uvicorn main:app --reload` (Runs on port 8000)

## Validation Scenarios

### Scenario 1: Greeting & Discovery
1. Open the web UI at `http://localhost:5173`.
2. Verify the AI greets you and provides quick-reply pills (e.g., "Laptop", "Gaming").
3. Type "I need a gaming laptop under 20 million VND".

### Scenario 2: Recommendations
1. Observe the AI querying the Go backend catalog (via logs).
2. Verify the UI renders a **Product Carousel** schema.
3. Verify each product includes an explainable AI insight (trade-offs).

### Scenario 3: Comparison
1. Ask "Compare the first two laptops".
2. Verify the UI renders a **Comparison Table** schema.

### Scenario 4: Checkout & Recovery
1. Ask "I want to buy the Asus TUF".
2. Verify the AI outputs a **Checkout Summary** schema.
3. Simulate an Out-of-Stock scenario (by modifying the mock JSON). Verify the AI informs you and suggests alternatives instead of auto-replacing the item.
