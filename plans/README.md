# Animation Improvement Plans

These plans were generated to unify and improve the motion design system of SHOPWISE, balancing landing page expressiveness with application utility.

## Execution Order

| Plan | Title | Severity | Status | Dependencies |
| :--- | :--- | :--- | :--- | :--- |
| `001-motion-tokens.md` | Establish Motion Tokens | HIGH | DONE | None (Must run first) |
| `002-fix-transition-all.md` | Fix transition-all in UI Primitives | HIGH | DONE | 001 |
| `003-gate-hover-animations.md` | Gate Hover Animations for Touch Devices | MEDIUM | DONE | None |
| `004-speed-up-audit-trail.md` | Fix Audit Trail Easing and Duration | MEDIUM | DONE | 001 |
| `005-spatial-workspace-transition.md` | Spatial Workspace Transitions | LOW | DONE | None |
| `006-tactile-buttons.md` | Tactile Button Press Feedback | LOW | DONE | 001 |

## Guidelines for Execution

1. Start with **001**. The motion tokens defined there are referenced by almost all subsequent plans (`var(--ease-out)`, etc.).
2. Plans **002** and **004** directly depend on the tokens being present to avoid breaking animations or using sluggish defaults.
3. Plans **003**, **005**, and **006** can be executed in any order, though **006** also leverages the motion tokens for its press-release curve.
