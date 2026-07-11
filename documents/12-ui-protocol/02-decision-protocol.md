# Decision Protocol (SDP)

The **Shopwise Decision Protocol (SDP)** is the interface between the AI reasoning engine and the UI Composer.

Instead of deciding what buttons or cards to render, the AI outputs its raw, structured decision.

## Format

```json
{
  "intent": "recommend",
  "products": [
    {
      "id": "prod_123",
      "match_score": 0.95
    }
  ],
  "reasoning": "This product matches the user's requirement for a lightweight running shoe.",
  "confidence": 0.93,
  "next_actions": ["compare", "checkout"]
}
```

## Workflow

1. The **Python Agent Runtime** analyzes the user's message using Planners, Tools, and Reasoners.
2. The **Decision Engine** formulates an SDP response.
3. The **UI Composer** intercepts the SDP response and dynamically builds a `UI AST` (SAUP) tailored to the user's client application.
