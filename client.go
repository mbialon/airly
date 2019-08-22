package airly

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

const defaultBaseURL = "https://airapi.airly.eu/"

// Client is an Airly API Client.
type Client struct {
	Client  *http.Client
	BaseURL *url.URL
	APIKey  string
}

// NewClient creates new Client instance with the given API key.
func NewClient(apikey string) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)
	return &Client{
		Client:  http.DefaultClient,
		BaseURL: baseURL,
		APIKey:  apikey,
	}
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
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	return json.NewDecoder(resp.Body).Decode(v)
}
