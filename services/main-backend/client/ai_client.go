package client

type AIClient struct {
	BaseURL string
}

func NewAIClient(baseURL string) *AIClient {
	return &AIClient{BaseURL: baseURL}
}

func (c *AIClient) SendMessage(sessionID, message string) error {
	// TODO: Send HTTP POST to AI Runtime
	return nil
}
