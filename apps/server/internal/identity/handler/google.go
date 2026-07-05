package handler

import (
	"fmt"
	"net/http"
	"net/url"

	"shopwise/apps/server/internal/platform/apperror"

	"github.com/gin-gonic/gin"
)

const oauthStateCookie = "shopwise_oauth_state"

// GoogleStart godoc
//
//	@Summary	Start Google OAuth flow
//	@Tags		identity
//	@Success	302
//	@Router		/identity/google/start [get]
func (h *Handler) GoogleStart(c *gin.Context) {
	state, err := newOAuthState()
	if err != nil {
		_ = c.Error(apperror.Internal("failed to start google oauth", err))
		return
	}
	authURL, err := h.svc.GoogleAuthURL(state)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.SetCookie(oauthStateCookie, state, 600, "/", "", false, true)
	c.Redirect(http.StatusFound, authURL)
}

// GoogleCallback godoc
//
//	@Summary	Google OAuth callback
//	@Tags		identity
//	@Success	302
//	@Router		/identity/google/callback [get]
func (h *Handler) GoogleCallback(c *gin.Context, webAppURL string) {
	stateCookie, err := c.Cookie(oauthStateCookie)
	if err != nil || stateCookie == "" || stateCookie != c.Query("state") {
		_ = c.Error(apperror.Unauthorized("invalid oauth state", nil))
		return
	}
	c.SetCookie(oauthStateCookie, "", -1, "/", "", false, true)

	code := c.Query("code")
	if code == "" {
		_ = c.Error(apperror.Validation("missing authorization code", nil))
		return
	}

	tokens, err := h.svc.GoogleCallback(c.Request.Context(), code)
	if err != nil {
		_ = c.Error(err)
		return
	}

	redirectURL := fmt.Sprintf(
		"%s/auth/callback?accessToken=%s&expiresIn=%d&refreshToken=%s",
		webAppURL,
		url.QueryEscape(tokens.AccessToken),
		tokens.ExpiresIn,
		url.QueryEscape(tokens.RefreshToken),
	)
	c.Redirect(http.StatusFound, redirectURL)
}

func RegisterGoogleRoutes(rg *gin.RouterGroup, h *Handler, webAppURL string) {
	rg.GET("/google/start", h.GoogleStart)
	rg.GET("/google/callback", func(c *gin.Context) {
		h.GoogleCallback(c, webAppURL)
	})
}
