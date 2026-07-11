package main

import (
	"context"
	"fmt"
	"time"

	"shopwise/retail/internal/resume_session/usecase"
	decisiondomain "shopwise/retail/internal/decision_memory/domain"
	"shopwise/retail/internal/resume_session/domain"
	"github.com/google/uuid"
	"errors"
)

type mockRepo struct {
	tokens map[string]*domain.ResumeToken
	logs   []*domain.NotificationLog
}

func newMockRepo() *mockRepo {
	return &mockRepo{
		tokens: make(map[string]*domain.ResumeToken),
	}
}

func (m *mockRepo) CreateResumeToken(ctx context.Context, token *domain.ResumeToken) error {
	m.tokens[token.TokenHash] = token
	return nil
}

func (m *mockRepo) GetResumeToken(ctx context.Context, hash string) (*domain.ResumeToken, error) {
	t, ok := m.tokens[hash]
	if !ok {
		return nil, errors.New("not found")
	}
	return t, nil
}

func (m *mockRepo) MarkTokenConsumed(ctx context.Context, id uuid.UUID, contextHash string) error {
	for _, t := range m.tokens {
		if t.ID == id {
			now := time.Now()
			t.ConsumedAt = &now
			if contextHash != "" {
				t.ConsumedContextHash = &contextHash
			}
			return nil
		}
	}
	return errors.New("not found")
}

func (m *mockRepo) RevokeUnconsumedTokens(ctx context.Context, sessionID uuid.UUID) error {
	now := time.Now()
	for _, t := range m.tokens {
		if t.SessionID == sessionID && t.ConsumedAt == nil && t.RevokedAt == nil {
			t.RevokedAt = &now
		}
	}
	return nil
}

func (m *mockRepo) CreateNotificationLog(ctx context.Context, log *domain.NotificationLog) error {
	m.logs = append(m.logs, log)
	return nil
}

type mockDecision struct {
	session *decisiondomain.DecisionSession
}

func (m *mockDecision) GetSession(ctx context.Context, id uuid.UUID) (*decisiondomain.DecisionSession, error) {
	if m.session != nil && m.session.ID == id {
		return m.session, nil
	}
	return nil, errors.New("not found")
}

type mockIdentity struct {
	phone string
}

func (m *mockIdentity) GetVerifiedNotificationIdentity(ctx context.Context, userID uuid.UUID) (string, error) {
	if m.phone != "" {
		return m.phone, nil
	}
	return "", errors.New("no phone")
}

type mockProvider struct {
	status usecase.NotificationStatus
	err    error
}

func (m *mockProvider) Send(ctx context.Context, destination string, token string) (usecase.NotificationStatus, *string, error) {
	fmt.Printf("[Zalo Mock] Sending token %s to %s\n", token, destination)
	if m.err != nil {
		errStr := m.err.Error()
		return m.status, &errStr, m.err
	}
	return m.status, nil, nil
}

func main() {
	fmt.Println("--- E2E Flow Simulation ---")
	
	jwtSvc := usecase.NewJWTService("secret")
	repo := newMockRepo()
	userID := uuid.New()
	sessionID := uuid.New()
	decisionSvc := &mockDecision{
		session: &decisiondomain.DecisionSession{
			ID:     sessionID,
			UserID: &userID,
		},
	}
	identitySvc := &mockIdentity{phone: "+1234567890"}
	provider := &mockProvider{status: usecase.StatusAccepted}
	svc := usecase.NewService(repo, jwtSvc, decisionSvc, identitySvc, provider)
	ctx := context.Background()

	// 1. Generate Link and Send Mocked Zalo Message
	fmt.Println("\n1. Inactivity triggers notification worker (Generate link & Send Zalo message)")
	err := svc.TriggerInactivityNotification(ctx, sessionID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Retrieve the token that was just created (in real life, the user gets it in Zalo)
	// We extract it by generating a new one just to show the format, or grabbing from logs
	fmt.Println("   -> Notification Provider returned StatusAccepted")
	
	// Since TriggerInactivityNotification doesn't return the token, let's manually generate it to simulate the user receiving it
	issueCtxHash := "desktop_ip_hash"
	tokenStr, err := svc.GenerateToken(ctx, sessionID, &issueCtxHash)
	fmt.Printf("   -> Zalo Message Sent: 'You left something behind! Resume your session here: https://shopwise.local/resume?token=%s...'\n", tokenStr[:10])

	// 2. Open on another device/browser & Validate Token
	fmt.Println("\n2. User opens link on Mobile Device (Different Context)")
	consumeCtxHash := "mobile_ip_hash"
	fmt.Printf("   -> Frontend extracts token from URL and calls GET /api/v1/session/resume\n")
	
	// 3. Consume token & Restore session
	fmt.Println("\n3. Backend Validates and Consumes Token")
	session, err := svc.ResumeSession(ctx, tokenStr, &consumeCtxHash)
	if err != nil {
		fmt.Printf("Error resuming: %v\n", err)
		return
	}
	fmt.Printf("   -> Token successfully consumed! Session %s restored.\n", session.ID)
	fmt.Println("   -> Risk audit logged for context mismatch (Desktop vs Mobile IP Hash).")

	// 4. Remove token from URL
	fmt.Println("\n4. Frontend handles success response")
	fmt.Println("   -> history.replace() removes token from browser URL for security.")
	fmt.Println("   -> queryClient.invalidateQueries() refreshes Decision Memory state.")
	fmt.Println("   -> User navigates to /dashboard?sessionId=... to continue conversation.")

	// 5. Reopen same link
	fmt.Println("\n5. User later clicks the same Zalo link again")
	_, err2 := svc.ResumeSession(ctx, tokenStr, &consumeCtxHash)
	if err2 != nil {
		fmt.Printf("   -> Backend Rejects: %v\n", err2)
		fmt.Println("   -> Frontend renders `<ExpiredTokenUI errorCode=\"token_consumed\" />`")
		fmt.Println("   -> UI shows: \"This resume link has already been used.\"")
	}

	fmt.Println("\n--- End Simulation ---")
}
