package formats

import (
	"encoding/csv"
	"os"

	"file-parser/internal/models"
)

// ParseCSV parses a CSV file
func ParseCSV(filePath string) (*models.ParseResult, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Allow variable fields

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return &models.ParseResult{}, nil
	}

	headers := records[0]
	var csvRecords []models.CSVRecord

	for _, row := range records[1:] {
		record := make(models.CSVRecord)
		for i, value := range row {
			if i < len(headers) {
				record[headers[i]] = value
			}
		}
		csvRecords = append(csvRecords, record)
	}

	return &models.ParseResult{
		TotalRecords: len(csvRecords),
		Headers:      headers,
		Records:      csvRecords,
	}, nil
}