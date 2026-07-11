# MCP Bridge

Because the Tool SDK abstracts business logic into language-agnostic JSON protocols, it is fully compatible with the **Model Context Protocol (MCP)**.

## Architecture

```text
AI Runtime
      │
Tool Registry
      │
 MCP Bridge
      │
Claude / Codex / Cursor / OpenAI
```

## Why MCP?

By exposing the Tool Registry via an MCP Bridge:

1.  Our internal tools can be used by third-party AI clients (e.g., Cursor, Claude Desktop).
2.  SHOPWISE becomes an **AI Commerce Platform** rather than just a single application.
3.  Developers can test retail workflows directly from their IDEs without running the full Python AI Runtime.
