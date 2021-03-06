package responses

import "github.com/etcinit/nexus-client-go/entities"

// FetchResponse represents a response from GET /v1/fetch
type FetchResponse struct {
	Application entities.Application `json:"application"`
	Values      map[string]string    `json:"files"`
	Status      string               `json:"status"`
}
