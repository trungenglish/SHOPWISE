# Tasks: Resume Shopping Session

## Implementation Strategy
Deliver the backend foundation first (data models and core JWT logic), followed by the notification delivery mechanism (Zalo integration and triggers). Finally, integrate the frontend to handle the `/resume` route and error states.

## Phase 1: Setup
- [ ] T001 Define JWT secret and Zalo OA API credentials in `services/main-backend/.env.example`

## Phase 2: Foundational
*Goal: Establish database persistence and token generation logic.*
- [x] T002 Implement `ResumeToken` and `NotificationLog` entities in `services/main-backend/internal/resume_session/domain/domain.go`
- [x] T003 Implement Postgres auto-migrations for new entities in `services/main-backend/internal/resume_session/repository/postgres/repository.go`
- [x] T004 Implement JWT generation/validation logic in `services/main-backend/internal/resume_session/usecase/jwt.go`
- [x] T004b Implement token supersession logic in `services/main-backend/internal/resume_session/usecase/service.go` (revoke previous unconsumed tokens within the same DB transaction upon new token generation)

## Phase 3: [US1] Resume Session After Inactivity
*Goal: Detect inactive sessions and securely dispatch a single-use resume link via Zalo.*
*Test Criteria: A simulated inactivity timeout triggers a Zalo message containing a valid 24h JWT link.*
- [x] T005 [US1] Implement `NotificationProvider` interface and `RetryableProvider` in `services/main-backend/internal/resume_session/usecase/provider.go`
- [x] T006 [P] [US1] Implement `ZaloNotificationProvider` mock in `services/main-backend/internal/resume_session/provider/zalo/zalo.go`
- [x] T007 [US1] Add `TriggerInactivityNotification` cron/worker logic to `services/main-backend/internal/resume_session/usecase/service.go` (must resolve recipient strictly from verified authenticated session owner; do not trust client-supplied phone)
- [x] T008 [US1] Wire providers and handlers into application registry in `services/main-backend/internal/bootstrap/bootstrap.go`

## Phase 4: [US2] Successful Session Restoration
*Goal: Restore the user's previous session context when they click a valid resume link.*
*Test Criteria: Clicking a valid link immediately loads the prior conversation history and product comparisons.*
- [x] T009 [US2] Implement `GET /api/v1/session/resume` endpoint in `services/main-backend/internal/resume_session/handler/handler.go`
- [x] T010 [US2] Add token consumption and validation logic to `services/main-backend/internal/resume_session/usecase/service.go`
- [x] T010b [US2] Implement resume context auditing and IP hash capture in `services/main-backend/internal/resume_session/usecase/service.go` (do not reject solely on IP/device mismatch; emit risk signals and test all token states)
- [x] T011 [P] [US2] Add `/resume` route configuration to `apps/web/src/App.tsx`
- [x] T012 [US2] Implement `ResumeSession` loading component in `apps/web/src/pages/ResumeSession.tsx`
- [x] T013 [US2] Integrate session restoration API call using `fetch` or React Query in `apps/web/src/pages/ResumeSession.tsx`

## Phase 5: [US3] Expired Session Handling
*Goal: Gracefully handle expired, invalid, or already-consumed tokens.*
*Test Criteria: Clicking an expired or consumed link displays an accessible Error State UI prompting a fresh start.*
- [x] T014 [US3] Implement `ExpiredTokenUI` component in `apps/web/src/components/ExpiredTokenUI.tsx`
- [x] T015 [US3] Add error state rendering logic for 401/403 responses in `apps/web/src/pages/ResumeSession.tsx`

## Phase 6: Polish & Cross-Cutting Concerns
- [x] T016 Add comprehensive audit logging for token consumption events in `services/main-backend/internal/resume_session/usecase/service.go`
- [x] T017 Verify all WCAG 2.1 AA accessibility standards are met on `ExpiredTokenUI.tsx`

## Dependencies
- US1 depends on Foundational (Phase 2).
- US2 depends on Foundational (Phase 2).
- US3 depends on US2.
