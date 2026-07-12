package usecase

import (
	"context"
	"time"
)

type NotificationStatus string

const (
	StatusAccepted         NotificationStatus = "Accepted"
	StatusFailed           NotificationStatus = "Failed"
	StatusRetryableFailure NotificationStatus = "RetryableFailure"
)

type NotificationProvider interface {
	// Send sends the resume link to the destination (e.g. phone number).
	// Returns the status and any error details as a string.
	Send(ctx context.Context, destination string, token string) (NotificationStatus, *string, error)
}

// RetryableProvider wraps a NotificationProvider and adds retry logic for transient failures
type RetryableProvider struct {
	provider   NotificationProvider
	maxRetries int
}

func NewRetryableProvider(p NotificationProvider, maxRetries int) *RetryableProvider {
	return &RetryableProvider{
		provider:   p,
		maxRetries: maxRetries,
	}
}

func (r *RetryableProvider) Send(ctx context.Context, destination string, token string) (NotificationStatus, *string, error) {
	var status NotificationStatus
	var details *string
	var err error

	for i := 0; i <= r.maxRetries; i++ {
		status, details, err = r.provider.Send(ctx, destination, token)
		if status == StatusAccepted || status == StatusFailed {
			return status, details, err
		}
		
		// If StatusRetryableFailure, wait before retrying (exponential backoff)
		if i < r.maxRetries {
			select {
			case <-ctx.Done():
				return StatusFailed, nil, ctx.Err()
			case <-time.After(time.Duration(1<<i) * time.Second):
				// wait and loop
			}
		}
	}
	return status, details, err
}
