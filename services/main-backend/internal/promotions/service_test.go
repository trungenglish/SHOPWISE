package promotions

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
)

type mockRepo struct {
	promotions map[string]*CheckoutPromotion
}

func newMockRepo() *mockRepo {
	return &mockRepo{
		promotions: make(map[string]*CheckoutPromotion),
	}
}

func (m *mockRepo) Create(ctx context.Context, promotion *CheckoutPromotion) error {
	m.promotions[promotion.SessionID] = promotion
	return nil
}

func (m *mockRepo) FindBySession(ctx context.Context, sessionID string) (*CheckoutPromotion, error) {
	if p, ok := m.promotions[sessionID]; ok {
		return p, nil
	}
	return nil, nil
}

func (m *mockRepo) FindByEligibilityKey(ctx context.Context, userID *string, deviceID *string) (*CheckoutPromotion, error) {
	var latest *CheckoutPromotion
	for _, p := range m.promotions {
		if userID != nil && p.UserID != nil && *userID == p.UserID.String() {
			if latest == nil || p.CreatedAt.After(latest.CreatedAt) {
				latest = p
			}
		} else if deviceID != nil && p.DeviceID != nil && *deviceID == *p.DeviceID {
			if latest == nil || p.CreatedAt.After(latest.CreatedAt) {
				latest = p
			}
		}
	}
	return latest, nil
}

func (m *mockRepo) UpdateStatus(ctx context.Context, promotionID string, status PromotionStatus) error {
	for _, p := range m.promotions {
		if p.PromotionID.String() == promotionID {
			p.Status = status
			return nil
		}
	}
	return errors.New("not found")
}

func (m *mockRepo) SaveVoucherResult(ctx context.Context, promotionID string, status PromotionStatus, voucherID *string, voucherStatus *VoucherStatus) error {
	for _, p := range m.promotions {
		if p.PromotionID.String() == promotionID {
			p.Status = status
			if voucherID != nil {
				uid := uuid.MustParse(*voucherID)
				p.VoucherID = &uid
			}
			if voucherStatus != nil {
				p.VoucherStatus = voucherStatus
			}
			return nil
		}
	}
	return errors.New("not found")
}

type mockVoucher struct{}

func (m *mockVoucher) GenerateVoucher(ctx context.Context, promotion *CheckoutPromotion) (*VoucherResult, error) {
	return &VoucherResult{
		VoucherID: uuid.New().String(),
		Code:      "MOCK-CODE",
	}, nil
}

func TestStartPromotion_Idempotent(t *testing.T) {
	repo := newMockRepo()
	svc := NewService(repo, &mockVoucher{})

	sessionID := "session-1"
	userID := uuid.New()

	p1, err := svc.StartPromotion(context.Background(), sessionID, &userID, nil, TriggerSavedProduct)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	p2, err := svc.StartPromotion(context.Background(), sessionID, &userID, nil, TriggerSavedProduct)
	if err != nil {
		t.Fatalf("expected no error on duplicate start, got %v", err)
	}

	if p1.PromotionID != p2.PromotionID {
		t.Errorf("expected idempotent promotion, got different IDs: %s != %s", p1.PromotionID, p2.PromotionID)
	}
}

func TestStartPromotion_Cooldown(t *testing.T) {
	repo := newMockRepo()
	svc := NewService(repo, &mockVoucher{})

	userID := uuid.New()

	// Simulate an existing promotion from a previous session that has a cooldown in the future
	now := time.Now()
	cooldown := now.Add(24 * time.Hour)
	repo.promotions["session-old"] = &CheckoutPromotion{
		SessionID:     "session-old",
		UserID:        &userID,
		Status:        StatusActive,
		CreatedAt:     now.Add(-1 * time.Hour),
		CooldownUntil: &cooldown,
	}

	_, err := svc.StartPromotion(context.Background(), "session-new", &userID, nil, TriggerSavedProduct)
	if err == nil || err.Error() != "user is in cooldown period and not eligible for a new promotion" {
		t.Fatalf("expected cooldown error, got %v", err)
	}
}

func TestHandlePaymentCompleted_BeforeDeadline(t *testing.T) {
	repo := newMockRepo()
	svc := NewService(repo, &mockVoucher{})

	sessionID := "session-1"
	userID := uuid.New()

	now := time.Now()
	promo := &CheckoutPromotion{
		PromotionID: uuid.New(),
		SessionID:   sessionID,
		UserID:      &userID,
		Status:      StatusActive,
		ExpiresAt:   now.Add(5 * time.Minute), // deadline in future
	}
	repo.promotions[sessionID] = promo

	err := svc.HandlePaymentCompleted(context.Background(), sessionID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if promo.Status != StatusVoucherIssued {
		t.Errorf("expected status to be VOUCHER_ISSUED, got %s", promo.Status)
	}
	if promo.VoucherID == nil {
		t.Errorf("expected voucher to be issued")
	}
}

func TestHandlePaymentCompleted_AfterDeadline(t *testing.T) {
	repo := newMockRepo()
	svc := NewService(repo, &mockVoucher{})

	sessionID := "session-1"
	userID := uuid.New()

	now := time.Now()
	promo := &CheckoutPromotion{
		PromotionID: uuid.New(),
		SessionID:   sessionID,
		UserID:      &userID,
		Status:      StatusActive,
		ExpiresAt:   now.Add(-5 * time.Minute), // deadline in past
	}
	repo.promotions[sessionID] = promo

	err := svc.HandlePaymentCompleted(context.Background(), sessionID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if promo.Status != StatusExpired {
		t.Errorf("expected status to be EXPIRED, got %s", promo.Status)
	}
	if promo.VoucherID != nil {
		t.Errorf("expected no voucher to be issued")
	}
}
