package pkg

import (
	"crypto/tls"
	"fmt"
	"github.com/go-resty/resty/v2"
)

// APIClient contains the API Client and the module, controller, command, and host for API calls
type APIClient struct {
	client     *resty.Client
	Module     string
	Controller string
	Command    string
	Host       string
}

// NewAPIClient creates a new Client with the given address and basic auth set from key and secret
func NewAPIClient(address, key, secret string) *APIClient {
	newClient := resty.New()
	newClient.BaseURL = fmt.Sprintf("%s%s", address, "/api")
	newClient.SetBasicAuth(key, secret)
	return &APIClient{
		client: newClient,
	}
}

// SetTLSVerify sets the TLS verification for the API client
func (c *APIClient) SetTLSVerify(verify bool) {
	c.client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: !verify})
}
