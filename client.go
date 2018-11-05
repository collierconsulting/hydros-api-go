package client

import (
	"errors"
	"net/http"
	"net/url"
	"time"
)

const (
	// DefaultURL default URL
	DefaultURL      = "https://localhost"
	// DefaultAuthType default authentication type
	DefaultAuthType = AuthTypeOpenID
)

// NewHydrosClient Creates instance of *HydrosClient
func NewHydrosClient(options ...HydrosClientOptionFunc) (*HydrosClient, error) {
	parsedURL, _ := url.Parse(DefaultURL)
	client := &HydrosClient{
		AuthType: DefaultAuthType,
		URL:      parsedURL,
		HTTPClient: http.Client{
			Timeout: time.Second * 60,
		},
	}

	// Run option functions
	for _, option := range options {
		if err := option(client); err != nil {
			return nil, err
		}
	}

	// Create service instances
	client.Driller = NewDrillerService(client)

	return client, nil
}

// HydrosClient Hydros API client
type HydrosClient struct {
	AuthType    AuthType
	AccessToken string
	URL         *url.URL
	HTTPClient  http.Client
	Driller     DrillerService
}

// HydrosClientOptionFunc Hydros API client option
type HydrosClientOptionFunc func(*HydrosClient) error

// SetHost updates host URL API calls are made against.  Otherwise, the default URL is used
func SetHost(host string) HydrosClientOptionFunc {
	return func(c *HydrosClient) error {
		parsedURL, err := url.Parse(host)
		if err != nil {
			return err
		}
		if parsedURL.Scheme == "" {
			return errors.New("missing scheme from host parameter: please prefix with http:// or https://")
		}
		c.URL = parsedURL
		return nil
	}
}

// SetAccessToken sets a global access token to be used for client for duration of session
func SetAccessToken(accessToken string) HydrosClientOptionFunc {
	return func(c *HydrosClient) error {
		c.AuthType = AuthTypeOpenID
		c.AccessToken = accessToken
		return nil
	}
}
