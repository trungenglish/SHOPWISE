# Dynamic UI Protocol Tasks

The current `AgentEnvelope` renderer is a tested, bounded subset of this draft
feature. It supports questions, hydrated recommendations, comparisons, and a
checkout-ready confirmation action. It does not complete FR-001 through FR-009.

Completion is based on the evidence listed in each task, not on a checkbox.

| ID | Status | Task | Required evidence |
| --- | --- | --- | --- |
| DYN-001 | Pending | Define one versioned, framework-neutral JSON Schema shared by Python and TypeScript. | Schema validation tests in both runtimes. |
| DYN-002 | Pending | Cover the layout, shopping, AI, interaction, commerce, and error component families in FR-001 through FR-006. | Contract fixtures for every supported component. |
| DYN-003 | Pending | Add component IDs, actions, payloads, events, and state to every interactive component. | Invalid interaction payload tests and valid round-trip tests. |
| DYN-004 | Pending | Implement append/replace semantics and explicit conversation, UI, component, loading, and error states. | Reducer tests for partial and out-of-order updates. |
| DYN-005 | Pending | Route declarative interaction events through the Go backend to the AI Runtime/tools. | Go ownership tests plus frontend interaction tests. |
| DYN-006 | Pending | Reject embedded HTML/CSS/JavaScript and safely ignore unsupported future components. | Security and backward-compatibility tests. |
| DYN-007 | Pending | Add a second renderer independent of React. | Identical fixture rendered by React and the second renderer. |
| DYN-008 | Pending | Run the end-to-end discovery-to-checkout protocol scenarios. | Recorded automated E2E results with no text-only fallback. |

