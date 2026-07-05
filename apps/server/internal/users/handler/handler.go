package handler

import (
	"net/http"
	"strconv"

	"shopwise/apps/server/internal/platform/apperror"
	"shopwise/apps/server/internal/platform/middleware"
	"shopwise/apps/server/internal/users/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	svc *usecase.Service
}

func NewHandler(svc *usecase.Service) *Handler {
	return &Handler{svc: svc}
}

// List godoc
//
//	@Summary	List users
//	@Tags		users
//	@Produce	json
//	@Param		limit	query		int	false	"Page size"	default(20)
//	@Param		offset	query		int	false	"Offset"	default(0)
//	@Success	200		{object}	UserListResponse
//	@Failure	500		{object}	ErrorResponse
//	@Router		/users [get]
func (h *Handler) List(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	users, err := h.svc.List(c.Request.Context(), limit, offset)
	if err != nil {
		_ = c.Error(err)
		return
	}

	items := make([]UserResponse, 0, len(users))
	for i := range users {
		items = append(items, toUserResponse(users[i].ID, users[i].Email, users[i].Name, users[i].CreatedAt, users[i].UpdatedAt))
	}

	c.JSON(http.StatusOK, UserListResponse{
		Items:  items,
		Limit:  limit,
		Offset: offset,
	})
}

// Create godoc
//
//	@Summary	Create user
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		body	body		CreateUserRequest	true	"Create user"
//	@Success	201		{object}	UserResponse
//	@Failure	400		{object}	ErrorResponse
//	@Failure	409		{object}	ErrorResponse
//	@Failure	500		{object}	ErrorResponse
//	@Router		/users [post]
func (h *Handler) Create(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(apperror.Validation("invalid request body", err))
		return
	}

	user, err := h.svc.Create(c.Request.Context(), usecase.CreateInput{
		Email: req.Email,
		Name:  req.Name,
	})
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, toUserResponse(user.ID, user.Email, user.Name, user.CreatedAt, user.UpdatedAt))
}

// GetByID godoc
//
//	@Summary	Get user by ID
//	@Tags		users
//	@Produce	json
//	@Param		id	path		string	true	"User ID"	format(uuid)
//	@Success	200	{object}	UserResponse
//	@Failure	404	{object}	ErrorResponse
//	@Failure	500	{object}	ErrorResponse
//	@Router		/users/{id} [get]
func (h *Handler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		_ = c.Error(apperror.Validation("invalid user id", err))
		return
	}

	user, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, toUserResponse(user.ID, user.Email, user.Name, user.CreatedAt, user.UpdatedAt))
}

// Update godoc
//
//	@Summary	Update user
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		id		path		string				true	"User ID"	format(uuid)
//	@Param		body	body		UpdateUserRequest	true	"Update user"
//	@Success	200		{object}	UserResponse
//	@Failure	400		{object}	ErrorResponse
//	@Failure	404		{object}	ErrorResponse
//	@Failure	409		{object}	ErrorResponse
//	@Failure	500		{object}	ErrorResponse
//	@Router		/users/{id} [patch]
func (h *Handler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		_ = c.Error(apperror.Validation("invalid user id", err))
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(apperror.Validation("invalid request body", err))
		return
	}

	user, err := h.svc.Update(c.Request.Context(), id, usecase.UpdateInput{
		Email: req.Email,
		Name:  req.Name,
	})
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, toUserResponse(user.ID, user.Email, user.Name, user.CreatedAt, user.UpdatedAt))
}

// Delete godoc
//
//	@Summary	Delete user
//	@Tags		users
//	@Param		id	path	string	true	"User ID"	format(uuid)
//	@Success	204
//	@Failure	404	{object}	ErrorResponse
//	@Failure	500	{object}	ErrorResponse
//	@Router		/users/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		_ = c.Error(apperror.Validation("invalid user id", err))
		return
	}

	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		_ = c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}

// RegisterRoutes mounts user routes on the given group (prefix /users).
func RegisterRoutes(rg *gin.RouterGroup, h *Handler, verifier middleware.TokenVerifier) {
	authed := rg.Group("")
	authed.Use(middleware.Auth(verifier))
	authed.GET("/me", h.GetMe)
	authed.PATCH("/me", h.PatchMe)
	authed.DELETE("/me", h.DeleteMe)

	rg.GET("", h.List)
	rg.POST("", h.Create)
	rg.GET("/:id", h.GetByID)
	rg.PATCH("/:id", h.Update)
	rg.DELETE("/:id", h.Delete)
}
