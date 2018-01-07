package airly

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

const defaultBaseURL = "https://airapi.airly.eu/"

// Client is an Airly API client.
type Client struct {
	client *http.Client

	BaseURL *url.URL
	APIKey  string
}

// ClientOption represents a client option.
type ClientOption func(*Client)

// WithAPIKey sets client's API key.
func WithAPIKey(apikey string) ClientOption {
	return func(c *Client) {
		c.APIKey = apikey
	}
}

// NewClient creates new client instance with the given options applied.
func NewClient(opts ...ClientOption) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{
		client:  http.DefaultClient,
		BaseURL: baseURL,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// NewRequest creates new GET request with the given url.
func (c *Client) NewRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("apikey", c.APIKey)
	return req, nil
}

// Do sends an API request and decodes the response.
func (c *Client) Do(req *http.Request, v interface{}) error {
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	return json.NewDecoder(resp.Body).Decode(v)
}
