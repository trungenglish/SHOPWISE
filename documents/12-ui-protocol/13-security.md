# Security Boundaries

SAUP ensures strict security by preventing the LLM from generating arbitrary executable code.

## Explicit Restrictions

The AI is explicitly forbidden from generating:
*   HTML
*   CSS
*   JavaScript (including React components, hooks)
*   Flutter Widgets
*   Tailwind classes

## Why?
1. **XSS Protection**: By ensuring the AI only outputs structural Nodes and Props, we prevent malicious injection.
2. **Determinism**: We can guarantee exactly how a `checkout.panel` looks and behaves.
3. **App Store Compliance**: Native apps cannot execute arbitrary downloaded code, but they can render native UI dynamically from JSON (Server-Driven UI).
