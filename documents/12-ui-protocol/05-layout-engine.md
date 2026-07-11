# Layout Engine

The SAUP Layout Engine is declarative and deterministic. It allows the UI Composer to structure complex multi-panel workspaces without CSS.

Layout Nodes dictate the flow of their `children`.

## Workspace
The root layout for desktop and tablet, typically defining a multi-pane environment.
```json
{
  "type": "workspace",
  "children": [
    { "type": "sidebar" },
    { "type": "page" }
  ]
}
```

## Grid & Stack
For general alignment and content grouping.

*   `stack`: Lays out children sequentially along the z-axis (overlapping) or acts as a basic container.
*   `grid`: Lays out children in a defined row/column structure.
*   `row`: Horizontal linear layout.
*   `column`: Vertical linear layout.

## Responsiveness
The AI does not specify breakpoints. The Client Renderer interprets Layout Nodes contextually (e.g., `workspace` becomes a single stack or bottom-sheet-based view on Mobile).
