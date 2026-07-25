package geniusreferrals

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// GetTags lists tags.
func (c *Client) GetTags(accountSlug string, page, limit int64) (map[string]interface{}, error) {
	q := url.Values{}
	if page > 0 {
		q.Set("page", strconv.FormatInt(page, 10))
	}
	if limit > 0 {
		q.Set("limit", strconv.FormatInt(limit, 10))
	}
	path := fmt.Sprintf("/accounts/%s/tags", url.PathEscape(accountSlug))
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

// GetTag fetches a single tag.
func (c *Client) GetTag(accountSlug, tagSlug string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/tags/%s", url.PathEscape(accountSlug), url.PathEscape(tagSlug))
	resp, body, err := c.request("GET", path, nil, nil)
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

// CreateTag creates a tag.
func (c *Client) CreateTag(accountSlug string, form TagForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/tags", url.PathEscape(accountSlug))
	resp, body, err := c.request("POST", path, nil, form)
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

// UpdateTag updates a tag.
func (c *Client) UpdateTag(accountSlug, tagSlug string, form TagForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/tags/%s", url.PathEscape(accountSlug), url.PathEscape(tagSlug))
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

// DeleteTag deletes a tag.
func (c *Client) DeleteTag(accountSlug, tagSlug string) error {
	path := fmt.Sprintf("/accounts/%s/tags/%s", url.PathEscape(accountSlug), url.PathEscape(tagSlug))
	resp, body, err := c.request("DELETE", path, nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 206 {
		return newAPIError(resp.StatusCode, body)
	}
	return nil
}
