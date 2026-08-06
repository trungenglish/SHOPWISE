package promotions

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	StartPromotion(ctx context.Context, sessionID string, userID *uuid.UUID, deviceID *string, triggerType TriggerType) (*CheckoutPromotion, error)
	GetStatus(ctx context.Context, sessionID string) (*CheckoutPromotion, error)
	HandlePaymentCompleted(ctx context.Context, sessionID string) error
	RetryVoucherIssuance(ctx context.Context, sessionID string) error
}

type serviceImpl struct {
	repo    Repository
	voucher VoucherGenerator
}

func NewService(repo Repository, voucher VoucherGenerator) Service {
	return &serviceImpl{
		repo:    repo,
		voucher: voucher,
	}
}

func (s *serviceImpl) StartPromotion(ctx context.Context, sessionID string, userID *uuid.UUID, deviceID *string, triggerType TriggerType) (*CheckoutPromotion, error) {
	// Idempotency: Check if promotion already exists for this session
	existingPromo, err := s.repo.FindBySession(ctx, sessionID)
	if err != nil {
		return nil, err
	}

	if existingPromo != nil {
		// Existing promotion found for this session, it's idempotent, just return it.
		// NOTE: Even if it's expired, we do not restart it for the SAME session.
		return existingPromo, nil
	}

	// Check Cooldown Eligibility across previous sessions for this user/device
	var uidStr, devStr *string
	if userID != nil {
		id := userID.String()
		uidStr = &id
	}
	if deviceID != nil {
		devStr = deviceID
	}

	lastPromo, err := s.repo.FindByEligibilityKey(ctx, uidStr, devStr)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	if lastPromo != nil && lastPromo.CooldownUntil != nil {
		if now.Before(*lastPromo.CooldownUntil) {
			return nil, errors.New("user is in cooldown period and not eligible for a new promotion")
		}
	}

	// Eligible, create a new promotion
	expiresAt := now.Add(15 * time.Minute)
	// Cooldown defaults to 30 days
	cooldownUntil := now.Add(30 * 24 * time.Hour)

	promo := &CheckoutPromotion{
		SessionID:     sessionID,
		UserID:        userID,
		DeviceID:      deviceID,
		TriggerType:   triggerType,
		Status:        StatusActive,
		StartedAt:     now,
		ExpiresAt:     expiresAt,
		CooldownUntil: &cooldownUntil,
		// MVP defaults are applied by DB/GORM layer (campaign code, reward details)
	}

	if err := s.repo.Create(ctx, promo); err != nil {
		return nil, err
	}

	return promo, nil
}

func (s *serviceImpl) GetStatus(ctx context.Context, sessionID string) (*CheckoutPromotion, error) {
	promo, err := s.repo.FindBySession(ctx, sessionID)
	if err != nil {
		return nil, err
	}
	if promo == nil {
		return nil, nil // No active promotion for this session
	}

	// Lazy evaluation of ACTIVE -> EXPIRED based on backend time
	if promo.Status == StatusActive && time.Now().After(promo.ExpiresAt) {
		promo.Status = StatusExpired
		// Persist the lazy state change
		if err := s.repo.UpdateStatus(ctx, promo.PromotionID.String(), StatusExpired); err != nil {
			// Log error but continue with expired status
		}
	}

	return promo, nil
}

func (s *serviceImpl) HandlePaymentCompleted(ctx context.Context, sessionID string) error {
	promo, err := s.repo.FindBySession(ctx, sessionID)
	if err != nil {
		return err
	}
	if promo == nil {
		// No promotion for this session, nothing to do
		return nil
	}

	// Check deadline authoritative logic
	if time.Now().After(promo.ExpiresAt) {
		// Mark as expired if not already
		if promo.Status == StatusActive {
			s.repo.UpdateStatus(ctx, promo.PromotionID.String(), StatusExpired)
		}
		// Payment completed after deadline -> do not issue voucher
		return nil
	}

	// Payment completed before deadline, handle state transitions
	if promo.Status == StatusActive {
		// State transition
		err = s.repo.UpdateStatus(ctx, promo.PromotionID.String(), StatusCompletedPendingVoucher)
		if err != nil {
			return err
		}
		promo.Status = StatusCompletedPendingVoucher
	}

	// If pending voucher or issue retry required, attempt issuance
	if promo.Status == StatusCompletedPendingVoucher || promo.Status == StatusIssueRetryRequired {
		return s.issueVoucherLocally(ctx, promo)
	}

	return nil
}

func (s *serviceImpl) RetryVoucherIssuance(ctx context.Context, sessionID string) error {
	promo, err := s.repo.FindBySession(ctx, sessionID)
	if err != nil {
		return err
	}
	if promo == nil {
		return errors.New("promotion not found")
	}

	if promo.Status != StatusIssueRetryRequired {
		return errors.New("promotion is not in a retryable state")
	}

	return s.issueVoucherLocally(ctx, promo)
}

func (s *serviceImpl) issueVoucherLocally(ctx context.Context, promo *CheckoutPromotion) error {
	res, err := s.voucher.GenerateVoucher(ctx, promo)
	
	if err != nil {
		// Mark as issue retry required on failure
		_ = s.repo.UpdateStatus(ctx, promo.PromotionID.String(), StatusIssueRetryRequired)
		return err
	}

	vStatus := VoucherStatusIssued
	err = s.repo.SaveVoucherResult(ctx, promo.PromotionID.String(), StatusVoucherIssued, &res.VoucherID, &vStatus)
	if err != nil {
		// Update failed, but voucher was generated. This requires retry to sync state.
		return err
	}

	return nil
}
