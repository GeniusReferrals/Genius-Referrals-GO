package geniusreferrals

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// GetAccounts lists accounts.
func (c *Client) GetAccounts(page, limit int64, format, clientSlug string) (map[string]interface{}, error) {
	q := url.Values{}
	if page > 0 {
		q.Set("page", strconv.FormatInt(page, 10))
	}
	if limit > 0 {
		q.Set("limit", strconv.FormatInt(limit, 10))
	}
	if format != "" {
		q.Set("format", format)
	}
	if clientSlug != "" {
		q.Set("client_slug", clientSlug)
	}
	resp, body, err := c.request("GET", "/accounts", q, nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 206 {
		return nil, newAPIError(resp.StatusCode, body)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetAccount fetches a single account.
func (c *Client) GetAccount(accountSlug, clientSlug string) (map[string]interface{}, error) {
	q := url.Values{}
	if clientSlug != "" {
		q.Set("client_slug", clientSlug)
	}
	path := fmt.Sprintf("/accounts/%s", url.PathEscape(accountSlug))
	resp, body, err := c.request("GET", path, q, nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 206 {
		return nil, newAPIError(resp.StatusCode, body)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateAccount creates a new account.
func (c *Client) CreateAccount(form AccountForm) (map[string]interface{}, error) {
	resp, body, err := c.request("POST", "/accounts", nil, form)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 206 {
		return nil, newAPIError(resp.StatusCode, body)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateAccount updates an account.
func (c *Client) UpdateAccount(accountSlug string, form AccountForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s", url.PathEscape(accountSlug))
	resp, body, err := c.request("PUT", path, nil, form)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 206 {
		return nil, newAPIError(resp.StatusCode, body)
	}
	if resp.StatusCode == 204 || len(body) == 0 {
		return nil, nil
	}
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteAccount deletes an account.
func (c *Client) DeleteAccount(accountSlug string) error {
	path := fmt.Sprintf("/accounts/%s", url.PathEscape(accountSlug))
	resp, body, err := c.request("DELETE", path, nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 206 {
		return newAPIError(resp.StatusCode, body)
	}
	return nil
}
