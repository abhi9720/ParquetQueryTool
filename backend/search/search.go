package search

import (
	"apica-search/models"
	"apica-search/utils"
	"strings"
	"time"
)

type SearchResult struct {
	Results      []models.Record `json:"results"`
	Count        int             `json:"count"`
	SearchTimeMs int64           `json:"search_time_ms"`
}

// Simple keyword-based search (case-insensitive) over selected fields
func Search(keyword string) SearchResult {
	start := time.Now()

	records := utils.GetRecords()
	keyword = strings.ToLower(keyword)

	var matches []models.Record

	for _, record := range records {
		if containsKeyword(record, keyword) {
			matches = append(matches, record)
		}
	}

	elapsed := time.Since(start).Milliseconds()

	return SearchResult{
		Results:      matches,
		Count:        len(matches),
		SearchTimeMs: elapsed,
	}
}

// Match keyword in selected fields
func containsKeyword(r models.Record, keyword string) bool {
	return strings.Contains(strings.ToLower(r.Message), keyword) ||
		strings.Contains(strings.ToLower(r.Tag), keyword) ||
		strings.Contains(strings.ToLower(r.Sender), keyword) ||
		strings.Contains(strings.ToLower(r.StructuredData), keyword) ||
		strings.Contains(strings.ToLower(r.Groupings), keyword) ||
		strings.Contains(strings.ToLower(r.Namespace), keyword)
}
