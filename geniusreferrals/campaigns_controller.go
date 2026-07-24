package geniusreferrals

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// GetCampaigns lists campaigns.
func (c *Client) GetCampaigns(accountSlug string, page, limit int64, format, fields, sort string) (map[string]interface{}, error) {
	q := url.Values{}
	if page > 0 {
		q.Set("page", strconv.FormatInt(page, 10))
	}
	if limit > 0 {
		q.Set("limit", strconv.FormatInt(limit, 10))
	}
	setOpt(&q, "format", format)
	setOpt(&q, "fields", fields)
	setOpt(&q, "sort", sort)
	path := fmt.Sprintf("/accounts/%s/campaigns", url.PathEscape(accountSlug))
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

// GetCampaign fetches a single campaign.
func (c *Client) GetCampaign(accountSlug, campaignSlug string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/campaigns/%s", url.PathEscape(accountSlug), url.PathEscape(campaignSlug))
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
