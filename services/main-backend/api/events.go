package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type EventsAPI struct {
	// dependencies will be injected here
}

func NewEventsAPI() *EventsAPI {
	return &EventsAPI{}
}

// POST /api/events/continue-checkout
func (api *EventsAPI) ContinueCheckout(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// POST /api/checkout-readiness/customer-info
func (api *EventsAPI) CustomerInfo(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// POST /api/events/store-selected
func (api *EventsAPI) StoreSelected(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// POST /api/events/promotion-expired
func (api *EventsAPI) PromotionExpired(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}
