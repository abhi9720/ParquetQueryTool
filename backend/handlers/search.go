package handlers

import (
	"apica-search/search"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SearchHandler handles the search requests.
func SearchHandler(c *gin.Context) {
	// Get the search query from the URL
	query := c.DefaultQuery("query", "")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing query parameter"})
		return
	}

	// Perform the search using the Search function from the search package
	result := search.Search(query)

	// Return the search results
	c.JSON(http.StatusOK, result)
}
