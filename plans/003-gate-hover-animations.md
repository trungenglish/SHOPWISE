# 003 — Gate Hover Animations for Touch Devices

- **Status**: TODO
- **Commit**: ce8c2ee
- **Severity**: MEDIUM
- **Category**: Accessibility / UX
- **Estimated scope**: Multiple files in `apps/web/src/features/dashboard/components`

## Problem

Interactive cards and buttons in the dashboard (e.g., `sidebar.tsx`, `accessories-modal.tsx`) use `hover:scale-[1.02]` or `hover:scale-105` without pointer gating. On mobile devices, tapping these elements triggers the hover state permanently, which creates visual bugs and forces the user to double-tap or tap elsewhere to remove the hover state.

```tsx
/* apps/web/src/features/dashboard/components/sidebar.tsx:54 — current */
"hover:scale-[1.02] hover:border-[#4F7CFF]/60 hover:bg-[#4F7CFF]/15 active:scale-[0.98]"
```

## Target

Gate all hover scaling behind the `@media (hover: hover) and (pointer: fine)` query. In Tailwind CSS, this can be done natively or with a specific prefix. Since we are using Tailwind CSS v4, we can wrap the hover utility or rely on native touch support configurations if active, but the most explicit way is using `@media (hover: hover)` for transforms.

```tsx
/* target sidebar.tsx */
"max-md:hover:scale-100 hover:scale-[1.02] hover:border-[#4F7CFF]/60 hover:bg-[#4F7CFF]/15 active:scale-[0.98]"
```
*Note: We can also use a custom variant or just ensure Tailwind is configured with `future: { hoverOnlyWhenSupported: true }` (if v3), or the equivalent v4 native handling.*

## Repo conventions to follow

- Do not break existing desktop hovers.
- Ensure tap feedback (`active:scale-[0.98]`) still fires on mobile (that is the correct tactile response).

## Steps

1. Find all `hover:scale-*` and `hover:translate-*` classes in the dashboard components.
2. Ensure they are correctly reset or ignored on mobile, either by adding `max-md:hover:scale-100` (or `max-md:hover:transform-none`) to override the transform on touch devices, or by wrapping in a custom `touch-hover:` variant if established.
3. Keep `active:scale-*` on all items for touch feedback.

## Boundaries

- Do NOT change structural classes.

## Verification

- **Mechanical**: None.
- **Feel check**: Open Chrome DevTools, toggle Device Toolbar (mobile simulation).
  - Tap on the sidebar items. They should scale down (active) and return to normal size (no hover stuck).
- **Done when**: No sticky hover transforms occur on touch simulation.
