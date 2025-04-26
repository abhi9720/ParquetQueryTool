package main

import (
	"apica-search/handlers"
	"apica-search/parser"
	"apica-search/utils"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	records, err := parser.LoadParquetFile("parquet_files\\File 1")
	if err != nil {
		log.Fatal("Failed to load parquet:", err)
	}
	utils.SetRecords(records)

	fmt.Println(utils.GetRecords())

	// Initialize Gin router
	r := gin.Default()

	// Register the /search route
	r.GET("/search", handlers.SearchHandler)

	// Start the server
	if err := r.Run(":8181"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
