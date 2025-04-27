package main

import (
	"apica-search/handlers"
	"apica-search/parser"
	"apica-search/utils"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.InitStore()
	LoadAllParquetFiles()

	r := gin.Default()

	r.GET("/search", handlers.SearchHandler)
	r.POST("/upload", handlers.UploadAndStore)

	if err := r.Run(":8181"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

const uploadDir = "./parquet_files"

func LoadAllParquetFiles() {

	err := filepath.Walk(uploadDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println("Error accessing path:", err)
			return nil
		}

		if !info.IsDir() {
			log.Printf("Loading file: %s\n", path)

			records, err := parser.LoadParquetFile(path)
			if err != nil {
				log.Printf("Failed to load Parquet file: %s, Error: %v\n", path, err)
				return nil
			}

			utils.SetRecords(records)
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error scanning directory: %v", err)
	}
}
