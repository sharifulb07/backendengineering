package utils

import (
	"file-parser/internal/models"
	"fmt"
	"os"
	"time"
)

// FileExists checks if file exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// PrintStats prints parsing statistics
func PrintStats(result *models.ParseResult) {
	fmt.Printf("\n✅ Parsing completed!\n")
	fmt.Printf("Total Records : %d\n", result.TotalRecords)
	fmt.Printf("Time          : %v\n", time.Now().Format(time.RFC3339))
}