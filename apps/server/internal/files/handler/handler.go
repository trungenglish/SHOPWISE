package handler

import (
	"net/http"

	"shopwise/apps/server/internal/files/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *usecase.Service
}

func NewHandler(svc *usecase.Service) *Handler {
	return &Handler{svc: svc}
}

type statusResponse struct {
	Module string `json:"module" example:"files"`
	Status string `json:"status" example:"ok"`
}

// Status godoc
//
//	@Summary	Files module status
//	@Tags		files
//	@Produce	json
//	@Success	200	{object}	statusResponse
//	@Router		/files/status [get]
func (h *Handler) Status(c *gin.Context) {
	c.JSON(http.StatusOK, statusResponse{Module: "files", Status: "ok"})
}

func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {
	rg.GET("/status", h.Status)
}
