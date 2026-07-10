# Architecture & Philosophy

## Separation of Concerns

The SHOPWISE ecosystem distinctly separates **Decision Intelligence** from **Presentation Intelligence**:

*   **AI (Decision Engine)** owns the **WHAT**: What products to recommend, what reasoning to show, and what the intent of the user is.
*   **UI Composer** owns the **STRUCTURE**: Translating decisions into an abstract UI tree (AST).
*   **Frontend (React/Flutter)** owns the **HOW**: Rendering the UI tree natively, managing animations, responsiveness, and accessibility.

## Why Two Protocols?

By splitting the communication into **SDP (Decision Protocol)** and **SAUP (UI Protocol)**:
1.  **Focused AI**: The LLM focuses purely on data and decisions, rather than trying to construct complex UI layouts.
2.  **Platform Agnostic**: The UI Composer can generate different ASTs for Web vs Mobile vs POS without the AI needing to know the target platform.
3.  **Independent Scaling**: The Decision Engine and the UI Composer can be scaled, tested, and optimized independently.

## AI Constraints (Security & Determinism)

The AI (and UI Composer) are strictly prohibited from generating raw presentation code:
*   NO HTML
*   NO CSS / Tailwind classes
*   NO React Components
*   NO Flutter Widgets
*   NO JavaScript / DOM manipulation

The protocol relies entirely on typed Nodes, Props, Actions, and Children.
