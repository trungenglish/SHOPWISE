# Tool SDK Overview

## Philosophy

The core philosophy of the Tool SDK is that the **LLM never interacts with systems directly**.

- LLM **does not call REST APIs**.
- LLM **does not query Databases**.

It only executes predefined Tools via the **Tool Protocol (JSON)**.

```text
LLM
 ↓
Planning
 ↓
Tool Call
 ↓
Go Backend
 ↓
Observation
 ↓
Reasoning
```

By abstracting away the underlying business logic, the LLM only operates on semantic intentions (e.g., "search catalog", "add to cart") and observes results.
