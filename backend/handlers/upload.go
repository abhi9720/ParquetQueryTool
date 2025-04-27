package handlers

import (
	"apica-search/parser"
	"apica-search/utils"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const uploadDir = "./parquet_files"

func UploadAndStore(c *gin.Context) {

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed"})
		return
	}

	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
		return
	}

	filePath := filepath.Join(uploadDir, fileHeader.Filename)

	dst, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy file"})
		return
	}

	records, err := parser.LoadParquetFile(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing Parquet"})
		return
	}

	utils.SetRecords(records)

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("File '%s' uploaded and indexed successfully", fileHeader.Filename)})
}
