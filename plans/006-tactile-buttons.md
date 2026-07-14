# 006 — Tactile Button Press Feedback

- **Status**: TODO
- **Commit**: ce8c2ee
- **Severity**: LOW (Missed Opportunity)
- **Category**: Polish
- **Estimated scope**: `packages/ui/src/components/button.tsx`

## Problem

Buttons currently use `active:not-aria-[haspopup]:translate-y-px`. This provides a very subtle 1px shift down, but lacks the tactile scale response that feels modern and responsive, particularly on touch devices or primary actions.

```tsx
/* packages/ui/src/components/button.tsx:8 — current */
"active:not-aria-[haspopup]:translate-y-px"
```

## Target

Replace the `translate-y-px` with a scale down `active:scale-[0.98]` and ensure it has a snappy transition back to 1.

```tsx
/* target button.tsx */
"active:not-aria-[haspopup]:scale-[0.98] transition-[color,transform,background-color,border-color,opacity] duration-200 ease-[var(--ease-out)]"
```

## Repo conventions to follow

- Keep the `not-aria-[haspopup]` selector to avoid scaling buttons that open dropdowns (which feels weird when the dropdown anchors to them).
- Maintain existing `disabled` and `focus-visible` styles.

## Steps

1. Edit `packages/ui/src/components/button.tsx`.
2. Replace `active:not-aria-[haspopup]:translate-y-px` with `active:not-aria-[haspopup]:scale-[0.98]`.
3. Ensure the button has `transition-transform` enabled.

## Boundaries

- Do NOT change button variants or sizes.
- Do NOT add new libraries.

## Verification

- **Mechanical**: Lint check passes.
- **Feel check**: Click and hold a primary button.
  - It should shrink slightly (0.98 scale) immediately.
  - Releasing it should snap back to normal size with a crisp curve.
- **Done when**: All buttons exhibit scale-based press feedback instead of Y-translation.
