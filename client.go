package hydros

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	// DefaultURL default URL
	DefaultURL = "https://localhost"
	// DefaultAuthType default authentication type
	DefaultAuthType = AuthTypeOpenID
)

// NewClient Creates instance of *Client
func NewClient(options ...ClientOptionFunc) (*Client, error) {
	parsedURL, _ := url.Parse(DefaultURL)
	client := &Client{
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

	client.CreateHeadersFunc = func() []RequestHeader {
		headers := make([]RequestHeader, 3)
		headers[0] = RequestHeader{Key: "Authorization", Value: fmt.Sprintf("Bearer %s", client.AccessToken)}
		headers[1] = RequestHeader{Key: "Content-Type", Value: "application/json"}
		headers[2] = RequestHeader{Key: "Accept", Value: "application/json"}
		return headers
	}

	// Create service instances
	client.Driller = NewDrillerService(client)
	client.Well = NewWellService(client)

	return client, nil
}

// Client Hydros API client
type Client struct {
	AuthType          AuthType
	AccessToken       string
	CreateHeadersFunc func() []RequestHeader
	URL               *url.URL
	HTTPClient        http.Client
	Driller           DrillerService
	Well              WellService
}

// RequestHeader hold key value pairs
type RequestHeader struct {
	Key   string
	Value string
}

// ClientOptionFunc Hydros API client option
type ClientOptionFunc func(*Client) error

// SetHost updates host URL API calls are made against.  Otherwise, the default URL is used
func SetHost(host string) ClientOptionFunc {
	return func(c *Client) error {
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
func SetAccessToken(accessToken string) ClientOptionFunc {
	return func(c *Client) error {
		c.AuthType = AuthTypeOpenID
		c.AccessToken = accessToken
		return nil
	}
}
