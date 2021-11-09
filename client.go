package paperspace

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/Paperspace/go-graphql-client"
)

type RequestParams struct {
	Context context.Context   `json:"-"`
	Headers map[string]string `json:"-"`
}

type Client struct {
	APIKey  string
	Backend Backend
	graphql *graphql.Client
}

// client that makes requests to Gradient API
func NewClient() *Client {
	apiKey := os.Getenv("PAPERSPACE_APIKEY")
	client := Client{
		Backend: NewAPIBackend(),
		graphql: graphql.NewClientWithHeaders(os.Getenv("PAPERSPACE_BASEURL")+"/graphql", http.DefaultClient, func(h http.Header) error {
			h.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))
			return nil
		}),
	}

	if apiKey != "" {
		client.APIKey = apiKey
	}

	return &client
}

func NewClientWithBackend(backend Backend) *Client {
	client := NewClient()
	client.Backend = backend

	return client
}

func (c *Client) Request(method string, url string, params, result interface{}, requestParams RequestParams) (*http.Response, error) {
	if requestParams.Headers == nil {
		requestParams.Headers = make(map[string]string)
	}
	requestParams.Headers["x-api-key"] = c.APIKey

	return c.Backend.Request(method, url, params, result, requestParams)
}

func (c *Client) Debug() {
	c.graphql.Debug = true
}
