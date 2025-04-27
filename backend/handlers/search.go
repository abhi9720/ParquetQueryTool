package handlers

import (
	"apica-search/search"
	"apica-search/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchHandler(c *gin.Context) {

	query := c.DefaultQuery("query", "")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing query parameter"})
		return
	}

	result, err := search.Search(query)

	if err != nil {
		if err == utils.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"results":        []interface{}{},
				"count":          0,
				"search_time_ms": result.SearchTimeMs,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "search failed"})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"results":        result.Results,
		"count":          result.Count,
		"search_time_ms": result.SearchTimeMs,
	})
}
