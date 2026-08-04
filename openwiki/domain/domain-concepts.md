# Domain Concepts

This page captures the main product and business concepts used across SHOPWISE.

## Decision session

A **decision session** is the core unit of shopping continuity.

In current code (`/services/main-backend/internal/decision_memory/usecase/service.go` and `/apps/web/src/api/decision-memory.ts`), a session has:
- an ID
- either a user ID or an anonymous ID
- a title
- a status
- timestamps
- messages
- optional parent session relationships for branching

### Why it exists
The product is not just a cart or chat transcript. It wants to preserve an evolving shopping decision with:
- requirements
- compared products
- pinned items
- preference signals
- AI reasoning and supporting evidence

This is described most clearly in `/specs/005-decision-memory/spec.md`.

## Anonymous identity

Anonymous users are still allowed to have session continuity. The web client generates and persists an `anonymousId` locally, then sends it as `X-Anonymous-ID` (`/apps/web/src/api/decision-memory.ts`).

### Important invariant
Anonymous users keep a rolling window of the **10 most recent sessions**. The oldest session is evicted when a new one is created beyond the limit.

This is a core product tradeoff: preserve continuity without treating anonymous users as unlimited-storage accounts.

## User preference

A **user preference** is a reusable shopping signal extracted from prior interactions, such as budget, brand, or usage constraints.

The decision memory spec treats this as a first-class feature because the product promise is cumulative intelligence, not one-off chat.

In current code, preference listing and mutation endpoints exist through the decision memory API helper and backend session routes.

## Session branching

A **branch** is a new session created from an earlier session state without overwriting the original. The spec explicitly chooses a **full copy** model rather than a lightweight pointer-based branch.

Why this matters:
- a buyer may want to explore different budgets or constraints
- the original reasoning trail should remain intact
- new recommendations should not corrupt the original exploration path

## Resume token

A **resume token** is a secure single-use credential for reopening an interrupted session.

The resume session spec and service define the main properties:
- JWT-backed
- short-lived / time-bounded
- no PII in payload
- invalidated after successful use
- older unconsumed tokens revoked when a newer one is issued
- context mismatch logged as risk signal, but not automatically blocked

This concept matters because SHOPWISE wants cross-device continuity without exposing internal session details.

## Notification identity

The resume flow assumes a **verified notification identity** for a user, currently modeled around a Zalo destination in `/services/main-backend/internal/resume_session/usecase/service.go` and the feature spec.

This is not just a contact detail. It is the trusted outbound channel used for re-engagement.

## Product catalog item

A **product** is the comparable unit shown to the shopper. In the current implementation, the catalog handler returns product records from the backend database, while the dashboard maps those records into UI-specific comparison cards.

In the AI runtime schema (`/services/ai-runtime/src/models/schemas.py`), products are also part of the structured response contract, including:
- price
- specs
- match score
- explanation
- performance-related dimensions

This reinforces that products are represented in multiple layers:
- backend database/API
- dashboard UI card model
- AI structured response schema

## Store

A **store** is a fulfillment/retail endpoint surfaced to the shopper. The current `/stores` handler returns a static list of physical store options. This is a lightweight implementation now, but it supports the broader procurement flow around availability and pickup.

## Order / checkout

An **order** is the backend representation of a confirmed checkout action. The checkout handler expects:
- customer identity
- selected items
- fulfillment method
- shipping address
- optional coupon code

From a product perspective, checkout is downstream of the earlier reasoning flows. The user journey is meant to go from AI-guided exploration to ready-to-execute purchase.

## Trust score, audit trail, and agent hub

These are currently UI-facing domain concepts, especially on the dashboard:
- **Trust score**: a confidence-oriented summary shown to the shopper
- **Audit trail**: visible trace of what the system evaluated
- **Agent hub**: a mental model of multiple collaborating agents or specialist processes

Evidence appears in dashboard components such as:
- `/apps/web/src/features/dashboard/components/agent-hub.tsx`
- `/apps/web/src/features/dashboard/components/audit-trail.tsx`
- `/apps/web/src/features/dashboard/components/trust-center-header.tsx`

These concepts matter because the product is trying to make AI-assisted decision-making inspectable, not opaque.

## Reasoning replay

A **reasoning replay** is the ability to review why the assistant recommended something. This appears in both the decision memory spec and dashboard UX concepts.

Even where the full backend implementation is still evolving, this is a core business promise: recommendations should be explainable enough to revisit later.

## Practical concept map

If you are reading code, think of the core model like this:

- **Session** is the top-level shopping container
- **Messages** record the dialogue inside it
- **Preferences** are reusable user signals derived from or attached to sessions
- **Branching** creates alternate futures from a session
- **Resume tokens** reopen sessions securely
- **Products/stores/orders** are commerce entities the session reasons about
- **Trust/audit/reasoning** are transparency layers that make the AI workflow legible to users

## Watch-outs for future changes

- Do not confuse product intent from specs with full implementation coverage in code.
- Preserve the session identity rules: authenticated and anonymous paths behave differently for storage and retention.
- Treat resume tokens as security-critical, not as convenience URLs.
- If you change structured AI response fields, verify they still support trust/audit/agent-oriented UI concepts rather than only plain chat text.
