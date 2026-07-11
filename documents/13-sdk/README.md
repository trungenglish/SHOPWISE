# SHOPWISE Dual SDK Protocol

To achieve total decoupling between the AI decision-making layer and the Retail Business Logic layer, SHOPWISE defines two strictly separated SDK contracts:

- **Tool SDK**: Used by the AI Runtime to interact with the backend.
- **Retail SDK**: Used by the backend to interact with Retailers (Shopee, PhongVu, Tiki, etc.).

## Architecture

```text
                    Python AI Runtime
                            │
                  Planning / Reasoning
                            │
                    Tool SDK Contract
                            │
────────────────────────────────────────────────────────
                    Go Business Backend
                            │
                     Tool Registry
                            │
               Catalog / Commerce / Memory
                            │
                    Retail SDK Contract
                            │
────────────────────────────────────────────────────────
         PhongVu     Shopee     Lazada     CellphoneS
             │           │           │            │
             └───────────┼───────────┴────────────┘
                         ▼
                    Retail APIs
```

This ensures that the AI never needs to know about Go or Retail APIs, and Go never needs to know about LangGraph or PydanticAI.
