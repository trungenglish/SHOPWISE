# Event Protocol

Events flow from the UI back to the AI Runtime, creating an interactive loop.

```text
User Click
    ↓
Trigger UIAction
    ↓
Tool Execution (Frontend or Backend)
    ↓
Observation (Result)
    ↓
UI Composer generates Patch
    ↓
Frontend Re-renders
```

The Frontend doesn't implement business logic; it merely forwards actions to the Conversation API and applies SAUP state patches in response.
