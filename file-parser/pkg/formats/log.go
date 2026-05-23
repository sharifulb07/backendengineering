package formats

import (
	"bufio"
	"os"
	"regexp"
	"strconv"

	"file-parser/internal/models"
)

// ParseNginxLog parses nginx access log
func ParseNginxLog(filePath string) (*models.ParseResult, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var entries []models.LogEntry
	scanner := bufio.NewScanner(file)

	// Common nginx log regex
	re := regexp.MustCompile(`^(\S+) \S+ \S+ \[([^\]]+)\] "(\S+) (\S+) \S+" (\d+) (\d+)`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)

		if len(matches) > 6 {
			status, _ := strconv.Atoi(matches[5])
			bytes, _ := strconv.Atoi(matches[6])

			entry := models.LogEntry{
				IP:        matches[1],
				Timestamp: matches[2],
				Method:    matches[3],
				Path:      matches[4],
				Status:    status,
				Bytes:     bytes,
				UserAgent: "", // Can be extended
			}
			entries = append(entries, entry)
		}
	}

	return &models.ParseResult{
		TotalRecords: len(entries),
		LogEntries:   entries,
	}, nil
}