# System Architecture

Version: 1.0

---

# Overview

SHOPWISE adopts a distributed architecture where business services, AI orchestration, and retailer integrations are separated into independent runtime components.

The system is designed around three architectural principles:

- AI-native
- Retail-agnostic
- Service-oriented

---

# Context Diagram

```mermaid
flowchart TD
    Customer([Customer]) --> App[React Web / Mobile App]
    App --> Backend[Go Backend API]

    Backend --> BA_API[Business API]
    Backend --> SA_API[Session API]
    Backend --> EE_API[Event API]

    BA_API --> AI[Python AI Service]
    SA_API --> AI[Python AI Service]
    EE_API --> AI[Python AI Service]
    AI --> Planner[Planner]
    AI --> ToolRuntime[Tool Runtime]
    AI --> Memory[Memory]

    ToolRuntime --> Retail[Retail Connector]
    Retail --> PV[Phong Vu]
    Retail --> Shopee[Shopee]
    Retail --> CellphoneS[CellphoneS]
```

---

# Architectural Principles

## Separation of Concerns

Business logic and AI orchestration are deployed independently.

---

## Retail Independence

No service depends directly on retailer-specific APIs.

---

## Tool-based Integration

AI communicates exclusively through tools.

---

## UI Independence

Frontend owns rendering.

AI owns interaction intent.

---

## Stateless AI Runtime

The AI service is stateless.

Persistent data is managed by backend services.
