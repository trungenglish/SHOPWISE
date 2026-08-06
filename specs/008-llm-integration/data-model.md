# Data Model: AI Runtime LLM Integration

This document defines the core data models and entities used within the AI Runtime for LLM Integration.

## Entities

### `LLMProviderConfig`
Configuration for connecting to an LLM provider.
- `provider_name` (string): e.g., "openai", "anthropic".
- `api_key` (SecretStr): The secret API key.
- `base_url` (string, optional): Custom endpoint.
- `model_name` (string): e.g., "gpt-4o".
- `timeout_ms` (integer): Request timeout in milliseconds.
- `retry_policy` (object): Max retries, backoff factor.

### `AIErrorState`
Structured representation of a failure to communicate with or process data from the LLM provider.
- `error_code` (string): e.g., `PROVIDER_TIMEOUT`, `RATE_LIMITED`, `VALIDATION_FAILED`.
- `message` (string): Human-readable error description.
- `retryable` (boolean): Whether the client should attempt the request again later.
- `provider` (string): The provider that failed (if applicable).

### `DynamicUIResponse`
The expected structured output from the LLM when generating UI updates.
- `schema_version` (string): Protocol version.
- `type` (string): e.g., "product_list", "comparison_table", "checkout_status".
- `data` (object): The payload conforming to the specific type.
- `tool_calls` (array, optional): Any tools the AI decided to invoke.

### `ChatRequest`
The incoming request from the Go Backend/Frontend to the AI Runtime.
- `session_id` (string): Unique identifier for logging and tracing.
- `messages` (array): List of conversation messages (role, content).
- `system_prompt` (string, optional): Overrides the default system prompt.
- `tools` (array): List of available tool schemas for this interaction.
- `temperature` (float, optional): Model temperature.

## Validation Rules
- All incoming requests must be validated using Pydantic.
- `LLMProviderConfig` must validate the presence of an API key on startup.
- Responses from the LLM mapped to `DynamicUIResponse` must pass strict schema validation. Failure triggers self-correction.
