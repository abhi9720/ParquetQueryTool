package search

import (
	"apica-search/models"
	"apica-search/utils"
	"math"
	"strings"
	"time"
)

type SearchResult struct {
	Results      []models.Record `json:"results"`
	Count        int             `json:"count"`
	SearchTimeMs int64           `json:"search_time_ms"`
}

func Search(query string) (SearchResult, error) {
	start := time.Now()

	words := strings.Fields(strings.ToLower(query))
	if len(words) == 0 {
		return SearchResult{
			Results:      []models.Record{},
			Count:        0,
			SearchTimeMs: time.Since(start).Milliseconds(),
		}, nil
	}

	recordMap := make(map[*models.Record]int)

	for _, word := range words {
		matchedRecords, err := utils.Store.Search(word)
		if err != nil {
			return SearchResult{
				Results:      []models.Record{},
				Count:        0,
				SearchTimeMs: time.Since(start).Milliseconds(),
			}, err
		}
		for _, rec := range matchedRecords {
			recordMap[rec]++
		}
	}

	var results []models.Record
	for rec, count := range recordMap {
		if count >= 1 {
			results = append(results, *rec)
		}
	}

	return SearchResult{
		Results:      results[:int(math.Min(float64(len(results)), 1000))], // showing only top 1000 records only
		Count:        len(results),
		SearchTimeMs: time.Since(start).Milliseconds(),
	}, nil
}
