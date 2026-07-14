# 005 — Spatial Workspace Transitions

- **Status**: TODO
- **Commit**: ce8c2ee
- **Severity**: LOW (Missed Opportunity)
- **Category**: Polish
- **Estimated scope**: `apps/web/src/features/dashboard/components/spatial-workspace.tsx`

## Problem

The transition between selecting different products in the 3D workspace currently snaps harshly or abruptly swaps out data. A shared element transition or a subtle mask would make the swap feel physical and continuous.

## Target

Apply a Framer Motion `AnimatePresence` and `motion.div` around the product display card, using a subtle `blur` filter and `opacity` crossfade during swaps.

```tsx
/* target spatial-workspace.tsx */
import { motion, AnimatePresence } from "framer-motion";

<AnimatePresence mode="wait">
  {displayProduct ? (
    <motion.div
      key={displayProduct.id}
      initial={{ opacity: 0, filter: "blur(4px)", scale: 0.98 }}
      animate={{ opacity: 1, filter: "blur(0px)", scale: 1 }}
      exit={{ opacity: 0, filter: "blur(4px)", scale: 0.98 }}
      transition={{ duration: 0.2, ease: [0.23, 1, 0.32, 1] }}
      className="glass-card absolute top-4 left-4 z-20 flex w-[300px] flex-col gap-4 rounded-2xl border p-5 shadow-2xl"
    >
      {/* ... */}
    </motion.div>
  ) : null}
</AnimatePresence>
```

## Repo conventions to follow

- Use `framer-motion` for layout and complex state transitions that Tailwind cannot handle easily.
- Keep duration short (0.2s) since this is within the application workspace.

## Steps

1. Import `motion` and `AnimatePresence` from `framer-motion` in `spatial-workspace.tsx`.
2. Wrap the `displayProduct` card in `<AnimatePresence mode="wait">`.
3. Change the card `div` to `motion.div` and apply the `initial`, `animate`, and `exit` props.

## Boundaries

- Do NOT change the 3D canvas logic.
- Do NOT change the product data model.

## Verification

- **Mechanical**: TypeScript compilation passes.
- **Feel check**: Click on a node in the spatial workspace to open the product details.
  - The card should smoothly fade and unblur in.
  - Clicking another node should fade out the old card and blur in the new one without snapping.
- **Done when**: Product switching feels continuous and physical.
