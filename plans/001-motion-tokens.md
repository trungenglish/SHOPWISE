# 001 — Establish Motion Tokens

- **Status**: TODO
- **Commit**: ce8c2ee
- **Severity**: HIGH
- **Category**: Cohesion & Tokens
- **Estimated scope**: 2 files (globals.css, index.css)

## Problem

The SHOPWISE project currently lacks shared motion tokens. Developers rely on default Tailwind animations (which use `ease-in-out` implicitly) or hardcode sluggish durations. This breaks the cohesion between the application UI (which should be snappy and utility-focused) and the landing page (which is expressive).

```css
/* apps/web/src/index.css — current */
@theme {
  /* No motion tokens defined */
}
```

## Target

Introduce standardized easing curves using Emil Kowalski's recommended values.

```css
@theme {
  /* UI Transitions (App Experience) */
  --ease-out: cubic-bezier(0.23, 1, 0.32, 1);
  --ease-in-out: cubic-bezier(0.77, 0, 0.175, 1);
  --ease-drawer: cubic-bezier(0.32, 0.72, 0, 1);
  
  /* Landing Page Expressiveness */
  --ease-fluid: cubic-bezier(0.4, 0, 0, 1);
  --spring-bounce: cubic-bezier(0.175, 0.885, 0.32, 1.275);
}
```

## Repo conventions to follow

- Tokens live in `packages/ui/src/styles/globals.css` (for the UI package) and `apps/web/src/index.css` (for app-specific styling) inside the `@theme` block or `:root`.

## Steps

1. Edit `packages/ui/src/styles/globals.css` to add the custom `--ease-*` variables to the `:root` scope or `@theme` block (following Tailwind v4 structure).
2. Edit `apps/web/src/index.css` to ensure they are available project-wide if not inherited correctly.

## Boundaries

- Do NOT change any components in this step. This is purely setting up the variables.
- Do NOT add new dependencies.

## Verification

- **Mechanical**: Run `pnpm run build` and ensure CSS compiles correctly.
- **Feel check**: Verify that `var(--ease-out)` is recognized by the browser inspector on any element you manually apply it to.
- **Done when**: The CSS variables are present in the final bundled CSS.
