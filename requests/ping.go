package requests

// PingRequest represents a POST /v1/ping requests
type PingRequest struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}
