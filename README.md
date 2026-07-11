# SHOPWISE

This project was created with [Better-T-Stack](https://github.com/AmanVarshney01/create-better-t-stack), a modern TypeScript stack that combines React, TanStack Router, and more.

## Features

- **TypeScript** - For type safety and improved developer experience
- **TanStack Router** - File-based routing with full type safety
- **React Native** - Build mobile apps using React
- **Expo** - Tools for React Native development
- **TailwindCSS** - Utility-first CSS for rapid UI development
- **Shared UI package** - shadcn/ui primitives live in `packages/ui`
- **Turborepo** - Optimized monorepo build system
- **Go Gin API** - Backend server with GORM, Postgres, Redis, and Docker support

## Getting Started

First, install the dependencies:

```bash
pnpm install
```

### Backend (Go Gin)

Start Postgres and Redis:

```bash
pnpm --filter server dev:deps
```

Copy the server environment file and start the API:

```bash
cp apps/server/.env.example apps/server/.env
pnpm dev:server
```

Or run the full Docker stack (Postgres, Redis, and API):

```bash
pnpm dev:server:docker
```

The API listens on [http://localhost:18080](http://localhost:18080). Port `18080` avoids a common Windows + WSL conflict where `wslrelay` hijacks `localhost:8080` and resets connections.

Swagger UI (debug mode): [http://localhost:18080/swagger/index.html](http://localhost:18080/swagger/index.html)

### Frontend

Then, run the development servers:

```bash
pnpm run dev
```

Open [http://localhost:3001](http://localhost:3001) in your browser to see the web application. Use the Expo Go app to run the mobile application.

For native development, copy the Expo environment file:

```bash
cp apps/mobile/.env.example apps/mobile/.env
```

## UI Customization

React web apps in this stack share shadcn/ui primitives through `packages/ui`.

- Change design tokens and global styles in `packages/ui/src/styles/globals.css`
- Update shared primitives in `packages/ui/src/components/*`
- Adjust shadcn aliases or style config in `packages/ui/components.json` and `apps/web/components.json`

### Add more shared components

Run this from the project root to add more primitives to the shared UI package:

```bash
npx shadcn@latest add accordion dialog popover sheet table -c packages/ui
```

Import shared components like this:

```tsx
import { Button } from "@shopwise/ui/components/button";
```

### Add app-specific blocks

If you want to add app-specific blocks instead of shared primitives, run the shadcn CLI from `apps/web`.

## Deployment (Cloudflare via Alchemy)

- Target: web
- Dev: pnpm run dev
- Deploy: pnpm run deploy
- Destroy: pnpm run destroy

For more details, see the guide on [Deploying to Cloudflare with Alchemy](https://www.better-t-stack.dev/docs/guides/cloudflare-alchemy).

## Project Structure

```
shopwise/
├── apps/            # Applications
│   ├── web/         # Frontend web application (React + TanStack Router)
│   └── mobile/      # Mobile application (React Native, Expo)
├── packages/        # Shared libraries and internal configurations
│   ├── ui/          # Shared shadcn/ui components and styles
│   ├── env/         # Shared environment validation
│   └── ...          # Other internal packages (types, schemas, config)
├── services/        # Core backend domains/engines (Go, Python)
│   ├── main-backend/# Core retail engine
│   └── agentic/     # AI reasoning and memory components
├── conformance/     # Language-agnostic test suites and verification
├── infra/           # Root-level infrastructure configurations
├── deploy/          # Deployment definitions
└── scripts/         # Developer utilities
```

## Web Environment Variables

Configure `apps/server/.env` using `apps/server/.env.example`:

| Variable          | Default                  | Description    |
| ----------------- | ------------------------ | -------------- |
| `VITE_NODE_ENV`   | `development`            |                |
| `VITE_SERVER_URL` | `http://localhost:18080` | Gin server URL |

## Server Environment Variables

Configure `apps/server/.env` using `apps/server/.env.example`:

| Variable | Default | Description |
| --- | --- | --- |
| `PORT` | `18080` | API bind port |
| `GIN_MODE` | `debug` | Gin mode (`release` in production) |
| `ALLOWED_ORIGINS` | `http://localhost:3001` | Comma-separated CORS origins |
| `DATABASE_URL` | — | Postgres connection string |
| `REDIS_URL` | — | Redis connection string |
| `LOG_LEVEL` | `info` | Structured log level |
| `APP_NAME` | `server` | Application name in logs |

## AI service Environment Variables

Configure `apps/server/.env` using `apps/server/.env.example`:

| Variable          | Default                 | Description                  |
| ----------------- | ----------------------- | ---------------------------- |
| `PORT`            | `18080`                 | API bind port                |
| `ALLOWED_ORIGINS` | `http://localhost:3001` | Comma-separated CORS origins |

## Available Scripts

- `pnpm run dev`: Start all applications in development mode
- `pnpm run build`: Build all applications
- `pnpm run dev:web`: Start only the web application
- `pnpm run dev:server`: Start only the Go API (requires Postgres + Redis)
- `pnpm run dev:server:docker`: Start Postgres, Redis, and API via Docker Compose
- `pnpm run check-types`: Check TypeScript types across all apps
- `pnpm run dev:native`: Start the React Native/Expo development server
