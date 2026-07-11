package handler

import (
	"net/http"

	"shopwise/retail/internal/administration/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *usecase.Service
}

func NewHandler(svc *usecase.Service) *Handler {
	return &Handler{svc: svc}
}

type statusResponse struct {
	Module string `json:"module" example:"administration"`
	Status string `json:"status" example:"ok"`
}

// Status godoc
//
//	@Summary	Administration module status
//	@Tags		administration
//	@Produce	json
//	@Success	200	{object}	statusResponse
//	@Router		/admin/status [get]
func (h *Handler) Status(c *gin.Context) {
	c.JSON(http.StatusOK, statusResponse{Module: "administration", Status: "ok"})
}

func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {
	rg.GET("/status", h.Status)
}
