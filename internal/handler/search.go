package handler

import (
	"net/http"

	"github.com/abrar-hossain/supplier-risk-radar/internal/company"
	"github.com/gin-gonic/gin"
)

type SearchHandler struct {
	CompanyClient *company.CompanyClient
}

func NewSearchHandler(client *company.CompanyClient) *SearchHandler {
	return &SearchHandler{
		CompanyClient: client,
	}
}

func (h *SearchHandler) Search(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query parameter q is required"})
		return
	}

	results, err := h.CompanyClient.SearchCompanies(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": results})
}
