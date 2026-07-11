# Phase 0: Research & Technical Validation

**Feature**: AI Runtime LLM Integration (008-llm-integration)

## 1. Provider Abstraction & Model Invocation
- **Decision**: Use `langchain-openai` for initial provider integration, wrapped in a generic `ProviderBase` interface.
- **Rationale**: LangChain provides built-in mechanisms for standardizing tool binding, schema forcing (e.g. `with_structured_output`), and streaming across different LLMs. Wrapping it in our own interface ensures we can later add Azure, Anthropic, or Gemini without leaking LangChain-specific types to our core application logic.
- **Alternatives considered**: Raw OpenAI Python SDK. Rejected because implementing multi-provider support, structured outputs across different APIs, and standardizing tool parsing is heavy boilerplate that LangChain already solves.

## 2. Streaming (Server-Sent Events)
- **Decision**: Use FastAPI with `sse-starlette` `EventSourceResponse` returning a stream of token chunks and JSON patches.
- **Rationale**: It perfectly matches the requirement for a transport-agnostic yet HTTP-based real-time protocol. `sse-starlette` integrates natively with FastAPI's async generators. 
- **Alternatives considered**: WebSockets (rejected as initial scope per specification clarification; harder to scale, requires stateful connection tracking).

## 3. Structured Outputs & Self-Correction
- **Decision**: Use LangGraph for the execution loop.
- **Rationale**: LangGraph is purpose-built for cyclical LLM workflows like self-correction loops. We can define a node for "invoke_model", a node for "validate_json", and a conditional edge that routes back to "invoke_model" (up to 2 retries) if validation fails.
- **Alternatives considered**: Standard Python `while` loop. Rejected because LangGraph provides built-in state management, persistence (if ever needed), and observability for complex AI workflows.

## 4. Bounded Retries for Transient Errors
- **Decision**: Use `tenacity` for exponential backoff and jitter on HTTP/Provider errors.
- **Rationale**: `tenacity` is the Python standard for retry logic. It integrates well with async functions and can be configured to catch specific provider exceptions (e.g., `openai.RateLimitError`, `openai.APIConnectionError`) while failing fast on `openai.AuthenticationError`.

## 5. Security & PII
- **Decision**: Use Pydantic `SecretStr` for API keys in configuration.
- **Rationale**: `SecretStr` prevents accidental printing or logging of the key. Environment variables are loaded via `pydantic-settings`.
