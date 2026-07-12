package usecase

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"shopwise/retail/internal/resume_session/domain"
	decisiondomain "shopwise/retail/internal/decision_memory/domain"

	"github.com/google/uuid"
)
var (
	ErrTokenExpired       = errors.New("token expired")
	ErrTokenConsumed      = errors.New("token already consumed")
	ErrTokenRevoked       = errors.New("token revoked")
	ErrTokenInvalid       = errors.New("invalid token")
	ErrNoVerifiedIdentity = errors.New("no verified notification identity for session")
)

type DecisionService interface {
	GetSession(ctx context.Context, id uuid.UUID) (*decisiondomain.DecisionSession, error)
}

type IdentityService interface {
	GetVerifiedNotificationIdentity(ctx context.Context, userID uuid.UUID) (string, error)
}

type Service struct {
	repo       Repository
	jwtSvc     *JWTService
	decisionSvc DecisionService
	identitySvc IdentityService
	provider   NotificationProvider
}

func NewService(repo Repository, jwtSvc *JWTService, decisionSvc DecisionService, identitySvc IdentityService, provider NotificationProvider) *Service {
	return &Service{
		repo:       repo,
		jwtSvc:     jwtSvc,
		decisionSvc: decisionSvc,
		identitySvc: identitySvc,
		provider:   provider,
	}
}

// GenerateToken creates and stores a new resume token for a session
func (s *Service) GenerateToken(ctx context.Context, sessionID uuid.UUID, contextHash *string) (string, error) {
	session, err := s.decisionSvc.GetSession(ctx, sessionID)
	if err != nil {
		return "", err
	}

	if err := s.repo.RevokeUnconsumedTokens(ctx, sessionID); err != nil {
		return "", fmt.Errorf("revoke unconsumed tokens: %w", err)
	}

	tokenStr, jti, expiresAt, err := s.jwtSvc.IssueResumeToken(sessionID, session.UserID)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256([]byte(tokenStr))
	hashStr := hex.EncodeToString(hash[:])

	token := &domain.ResumeToken{
		ID:                jti,
		SessionID:         sessionID,
		UserID:            session.UserID,
		TokenHash:         hashStr,
		ExpiresAt:         expiresAt,
		IssuedContextHash: contextHash,
		CreatedAt:         time.Now().UTC(),
	}

	if err := s.repo.CreateResumeToken(ctx, token); err != nil {
		return "", err
	}

	return tokenStr, nil
}

// ResumeSession validates the token, marks it as consumed, and returns the session
func (s *Service) ResumeSession(ctx context.Context, tokenStr string, contextHash *string) (*decisiondomain.DecisionSession, error) {
	sessionID, jti, err := s.jwtSvc.Verify(tokenStr)
	if err != nil {
		return nil, ErrTokenInvalid
	}

	hash := sha256.Sum256([]byte(tokenStr))
	hashStr := hex.EncodeToString(hash[:])

	token, err := s.repo.GetResumeToken(ctx, hashStr)
	if err != nil {
		return nil, ErrTokenInvalid
	}

	if token.ID != jti || token.SessionID != sessionID {
		return nil, ErrTokenInvalid
	}

	if token.RevokedAt != nil {
		return nil, ErrTokenRevoked
	}

	if token.ConsumedAt != nil {
		return nil, ErrTokenConsumed
	}

	if time.Now().UTC().After(token.ExpiresAt) {
		return nil, ErrTokenExpired
	}

	// Context auditing (risk signal)
	if token.IssuedContextHash != nil && contextHash != nil {
		if *token.IssuedContextHash != *contextHash {
			fmt.Printf("[AUDIT RISK] Session %s resumed with context mismatch. Issued: %s, Consumed: %s\n", sessionID, *token.IssuedContextHash, *contextHash)
		}
	}

	contextHashStr := ""
	if contextHash != nil {
		contextHashStr = *contextHash
	}

	if err := s.repo.MarkTokenConsumed(ctx, token.ID, contextHashStr); err != nil {
		return nil, err
	}

	session, err := s.decisionSvc.GetSession(ctx, sessionID)
	if err != nil {
		return nil, err
	}

	return session, nil
}

// TriggerInactivityNotification is a helper to manually trigger the resume flow for a specific session.
func (s *Service) TriggerInactivityNotification(ctx context.Context, sessionID uuid.UUID) error {
	session, err := s.decisionSvc.GetSession(ctx, sessionID)
	if err != nil {
		return fmt.Errorf("get session: %w", err)
	}

	if session.UserID == nil {
		return ErrNoVerifiedIdentity
	}

	destination, err := s.identitySvc.GetVerifiedNotificationIdentity(ctx, *session.UserID)
	if err != nil || destination == "" {
		return ErrNoVerifiedIdentity
	}

	tokenStr, err := s.GenerateToken(ctx, sessionID, nil)
	if err != nil {
		return fmt.Errorf("generate token: %w", err)
	}

	status, errDetails, err := s.provider.Send(ctx, destination, tokenStr)
	
	log := &domain.NotificationLog{
		ID:        uuid.New(),
		SessionID: sessionID,
		Provider:  "Zalo",
		Status:    string(status),
		CreatedAt: time.Now().UTC(),
	}
	if errDetails != nil {
		log.ErrorDetails = errDetails
	}

	// Persist the log
	_ = s.repo.CreateNotificationLog(ctx, log)

	if err != nil {
		return fmt.Errorf("send notification: %w", err)
	}

	return nil
}
