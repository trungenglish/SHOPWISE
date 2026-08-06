# Resume Shopping Session

**Version**: 1.0.0
**Status**: Draft
**Feature Directory**: specs/007-resume-shopping-session
**Created**: 2026-07-12
**Last Updated**: 2026-07-12

---

## Overview

The Resume Shopping Session feature allows customers to seamlessly return to an interrupted shopping journey across channels. It preserves the entire context—including conversation history, product recommendations, comparison workspaces, and checkout progress—and proactively re-engages users via Zalo Official Account (OA) API after inactivity.

## Clarifications

### Session 2026-07-12
- Q: How should the system handle a scenario where a user clicks the resume link on a different device? → A: Automatically resume the session on any device using a secure, single-use, short-lived bearer token. The token must expire after a limited time, be invalidated after successful use, and never expose sensitive information. If invalid/expired, the user must request a new link.
- Q: How long should an inactive session's data be retained in the database before it is permanently deleted? → A: 24-hour single-use token.

---

## Problem Statement

Customers often leave during product discovery or before checkout, resulting in lost sales. Without session persistence, returning users are forced to restart their entire shopping conversation and recreate their comparison workspace, creating friction and degrading the premium experience of the AI Shopping Agent.

---

## Goals

- Seamlessly preserve and restore the entire shopping session (retaining 100% of the conversation history, selected product options, comparison workspace state, and checkout readiness).
- Detect customer inactivity (configurable, default 10 minutes) before checkout.
- Proactively send a personalized, secure resume link.
- Ensure restoration happens in under 3 seconds without manual reconstruction.
- The system shall integrate with the Zalo Official Account API through a pluggable Notification Provider interface.

## Non-Goals

- Managing notifications across other platforms (e.g., SMS, email) in this iteration, though the architecture will support them.
- Exposing internal session identifiers or sensitive customer data in resume links.
- Re-engaging users who have successfully completed checkout.

---

## Target Users

| User Type | Description | Primary Need |
|---|---|---|
| Returning Shopper | A customer who started a shopping session but left before checkout | To continue their previous journey exactly where they left off without losing context or comparisons |
| Zalo User | A customer interacting via Zalo Official Account | To receive timely and secure links to return to the AI Shopping Agent |

---

## User Scenarios & Testing

### Scenario 1: Resume Session After Inactivity

**Given**: A customer provides their phone number, starts a shopping session, and becomes inactive for more than 10 minutes before completing checkout.
**When**: The system detects the inactivity timeout.
**Then**: The backend triggers a resume workflow and sends a proactive Zalo notification with a secure resume link.
**Acceptance**: Ensure a Zalo message is dispatched with the correct personalized content and resume link after exactly the configured timeout.

### Scenario 2: Successful Session Restoration

**Given**: A customer receives a Zalo notification containing a secure resume link.
**When**: The customer clicks the resume link within its validity period.
**Then**: The system restores the conversation history, dynamic UI, product comparison, and decision memory exactly as they were.
**Acceptance**: The full session state must be restored and rendered on the frontend within 3 seconds, without requiring manual user input.

### Scenario 3: Expired Session Handling

**Given**: A customer clicks a resume link that has expired.
**When**: The system attempts to resolve the session.
**Then**: The system notifies the customer that the session has expired and offers to start a new shopping session, reusing past preferences if available.
**Acceptance**: The frontend correctly renders a structured Error State UI indicating session expiration and prompts a fresh start.

---

## Functional Requirements

### Customer Identification

| ID | Requirement | Priority | Acceptance Criteria |
|---|---|---|---|
| FR-01 | Support identification via phone number (primary), name, and optional account | Must Have | Sessions are uniquely linked to a valid phone number |

### Session Persistence

| ID | Requirement | Priority | Acceptance Criteria |
|---|---|---|---|
| FR-02 | Continuously persist conversation history, preferences, AI context, recommendations, and checkout readiness on the backend with a 24-hour retention period | Must Have | Frontend reload proves backend is the sole source of truth for session state; data is automatically purged after exactly 24 hours, verifiable via TTL or cron |

### Resume Trigger

| ID | Requirement | Priority | Acceptance Criteria |
|---|---|---|---|
| FR-03 | Detect inactivity (defined as 10 consecutive minutes with no inbound user messages or UI interactions) and trigger workflow | Must Have | Timeout events reliably trigger the notification service if checkout is incomplete |

### Resume Notification

| ID | Requirement | Priority | Acceptance Criteria |
|---|---|---|---|
| FR-04 | The system shall integrate with the Zalo Official Account API through a pluggable Notification Provider interface | Must Have | Notifications are sent without exposing sensitive customer info; architecture supports future channels |
| FR-05 | Notification Provider shall return synchronous delivery results (Accepted, Failed, Retryable Failure) | Must Have | Core resume service applies a bounded exponential backoff (max 3 retries) for RetryableFailures, and halts immediately on Failed states |
| FR-05b| Fallback behavior during prolonged Provider downtime | Must Have | If Provider downtime exceeds the backoff window, abandon the proactive notification but retain session data |

### Security

| ID | Requirement | Priority | Acceptance Criteria |
|---|---|---|---|
| FR-06 | Resume links must act as secure, 24-hour single-use bearer tokens that are invalidated immediately in the database upon first successful use | Must Have | Reusing a consumed single-use link results in rejection, explicitly preventing replay attacks |
| FR-06b | Handle invalid/expired tokens by prompting the user to request a new resume link | Must Have | System rejects old tokens safely and displays UI guiding the user to request a new link |
| FR-06c | The bearer token payload must be a JWT containing required claims | Must Have | Payload must include `sub`, `session_id`, `exp`, `iat`, `jti`, and `iss`/`aud`, containing zero PII |
| FR-06d | Multiple link generation invalidates previous links | Must Have | If a user generates multiple resume links, all prior unconsumed links are immediately revoked; only the newest is valid |
| FR-06e | Risk-aware context validation | Must Have | Do not enforce strict IP/device equality for cross-device resume. Capture and log issuance and consumption context (IP hash, user-agent). Mismatch emits a risk signal but does not block. Token signature failure, expiration, revocation, or previous consumption always block. |

### Error Handling

| ID | Requirement | Priority | Acceptance Criteria |
|---|---|---|---|
| FR-07 | Handle Accepted, Failed, and RetryableFailure notification states explicitly, alongside token/session errors | Must Have | Frontend renders proper Error State UI; backend explicitly logs delivery attempts and halt conditions |
| FR-07b | Expired/invalid token screens must be accessible and actionable | Must Have | Screens meet WCAG 2.1 AA standards and provide clear recovery paths (e.g., "Request new link" button) |
| FR-07c | Comprehensive audit logging for security events | Must Have | All token consumptions, validation failures, and notification delivery states are logged for auditability |

---

## Success Criteria

| Criterion | Metric | Target |
|---|---|---|
| Session Restoration Time | Time from clicking resume link to fully interactive UI | < 3 seconds |
| Recovery Rate | Percentage of inactive sessions successfully resumed via link click | > 20% |
| Zalo Delivery Success | Percentage of resume notifications successfully delivered to Zalo | > 95% |
| State Consistency | Percentage of restored sessions missing data | 0% |

---

## Assumptions

- Customers will willingly provide their phone number when prompted by the AI.
- Zalo Official Account API is accessible and has appropriate permissions to message users proactively.
- The 10-minute timeout is appropriate for the typical shopping journey on this platform.
- Cryptographically signed links can be generated and validated securely by the backend API.
- The AI runtime remains completely stateless; all session persistence is managed exclusively by the backend database, ensuring session restoration does not depend on in-memory AI state.

## Dependencies

- Zalo Official Account (OA) API integration credentials and approval.
- Centralized session storage capable of holding complex state (e.g., Redis or Postgres JSONB).

## Risks

| Risk | Likelihood | Impact | Mitigation |
|---|---|---|---|
| Zalo API Rate Limits / Delivery Failures | Medium | High | Implement retry mechanism for transient failures and robust logging; fallback Notification Provider abstraction |
| Large Session Payload Restoration Latency | Low | Medium | Optimize backend payload sizes; implement lazy loading for images/assets upon session resume |
| Privacy/Security Breach via Resume Link | Low | High | Enforce strict cryptographic signing, expiration, and no exposure of internal IDs |

---

## Open Questions

- None at this time.
