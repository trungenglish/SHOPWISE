package api

import (
	"net/http"
	"shopwise/retail/retail"

	"github.com/gin-gonic/gin"
)

type CatalogAPI struct {
	provider *retail.MockRetailProvider
}

func NewCatalogAPI(provider *retail.MockRetailProvider) *CatalogAPI {
	return &CatalogAPI{provider: provider}
}

func (api *CatalogAPI) Search(c *gin.Context) {
	query := c.Query("q")
	results := api.provider.Search(query)
	c.JSON(http.StatusOK, gin.H{"success": true, "data": results})
}

