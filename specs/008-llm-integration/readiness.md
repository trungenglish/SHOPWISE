# SHOPWISE Agent Readiness

Snapshot: 2026-08-04. This document records observed code and runnable checks;
historical task checkboxes are not accepted as completion evidence.

## Current status

| Feature | Status | Current evidence | Remaining gate |
| --- | --- | --- | --- |
| 002 onboarding | Partial | The current identity/user OpenAPI contract is loaded by Go contract tests. | Run browser E2E against a live stack and retain the result. |
| 003 shopping conversation | MVP verified; broader feature partial | Session ownership, ordered multi-turn history, assistant envelope persistence, catalog-only recommendations, comparison restrictions, and checkout-ready confirmation have automated tests. Live Go -> AI Runtime -> OpenAI -> catalog scenarios passed for clarification, recommendation, comparison streaming, upstream failure, and two-session isolation. | Add a browser-level E2E for the dashboard before treating the full feature as complete. |
| 004 dynamic UI | Partial | React renders the four current `AgentEnvelope` variants; checkout requires an explicit click in its component test. | Complete `../004-dynamic-ui-protocol/tasks.md`; the general declarative protocol remains draft. |
| 008 LLM integration | MVP verified; advanced gates partial | Python tests cover configuration, exact model forwarding, OpenAI-compatible strict schema, structured output, two self-corrections, catalog failures, unknown product rejection, and SSE events. Live requests using the existing key succeeded with requested model `gpt-5.4-mini`. | Measure TTFT only after true token streaming exists; add real OpenTelemetry only when an exporter/collector is selected. |

## Verification evidence

| Gate | Result |
| --- | --- |
| Python `uv run pytest -q` | Passed: 15 tests. |
| Python Ruff and strict mypy | Passed. |
| Go `go test ./...` | Passed. |
| Go race test for decision memory | Blocked before execution: `CGO_ENABLED=0` and no C compiler. |
| Web `vitest run` | Passed: 14 tests across 5 files. |
| Root `pnpm run check-types` | Passed; includes the web production build and TypeScript check. |
| Root `pnpm run build` | Passed without Windows `EPERM`; bundle-size warning only. |
| Root `pnpm run check` | Not passed: the full legacy tree timed out; targeted checks for the new renderer and mapping files passed. |
| Live OpenAI smoke | Passed: requested `gpt-5.4-mini`; response metadata was `gpt-5.4-mini-2026-03-17`; response contained a valid choice. No key or response content was printed. |
| Live multi-service API E2E | Passed: clarification; 3 catalog-only VND recommendations; SSE `token,comparison,done`; cross-session read 404; AI Runtime down returned 502 with no fake assistant message. |
| Live browser E2E | Not run. Web component/API tests cover envelope, error, retry, and explicit checkout confirmation, but do not replace this gate. |

## Remaining work in gate order

1. Make the root Ultracite check bounded and repeatable, then fix its real
   diagnostics without sweeping unrelated example code into this feature.
2. Add a browser-level dashboard E2E over the already verified live API paths.
3. Enable CGO with a supported C compiler and run the decision-memory race test.
4. Complete the Dynamic UI tasks in `../004-dynamic-ui-protocol/tasks.md` before
   changing feature 004 from Partial to Complete.
5. Replace the validated single-message SSE token with true LLM token streaming
   only when measured latency justifies the added parsing and validation path.
