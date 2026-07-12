package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"shopwise/retail/internal/resume_session/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockFailProvider struct {
	calls  int
	status usecase.NotificationStatus
}

func (m *mockFailProvider) Send(ctx context.Context, destination string, token string) (usecase.NotificationStatus, *string, error) {
	m.calls++
	if m.status == usecase.StatusRetryableFailure {
		errStr := "transient error"
		return m.status, &errStr, errors.New("timeout")
	}
	if m.status == usecase.StatusFailed {
		errStr := "permanent error"
		return m.status, &errStr, errors.New("bad request")
	}
	return usecase.StatusAccepted, nil, nil
}

func TestRetryableProvider_TransientFailure(t *testing.T) {
	mockProv := &mockFailProvider{status: usecase.StatusRetryableFailure}
	// Configure for 1 retry
	retryable := usecase.NewRetryableProvider(mockProv, 1)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	status, _, err := retryable.Send(ctx, "+1234567890", "token123")
	require.Error(t, err)
	assert.Equal(t, usecase.StatusRetryableFailure, status)
	// Initial call + 1 retry = 2 calls
	assert.Equal(t, 2, mockProv.calls)
}

func TestRetryableProvider_PermanentFailure(t *testing.T) {
	mockProv := &mockFailProvider{status: usecase.StatusFailed}
	retryable := usecase.NewRetryableProvider(mockProv, 1)

	ctx := context.Background()

	status, _, err := retryable.Send(ctx, "+1234567890", "token123")
	require.Error(t, err)
	assert.Equal(t, usecase.StatusFailed, status)
	// Initial call returns StatusFailed immediately, no retries = 1 call
	assert.Equal(t, 1, mockProv.calls)
}
