package requests

// LogsRequest represents a request for sending log lines to Nexus
type LogsRequest struct {
	Name    string   `json:"instanceName"`
	LogName string   `json:"filename"`
	Lines   []string `json:"lines"`
}
