# 004 — Fix Audit Trail Easing and Duration

- **Status**: TODO
- **Commit**: ce8c2ee
- **Severity**: MEDIUM
- **Category**: Performance & Timing
- **Estimated scope**: `apps/web/src/features/dashboard/components/audit-trail.tsx`

## Problem

`transition-all duration-500` is used on interactive UI elements (like rotate triggers or chips) in the Audit Trail. 500ms is significantly too slow for utility actions, feeling sluggish rather than snappy.

```tsx
/* apps/web/src/features/dashboard/components/audit-trail.tsx:65 — current */
"transition-all duration-500 group-hover:rotate-180"

/* apps/web/src/features/dashboard/components/audit-trail.tsx:61 — current */
"transition-all hover:border-[#4F7CFF]/40"
```

## Target

Reduce duration to 200ms (or rely on default 150ms) and use a snappy ease-out curve.

```tsx
/* target audit-trail.tsx:65 */
"transition-transform duration-200 ease-[var(--ease-out)] group-hover:rotate-180"

/* target audit-trail.tsx:61 */
"transition-colors duration-200 ease-[var(--ease-out)] hover:border-[#4F7CFF]/40"
```

## Repo conventions to follow

- Replace `transition-all` with specific transitions (`transition-transform`, `transition-colors`).
- Use the new `--ease-out` token for UI entrances.

## Steps

1. Edit `audit-trail.tsx`.
2. Replace `transition-all duration-500` with `transition-transform duration-200 ease-[var(--ease-out)]`.
3. Replace generic `transition-all` on chips with `transition-[transform,colors] duration-200 ease-[var(--ease-out)]`.

## Boundaries

- Do NOT change component layout.

## Verification

- **Mechanical**: Lint check passes.
- **Feel check**: Hover over the refresh trigger in the audit trail.
  - The icon should rotate briskly (200ms) with a natural deceleration (ease-out).
- **Done when**: No `duration-500` remains on interactive UI hovers in this file.
