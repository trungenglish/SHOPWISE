---
trigger: glob
description: React + Vite — route to the matching skill in .agents/skills before implementing
globs: apps/web/**/*.{tsx,ts,jsx,js}, packages/ui/**/*.{tsx,ts,jsx,js}, apps/web/vite.config.ts
---

1. React & Vite skills

When editing matched files for **React + Vite** work (web app, shared UI package, Vite config):

- **Read the matching skill first** — open `.agents/skills/<skill>/SKILL.md` and follow it (including any `references/` it points to). Do not guess when a skill exists.
- **Pick the most specific skill** for the task. If several apply, start with the narrowest one, then pull in `vercel-react-best-practices` for data fetching / bundle / server-adjacent patterns or `react-render-optimization` for render profiling.
- **Profile before micro-optimizing** — for performance work, measure first; do not add `memo` / `useCallback` / `useMemo` without evidence.

2. Skill router

| Topic | Skill path |
| --- | --- |
| Bundle splitting, route-based chunks, reducing FCP/LCP | `.agents/skills/bundle-splitting/SKILL.md` |
| Client-side rendering (CSR), SPA architecture | `.agents/skills/client-side-rendering/SKILL.md` |
| Custom hooks, extracting shared stateful logic | `.agents/skills/hooks-pattern/SKILL.md` |
| Re-render reduction, memoization, state design | `.agents/skills/react-render-optimization/SKILL.md` |
| Render props for flexible composition | `.agents/skills/render-props-pattern/SKILL.md` |
| shadcn/ui — add, style, debug, `components.json` | `.agents/skills/shadcn/SKILL.md` |
| DOM components, web-in-native incremental migration | `.agents/skills/use-dom/SKILL.md` |
| Compound components, providers, composition APIs | `.agents/skills/vercel-composition-patterns/SKILL.md` |
| React/Next performance (waterfalls, bundles, RSC-adjacent) | `.agents/skills/vercel-react-best-practices/SKILL.md` |
| RN performance (lists, animations) — cross-platform only | `.agents/skills/vercel-react-native-skills/SKILL.md` |
| Virtual lists / windowing for large datasets | `.agents/skills/virtual-lists/SKILL.md` |
| Vite build config, code splitting, dependency tuning | `.agents/skills/vite-bundle-optimization/SKILL.md` |

If no skill matches, say so and fall back to React / Vite docs.