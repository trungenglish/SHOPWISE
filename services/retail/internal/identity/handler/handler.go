package handler

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"shopwise/retail/internal/identity/usecase"
	"shopwise/retail/internal/platform/apperror"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *usecase.Service
}

func NewHandler(svc *usecase.Service) *Handler {
	return &Handler{svc: svc}
}

type statusResponse struct {
	Module string `json:"module" example:"identity"`
	Status string `json:"status" example:"ok"`
}

// Status godoc
//
//	@Summary	Identity module status
//	@Tags		identity
//	@Produce	json
//	@Success	200	{object}	statusResponse
//	@Router		/identity/status [get]
func (h *Handler) Status(c *gin.Context) {
	c.JSON(http.StatusOK, statusResponse{Module: "identity", Status: "ok"})
}

// Register godoc
//
//	@Summary	Register with email and password
//	@Tags		identity
//	@Accept		json
//	@Produce	json
//	@Param		body	body	RegisterRequest	true	"Register"
//	@Success	201
//	@Failure	400	{object}	response.ProblemDetail
//	@Failure	409	{object}	response.ProblemDetail
//	@Router		/identity/register [post]
func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(apperror.Validation("invalid request body", err))
		return
	}
	if err := h.svc.Register(c.Request.Context(), usecase.RegisterInput{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	}); err != nil {
		_ = c.Error(err)
		return
	}
	c.Status(http.StatusCreated)
}

// Login godoc
//
//	@Summary	Login with email and password
//	@Tags		identity
//	@Accept		json
//	@Produce	json
//	@Param		body	body		LoginRequest	true	"Login"
//	@Success	200		{object}	TokenResponse
//	@Failure	401		{object}	response.ProblemDetail
//	@Router		/identity/login [post]
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(apperror.Validation("invalid request body", err))
		return
	}
	tokens, err := h.svc.Login(c.Request.Context(), usecase.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, toTokenResponse(tokens))
}

// Refresh godoc
//
//	@Summary	Refresh access token
//	@Tags		identity
//	@Accept		json
//	@Produce	json
//	@Param		body	body		RefreshRequest	true	"Refresh"
//	@Success	200		{object}	TokenResponse
//	@Failure	401		{object}	response.ProblemDetail
//	@Router		/identity/refresh [post]
func (h *Handler) Refresh(c *gin.Context) {
	var req RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(apperror.Validation("invalid request body", err))
		return
	}
	tokens, err := h.svc.Refresh(c.Request.Context(), usecase.RefreshInput{
		RefreshToken: req.RefreshToken,
	})
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, toTokenResponse(tokens))
}

func toTokenResponse(tokens *usecase.TokenPair) TokenResponse {
	return TokenResponse{
		AccessToken:  tokens.AccessToken,
		ExpiresIn:    tokens.ExpiresIn,
		RefreshToken: tokens.RefreshToken,
	}
}

func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {
	rg.GET("/status", h.Status)
	rg.POST("/register", h.Register)
	rg.POST("/login", h.Login)
	rg.POST("/refresh", h.Refresh)
	rg.POST("/verify-email", h.VerifyEmail)
}

func newOAuthState() (string, error) {
	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(buf), nil
}
