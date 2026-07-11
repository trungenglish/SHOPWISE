# SHOPWISE Agent UI Protocol (SAUP)

Version: 1.0

---

## Purpose

SAUP (Shopwise Agent UI Protocol) is the foundational UI tree protocol for the SHOPWISE AI ecosystem. It defines a declarative, streaming-ready protocol between the AI Runtime and UI Renderer.

To support SHOPWISE's goal of becoming the AI-native Operating System for Retail, the interaction is split into two layers:

1. **Decision Protocol (SDP)**: The AI reasoning layer where the agent decides on the intent, products, reasoning, and actions.
2. **UI Protocol (SAUP)**: A UI Composer receives the Decision Protocol output and generates a UI AST (Abstract Syntax Tree) representing the view.

The UI Protocol describes _what_ UI components are displayed and _what_ actions are available. The frontend decides _how_ they are rendered, animated, and adapted to platforms (Web, Mobile, POS, Smart TV).

---

## Architecture Overview

```text
                Customer
                     │
                     ▼
              Conversation API
                     │
                     ▼
              Python Agent Runtime
                     │
      ┌──────────────┼──────────────┐
      ▼              ▼              ▼
  Planner       Tool Calling    Reasoner
      │              │              │
      └──────────────┼──────────────┘
                     ▼
              Decision Engine
                     │
                     ▼
            SDP (Decision Protocol)
                     │
                     ▼
               UI Composer
                     │
                     ▼
            SAUP (UI AST Protocol)
                     │
          JSON + JSON Patch Stream
                     │
                     ▼
            React / Flutter Renderer
                     │
                     ▼
             Interactive Workspace
```

## Folder Structure

1. `01-overview.md` - Philosophy and Architecture
2. `02-decision-protocol.md` - SDP (Decision Protocol) details
3. `03-ui-ast.md` - Core UI Tree Model
4. `04-node-system.md` - Node Types
5. `05-layout-engine.md` - Layouts
6. `06-component-registry.md` - Component mapping
7. `07-actions.md` - Action handling
8. `08-events.md` - Event routing
9. `09-state.md` - UI State
10. `10-streaming.md` - Streaming & Patches
11. `11-renderer.md` - Rendering loop
12. `12-schema.md` - JSON schema
13. `13-security.md` - Security boundaries
14. `14-versioning.md` - Versioning and fallbacks
15. `15-examples.md` - Full examples
16. `16-reference.md` - Type reference
