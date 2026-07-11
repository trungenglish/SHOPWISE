# Polyglot Monorepo Strategy (Turborepo)

The SHOPWISE project uses [Turborepo](https://turbo.build/) to orchestrate a polyglot monolithic monorepo containing Node.js apps and packages, Go and Python services, conformance tests, and infrastructure tools.

## Workspace Structure

The project root is configured via `pnpm-workspace.yaml` and includes the following directories as native workspace packages:

- `apps/*`: Core applications including frontends (React-Vite, Expo).
- `packages/*`: Shared libraries, SDKs, UI components, and internal config (TypeScript/Node.js).
- `services/*`: Core backend domains and internal services (e.g., Python `agentic`, Go `main-backend`).
- `conformance/*`: Language-agnostic test suites and SDK bindings verification.
- `infra`: Root-level infrastructure configurations.
- `deploy`: Scripts and definitions for deployment.
- `scripts`: Developer utilities.

## Package Configuration (`package.json`)

To ensure Turborepo properly tracks and orchestrates tasks across non-JS components, a `package.json` is placed in **all** workspace directories (even for Go and Python). These `package.json` files act as "dummy" manifests that expose standardized task scripts (e.g., `"build": "go build ."`) which `turbo run build` can then execute.

By managing the monorepo this way:

1. Turborepo handles caching and dependency graphs for the entire project.
2. Cross-language dependencies are synchronized natively via standard build steps.
3. Commands like `pnpm dev` and `pnpm build` work uniformly across the entire repository.
