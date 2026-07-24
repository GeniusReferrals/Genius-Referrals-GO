package geniusreferrals

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// GetRedemptionRequests lists redemption requests.
func (c *Client) GetRedemptionRequests(accountSlug string, page, limit int64) (map[string]interface{}, error) {
	q := url.Values{}
	if page > 0 {
		q.Set("page", strconv.FormatInt(page, 10))
	}
	if limit > 0 {
		q.Set("limit", strconv.FormatInt(limit, 10))
	}
	path := fmt.Sprintf("/accounts/%s/redemption-requests", url.PathEscape(accountSlug))
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

// GetRedemptionRequest fetches a single redemption request.
func (c *Client) GetRedemptionRequest(accountSlug, redemptionRequestId string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/redemption-requests/%s", url.PathEscape(accountSlug), url.PathEscape(redemptionRequestId))
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

// PostRedemptionRequest creates a redemption request.
func (c *Client) PostRedemptionRequest(accountSlug string, form RedemptionRequestForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/redemption-requests", url.PathEscape(accountSlug))
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

// PatchRedemptionRequest updates a redemption request.
func (c *Client) PatchRedemptionRequest(accountSlug, redemptionRequestId string, form RedemptionRequestForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/redemption-requests/%s", url.PathEscape(accountSlug), url.PathEscape(redemptionRequestId))
	resp, body, err := c.request("PATCH", path, nil, form)
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

// RedeemRedemptionRequest processes redemption.
func (c *Client) RedeemRedemptionRequest(accountSlug, redemptionRequestId string, form map[string]interface{}) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/redemption-requests/%s/redemption", url.PathEscape(accountSlug), url.PathEscape(redemptionRequestId))
	resp, body, err := c.request("PATCH", path, nil, form)
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
