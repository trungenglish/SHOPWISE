# Error Contract

Errors are returned to the LLM in a standardized format so it can reason about what went wrong.

```json
{
  "code": "PRODUCT_NOT_FOUND",
  "message": "The product ID you provided does not exist.",
  "retryable": false
}
```

The LLM uses `retryable` to decide whether to attempt the tool call again with different parameters or immediately inform the user.
