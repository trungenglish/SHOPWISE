# Tool Lifecycle

The lifecycle of a tool execution in the SHOPWISE backend goes through strict phases:

```text
Discover
   ↓
Validate Input
   ↓
Authorize
   ↓
Execute
   ↓
Normalize
   ↓
Observe
   ↓
Return
```

1.  **Discover**: LLM discovers what tools are available dynamically.
2.  **Validate Input**: Backend validates the JSON schema of the LLM's tool call.
3.  **Authorize**: Checks if the current user session has permission to run the tool (e.g., `checkout.submit` requires `Customer` auth).
4.  **Execute**: Runs the business logic (calling Retail SDK or internal services).
5.  **Normalize**: Standardizes the raw response.
6.  **Observe**: Packages the response into an "Observation" (a summarized format easier for the LLM to reason about).
7.  **Return**: Returns the JSON observation to the AI Runtime.
