package zalo

import (
	"context"
	"fmt"
	"time"

	"shopwise/retail/internal/resume_session/usecase"
)

type Provider struct {
	oaID     string
	apiToken string
}

func NewProvider(oaID, apiToken string) *Provider {
	return &Provider{
		oaID:     oaID,
		apiToken: apiToken,
	}
}

func (p *Provider) Send(ctx context.Context, destination string, token string) (usecase.NotificationStatus, *string, error) {
	// Mock implementation for MVP
	// In reality, this would make an HTTP request to the Zalo OA API
	fmt.Printf("[Zalo Mock] Sending token %s to %s via OA %s\n", token, destination, p.oaID)
	
	// Simulate network latency
	time.Sleep(100 * time.Millisecond)

	// In a real implementation, we would inspect the HTTP response to determine if
	// the failure is transient (e.g. 503) or permanent (e.g. 400 Invalid phone number)
	return usecase.StatusAccepted, nil, nil
}
