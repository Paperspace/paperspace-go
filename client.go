package paperspace

import (
	"context"
	"net/http"
	"os"
)

type RequestParams struct {
	Context context.Context   `json:"-"`
	Headers map[string]string `json:"-"`
}

type Client struct {
	APIKey  string
	Backend Backend
}

func NewClientWithOptions(apiKey string, opts APIBackendOptions) *Client {
	client := Client{
		Backend: NewAPIBackendWithOptions(opts),
	}
	client.APIKey = apiKey
	return &client
}

// client that makes requests to Gradient API
func NewClient() *Client {
	return NewClientWithOptions(os.Getenv("PAPERSPACE_APIKEY"), APIBackendOptions{})
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
