// Package geniusreferrals provides a Go SDK for the Genius Referrals REST API.
package geniusreferrals

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const defaultBaseURI = "https://api.geniusreferrals.com"
const defaultUserAgent = "GeniusReferrals-GoSDK/2.0"

// Client is the main SDK client.
type Client struct {
	HTTPClient  *http.Client
	BaseURI     string
	AuthToken   string
	UserAgent   string
}

// NewClient creates a new Client with the given API token.
func NewClient(apiToken string) *Client {
	return &Client{
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
		BaseURI:    defaultBaseURI,
		AuthToken:  apiToken,
		UserAgent:  defaultUserAgent,
	}
}

// NewClientWithBaseURI creates a new Client with a custom base URI.
func NewClientWithBaseURI(apiToken, baseURI string) *Client {
	c := NewClient(apiToken)
	c.BaseURI = baseURI
	return c
}

// request performs an HTTP request and returns the parsed response.
func (c *Client) request(method, path string, query url.Values, body interface{}) (*http.Response, []byte, error) {
	u, err := url.Parse(c.BaseURI + path)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid url: %w", err)
	}
	if query != nil {
		u.RawQuery = query.Encode()
	}

	var bodyReader io.Reader
	if body != nil {
		if b, ok := body.([]byte); ok {
			bodyReader = bytes.NewReader(b)
		} else {
			data, err := json.Marshal(body)
			if err != nil {
				return nil, nil, fmt.Errorf("marshal body: %w", err)
			}
			bodyReader = bytes.NewReader(data)
		}
	}

	req, err := http.NewRequest(method, u.String(), bodyReader)
	if err != nil {
		return nil, nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Token", c.AuthToken)
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("http request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("read body: %w", err)
	}

	return resp, respBody, nil
}
