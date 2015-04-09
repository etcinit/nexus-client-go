package nexus

import (
	"encoding/json"
	"errors"
	"os"
	"unicode/utf8"

	"github.com/PuerkitoBio/purell"
	"github.com/etcinit/nexus-client-go/responses"
	"github.com/parnurzeal/gorequest"
)

// Client is an implementation of the Nexus configuration server client
type Client struct {
	endpoint string
	token    string
}

// NewClient creates a new instance of a NexusClient
func NewClient(endpoint string, token string) *Client {
	client := Client{
		endpoint: endpoint,
		token:    token,
	}

	return &client
}

// NewClientFromEnv creates a new instance of a NexusClient from environment
// variables
func NewClientFromEnv() (*Client, error) {
	endpoint := os.Getenv("NEXUS_ENDPOINT")
	token := os.Getenv("NEXUS_APIKEY")

	if utf8.RuneCountInString(endpoint) < 1 || utf8.RuneCountInString(endpoint) < 1 {
		return &Client{}, errors.New("Required environment is not present or invalid")
	}

	client := Client{
		endpoint: endpoint,
		token:    token,
	}

	return &client, nil
}

// buildURL constructs an URL to make a call to the Nexus API
func (c *Client) buildURL(path string) string {
	endpoint, err := purell.NormalizeURLString(
		c.endpoint+path,
		purell.FlagLowercaseScheme|purell.FlagLowercaseScheme|purell.FlagLowercaseHost|purell.FlagRemoveDuplicateSlashes,
	)

	if err != nil {
		panic(err)
	}

	return endpoint
}

// buildAuthHeader constructs the HTTP Authroization header contents
func (c *Client) buildAuthHeader() string {
	return "Bearer " + c.token
}

// Fetch fetches all configuration variables assigned to the token
func (c *Client) Fetch() (*responses.FetchResponse, []error) {
	url := c.buildURL("/v1/fetch")

	_, body, errs := gorequest.New().
		Get(url).
		Set("Authorization", c.buildAuthHeader()).
		End()

	if errs != nil {
		return &responses.FetchResponse{}, errs
	}

	var response responses.FetchResponse
	json.Unmarshal([]byte(body), &response)

	return &response, nil
}
