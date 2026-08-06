package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	decisiondomain "shopwise/retail/internal/decision_memory/domain"
	"shopwise/retail/internal/resume_session/domain"
	"shopwise/retail/internal/resume_session/usecase"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	if m.err != nil {
		errStr := m.err.Error()
		return m.status, &errStr, m.err
	}
	return m.status, nil, nil
}

func TestResumeSession(t *testing.T) {
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

	// Generate Token
	issueCtxHash := "hash123"
	tokenStr, err := svc.GenerateToken(ctx, sessionID, &issueCtxHash)
	require.NoError(t, err)
	require.NotEmpty(t, tokenStr)

	// Valid Resume
	consumeCtxHash := "hash123"
	session, err := svc.ResumeSession(ctx, tokenStr, &consumeCtxHash)
	require.NoError(t, err)
	assert.Equal(t, sessionID, session.ID)

	// Consumed Token Rejected
	_, err = svc.ResumeSession(ctx, tokenStr, &consumeCtxHash)
	assert.ErrorIs(t, err, usecase.ErrTokenConsumed)

	// Revoked Token Rejected
	tokenStr2, err := svc.GenerateToken(ctx, sessionID, &issueCtxHash)
	require.NoError(t, err)
	// generate another token, revoking tokenStr2
	tokenStr3, err := svc.GenerateToken(ctx, sessionID, &issueCtxHash)
	require.NoError(t, err)

	_, err = svc.ResumeSession(ctx, tokenStr2, &consumeCtxHash)
	assert.ErrorIs(t, err, usecase.ErrTokenRevoked)

	// tokenStr3 is still valid
	_, err = svc.ResumeSession(ctx, tokenStr3, &consumeCtxHash)
	require.NoError(t, err)

	// Invalid signature
	_, err = svc.ResumeSession(ctx, "invalid.token.str", &consumeCtxHash)
	assert.ErrorIs(t, err, usecase.ErrTokenInvalid)

	// Different IP does not hard block
	tokenStr4, err := svc.GenerateToken(ctx, sessionID, &issueCtxHash)
	require.NoError(t, err)
	diffCtxHash := "different_hash"
	session4, err := svc.ResumeSession(ctx, tokenStr4, &diffCtxHash)
	require.NoError(t, err)
	assert.Equal(t, sessionID, session4.ID)

	// Expired token rejected even when signature and database record are valid.
	expiredToken, err := svc.GenerateToken(ctx, sessionID, &issueCtxHash)
	require.NoError(t, err)
	for _, storedToken := range repo.tokens {
		if storedToken.ConsumedAt == nil && storedToken.RevokedAt == nil {
			storedToken.ExpiresAt = time.Now().UTC().Add(-time.Minute)
		}
	}
	_, err = svc.ResumeSession(ctx, expiredToken, &consumeCtxHash)
	assert.ErrorIs(t, err, usecase.ErrTokenExpired)
}

func TestTriggerInactivityNotification(t *testing.T) {
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

	err := svc.TriggerInactivityNotification(ctx, sessionID)
	require.NoError(t, err)
	assert.Len(t, repo.logs, 1)

	// Missing authenticated user
	decisionSvc.session.UserID = nil
	err = svc.TriggerInactivityNotification(ctx, sessionID)
	assert.ErrorIs(t, err, usecase.ErrNoVerifiedIdentity)

	// Unverified phone
	decisionSvc.session.UserID = &userID
	identitySvc.phone = ""
	err = svc.TriggerInactivityNotification(ctx, sessionID)
	assert.ErrorIs(t, err, usecase.ErrNoVerifiedIdentity)
}
