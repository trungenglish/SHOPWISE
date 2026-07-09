# AI Architecture

Version: 1.0

---

# Vision

SHOPWISE adopts an Agentic AI architecture rather than a traditional chatbot architecture.

The AI is responsible for:

- understanding customer intent
- planning actions
- invoking tools
- reasoning over retrieved information
- generating interactive user experiences

---

# Architecture

                 User
                    │
                    ▼
          Conversation Manager
                    │
                    ▼
           Shopping Decision Agent
                    │

──────────────────────────────────────────

Planning

↓

Tool Calling

↓

Observation

↓

Reasoning

↓

UI Generation

↓

Response

──────────────────────────────────────────

                    │
                    ▼
          Retail Connector
                    │
                    ▼
      Retail Systems / Knowledge Base

---

# Core Components

Conversation Manager

Maintains conversational context.

---

Shopping Decision Agent

The central orchestrator responsible for reasoning and decision support.

---

Planning Engine

Determines which tools should be executed.

---

Tool Runtime

Executes business operations.

---

Reasoning Engine

Explains recommendations and trade-offs.

---

UI Generator

Produces structured UI schemas for frontend rendering.

---

Memory Manager

Maintains short-term and long-term customer memory.
