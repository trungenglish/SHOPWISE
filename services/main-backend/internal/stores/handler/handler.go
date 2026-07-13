package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StoreResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ListStores(c *gin.Context) {
	stores := []StoreResponse{
		{ID: "pv-1", Name: "Phong Vũ Nguyễn Thị Minh Khai", Address: "264 Nguyễn Thị Minh Khai, P. Võ Thị Sáu, Q. 3, TP.HCM"},
		{ID: "pv-2", Name: "Phong Vũ Cách Mạng Tháng 8", Address: "288 Cách Mạng Tháng 8, P. 10, Q. 3, TP.HCM"},
		{ID: "pv-3", Name: "Phong Vũ Thái Hà", Address: "1 Thái Hà, Trung Liệt, Đống Đa, Hà Nội"},
		{ID: "pv-4", Name: "Phong Vũ Cầu Giấy", Address: "173 Xuân Thủy, Dịch Vọng Hậu, Cầu Giấy, Hà Nội"},
		{ID: "pv-5", Name: "Phong Vũ Lê Văn Việt", Address: "1A Lê Văn Việt, Hiệp Phú, TP. Thủ Đức, TP.HCM"},
	}
	c.JSON(http.StatusOK, gin.H{"items": stores})
}

func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {
	rg.GET("", h.ListStores)
}
