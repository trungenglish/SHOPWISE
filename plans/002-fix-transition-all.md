# 002 — Fix transition-all in UI Primitives

- **Status**: TODO
- **Commit**: ce8c2ee
- **Severity**: HIGH
- **Category**: Performance
- **Estimated scope**: ~10 files in `packages/ui/src/components`

## Problem

Rampant use of `transition-all` on components like `toggle.tsx`, `switch.tsx`, `tabs.tsx`, and `multi-select.tsx`. This animates unintended properties off-GPU (like `border-radius`, `box-shadow`, `width`, `height`), leading to dropped frames on lower-end devices.

```tsx
/* packages/ui/src/components/toggle.tsx:8 — current */
"group/toggle inline-flex items-center justify-center gap-1 rounded-md text-xs font-medium whitespace-nowrap transition-all outline-none hover:bg-muted"

/* packages/ui/src/components/multi-select.tsx:51 — current */
const multiSelectVariants = cva("m-1 transition-all duration-300 ease-in-out", {
```

## Target

Replace `transition-all` with specific properties (e.g., `transition-colors`, `transition-transform`) and replace slow implicit defaults with `--ease-out` and 200ms durations.

```tsx
/* target toggle.tsx */
"group/toggle inline-flex items-center justify-center gap-1 rounded-md text-xs font-medium whitespace-nowrap transition-colors duration-200 ease-[var(--ease-out)] outline-none hover:bg-muted"

/* target multi-select.tsx */
const multiSelectVariants = cva("m-1 transition-[color,background-color,border-color,text-decoration-color,fill,stroke,transform,opacity] duration-200 ease-[var(--ease-out)]", {
```

## Repo conventions to follow

- Modify the Tailwind utility string in the `cva` or `className` prop directly.
- Use `transition-colors` or `transition-[transform,opacity]` explicitly instead of `transition-all`.

## Steps

1. Find all instances of `transition-all` in `packages/ui/src/components/`.
2. Replace with `transition-colors duration-200 ease-[var(--ease-out)]` (or specific properties needed).
3. If structural properties (like width/height) actually need animating, move them to GPU-accelerated transforms (`scale`) or isolate them in separate `transition-[width]` utilities.

## Boundaries

- Do NOT change markup/structure — motion properties only.
- Do NOT change component logic.

## Verification

- **Mechanical**: `pnpm dlx ultracite check`
- **Feel check**: Hover over a toggle or multi-select.
  - The color change should happen quickly (200ms) with a snappy ease-out curve.
  - In DevTools, set playback to 10% and confirm only `background-color` or `color` animates, no layout shifts occur.
- **Done when**: `grep -r "transition-all" packages/ui/src/components` returns zero results.
