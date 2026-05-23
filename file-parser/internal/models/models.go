package models

// CSVRecord represents a row in CSV
type CSVRecord map[string]string

// LogEntry represents a parsed log line
type LogEntry struct {
	IP        string `json:"ip"`
	Timestamp string `json:"timestamp"`
	Method    string `json:"method"`
	Path      string `json:"path"`
	Status    int    `json:"status"`
	Bytes     int    `json:"bytes"`
	UserAgent string `json:"user_agent"`
}

// ParseResult holds parsing summary
type ParseResult struct {
	TotalRecords int         `json:"total_records"`
	Headers      []string    `json:"headers,omitempty"`
	Records      []CSVRecord `json:"records,omitempty"`
	LogEntries   []LogEntry  `json:"log_entries,omitempty"`
}