package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"shopwise/apps/server/internal/platform/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// const googleProvider = "google"

type GoogleOAuthService struct {
	config *oauth2.Config
}

func NewGoogleOAuthService(cfg *config.Config) *GoogleOAuthService {
	if !cfg.GoogleOAuthEnabled() {
		return nil
	}
	return &GoogleOAuthService{
		config: &oauth2.Config{
			ClientID:     cfg.GoogleClientID,
			ClientSecret: cfg.GoogleSecret,
			RedirectURL:  cfg.GoogleRedirect,
			Scopes:       []string{"openid", "email", "profile"},
			Endpoint:     google.Endpoint,
		},
	}
}

func (g *GoogleOAuthService) AuthCodeURL(state string) string {
	return g.config.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func (g *GoogleOAuthService) Exchange(ctx context.Context, code string) (*GoogleUserInfo, error) {
	token, err := g.config.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("exchange code: %w", err)
	}
	client := g.config.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return nil, fmt.Errorf("fetch userinfo: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("userinfo status %d: %s", resp.StatusCode, string(body))
	}
	var payload struct {
		Sub           string `json:"sub"`
		Email         string `json:"email"`
		Name          string `json:"name"`
		EmailVerified bool   `json:"email_verified"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, fmt.Errorf("decode userinfo: %w", err)
	}
	return &GoogleUserInfo{
		Subject:       payload.Sub,
		Email:         payload.Email,
		Name:          payload.Name,
		EmailVerified: payload.EmailVerified,
	}, nil
}
