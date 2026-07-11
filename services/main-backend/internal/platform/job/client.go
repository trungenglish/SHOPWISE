package job

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

const (
	TypeSendWelcomeEmail = "notification:send_welcome_email"
	TypePriceCheck       = "notification:price_check"
)

type WelcomeEmailPayload struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}

type Client struct {
	client *asynq.Client
}

func NewClient(redisURL string) (*Client, error) {
	opts, err := asynq.ParseRedisURI(redisURL)
	if err != nil {
		return nil, fmt.Errorf("parse redis url: %w", err)
	}
	return &Client{client: asynq.NewClient(opts)}, nil
}

func (c *Client) Close() error {
	return c.client.Close()
}

func (c *Client) EnqueueWelcomeEmail(ctx context.Context, userID, email string) error {
	payload, err := json.Marshal(WelcomeEmailPayload{UserID: userID, Email: email})
	if err != nil {
		return fmt.Errorf("marshal welcome email payload: %w", err)
	}

	task := asynq.NewTask(TypeSendWelcomeEmail, payload)
	if _, err := c.client.EnqueueContext(ctx, task); err != nil {
		return fmt.Errorf("enqueue welcome email: %w", err)
	}
	return nil
}

func (c *Client) EnqueuePriceCheck(ctx context.Context) error {
	task := asynq.NewTask(TypePriceCheck, nil)
	if _, err := c.client.EnqueueContext(ctx, task); err != nil {
		return fmt.Errorf("enqueue price check: %w", err)
	}
	return nil
}
