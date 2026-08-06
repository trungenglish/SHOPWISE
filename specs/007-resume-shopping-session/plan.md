# Implementation Plan: Resume Shopping Session

**Feature Directory**: specs/007-resume-shopping-session
**Status**: Draft

## Technical Architecture

### Frontend (React + Vite + TypeScript)
- The frontend will handle routing for the resume link (`/resume?token=...`).
- A new `ResumeSession` component will handle the loading state, parse the token, and call the backend to restore the session.
- Handles error states: Expired Token UI, Invalid Token UI, Network Error UI.
- No business logic; completely driven by the backend response.

### Backend (Go + Gin)
- **Token Generation**: A service to generate JWT-based, single-use, 24-hour bearer tokens containing the `session_id`.
- **Notification Provider**: A `NotificationProvider` interface with a `ZaloNotificationProvider` implementation.
  - Method `Send(ctx, token, destination)` returning `Accepted`, `Failed`, or `RetryableFailure`.
- **Resume Logic**: The `GET /api/v1/session/resume` endpoint validates the token signature, checks if it has been consumed in the database, marks it as consumed, and returns the serialized `DecisionSession` data.
- **Inactivity Detection**: For MVP, a background worker or cron that polls active, uncompleted sessions without recent activity (>10 mins) and triggers the notification flow.

### AI Runtime (Python + FastAPI)
- Remains completely stateless. The restored session context is passed to the AI on the next user interaction, exactly like a normal continuation of the conversation.

## Error Handling Strategy
- **Transient Notification Failures**: The backend handles `RetryableFailure` by retrying once with an exponential backoff.
- **Permanent Notification Failures**: The backend logs the failure (e.g., invalid phone number) and halts the resume attempt for that session.
- **Invalid/Expired Tokens**: The backend rejects the resume request. The frontend catches the 401/403 or custom error code and renders an "Expired Link" screen prompting the user to start a new session or request a new link.

## Security Considerations
- **Single-Use**: Token is invalidated in the database upon successful validation (recording `consumed_at` and `revoked_at`).
- **Expiration**: 24-hour expiration.
- **JWT**: Tokens are verified via JWT signature. Required claims: `sub`, `session_id`, `exp`, `iat`, `jti`, and `issuer` or `audience`.
- **Context Logging**: Risk-aware device/IP context logging using privacy-preserving hashes. No strict IP binding for cross-device resume.
- **Audit Logging**: Comprehensive auditing for issuance, consumption, expiration, revocation, and suspicious context mismatches.

## Testing Strategy
- **Unit Tests**:
  - NotificationProvider interface mocking and Zalo implementation (mocked HTTP client).
  - JWT generation and validation logic (expiration, signing).
- **Integration Tests**:
  - Full flow: Generate token -> validate token -> token is consumed -> second validation fails (replay prevention).
  - Test endpoint `GET /api/v1/session/resume` with valid, expired, and invalid tokens.

## Implementation Phases
1. **Phase 1: Data Model & Core API**: Implement JWT token generation, validation, and database persistence (consumed tokens). Add `GET /api/v1/session/resume` endpoint.
2. **Phase 2: Notification Provider**: Implement `NotificationProvider` interface and `ZaloNotificationProvider`. Add retry logic for transient failures.
3. **Phase 3: Trigger Mechanism**: Implement the inactivity detection to trigger the notification flow.
4. **Phase 4: Frontend**: Implement the `/resume` route, loading state, error states, and session restoration logic using the new endpoint.
