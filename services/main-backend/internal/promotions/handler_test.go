package promotions

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type mockService struct {
	startPromo func(ctx context.Context, sessionID string, userID *uuid.UUID, deviceID *string, triggerType TriggerType) (*CheckoutPromotion, error)
	getStatus  func(ctx context.Context, sessionID string) (*CheckoutPromotion, error)
	handlePay  func(ctx context.Context, sessionID string) error
	retryIssue func(ctx context.Context, sessionID string) error
}

func (m *mockService) StartPromotion(ctx context.Context, sessionID string, userID *uuid.UUID, deviceID *string, triggerType TriggerType) (*CheckoutPromotion, error) {
	if m.startPromo != nil {
		return m.startPromo(ctx, sessionID, userID, deviceID, triggerType)
	}
	return nil, nil
}
func (m *mockService) GetStatus(ctx context.Context, sessionID string) (*CheckoutPromotion, error) {
	if m.getStatus != nil {
		return m.getStatus(ctx, sessionID)
	}
	return nil, nil
}
func (m *mockService) HandlePaymentCompleted(ctx context.Context, sessionID string) error {
	if m.handlePay != nil {
		return m.handlePay(ctx, sessionID)
	}
	return nil
}
func (m *mockService) RetryVoucherIssuance(ctx context.Context, sessionID string) error {
	if m.retryIssue != nil {
		return m.retryIssue(ctx, sessionID)
	}
	return nil
}

func TestStartPromotion_Handler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	svc := &mockService{
		startPromo: func(ctx context.Context, sessionID string, userID *uuid.UUID, deviceID *string, triggerType TriggerType) (*CheckoutPromotion, error) {
			now := time.Now()
			return &CheckoutPromotion{
				PromotionID: uuid.New(),
				Status:      StatusActive,
				ExpiresAt:   now.Add(15 * time.Minute),
				RewardType:  "percentage",
				RewardValue: 15,
			}, nil
		},
	}

	h := NewHandler(svc)
	router := gin.New()
	h.RegisterRoutes(&router.RouterGroup)

	body := []byte(`{"trigger_type": "saved_product"}`)
	req, _ := http.NewRequest(http.MethodPost, "/checkout-incentive/start", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Session-ID", "sess-123")
	req.Header.Set("X-User-ID", uuid.New().String())

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var res map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &res)

	if res["status"] != string(StatusActive) {
		t.Errorf("expected status ACTIVE, got %v", res["status"])
	}
	if res["reward_value"].(float64) != 15 {
		t.Errorf("expected reward_value 15, got %v", res["reward_value"])
	}
}

func TestGetStatus_Handler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	svc := &mockService{
		getStatus: func(ctx context.Context, sessionID string) (*CheckoutPromotion, error) {
			now := time.Now()
			return &CheckoutPromotion{
				PromotionID:  uuid.New(),
				Status:       StatusActive,
				ExpiresAt:    now.Add(5 * time.Minute),
				CampaignCode: "ACCESSORY_NEXT_PURCHASE_15",
				RewardType:   "percentage",
				RewardValue:  15,
				RewardScope:  "accessories_next_purchase",
			}, nil
		},
	}

	h := NewHandler(svc)
	router := gin.New()
	h.RegisterRoutes(&router.RouterGroup)

	req, _ := http.NewRequest(http.MethodGet, "/checkout-incentive/status", nil)
	req.Header.Set("X-Session-ID", "sess-123")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var res map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &res)

	if res["status"] != string(StatusActive) {
		t.Errorf("expected status ACTIVE, got %v", res["status"])
	}
	if res["campaign_code"] != "ACCESSORY_NEXT_PURCHASE_15" {
		t.Errorf("expected campaign code, got %v", res["campaign_code"])
	}
}
