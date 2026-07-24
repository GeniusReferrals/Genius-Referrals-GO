package geniusreferrals

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// GetReferrals lists referrals for an advocate.
func (c *Client) GetReferrals(accountSlug, advocateToken string, page, limit int64) (map[string]interface{}, error) {
	q := url.Values{}
	if page > 0 {
		q.Set("page", strconv.FormatInt(page, 10))
	}
	if limit > 0 {
		q.Set("limit", strconv.FormatInt(limit, 10))
	}
	path := fmt.Sprintf("/accounts/%s/advocates/%s/referrals", url.PathEscape(accountSlug), url.PathEscape(advocateToken))
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

// GetReferral fetches a single referral.
func (c *Client) GetReferral(accountSlug, advocateToken, referralId string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/advocates/%s/referrals/%s", url.PathEscape(accountSlug), url.PathEscape(advocateToken), url.PathEscape(referralId))
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

// PostReferral creates a referral.
func (c *Client) PostReferral(accountSlug, advocateToken string, form ReferralForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/advocates/%s/referrals", url.PathEscape(accountSlug), url.PathEscape(advocateToken))
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

// PutReferral updates a referral.
func (c *Client) PutReferral(accountSlug, advocateToken, referralId string, form ReferralForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/advocates/%s/referrals/%s", url.PathEscape(accountSlug), url.PathEscape(advocateToken), url.PathEscape(referralId))
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

// DeleteReferral deletes a referral.
func (c *Client) DeleteReferral(accountSlug, advocateToken, referralId string) error {
	path := fmt.Sprintf("/accounts/%s/advocates/%s/referrals/%s", url.PathEscape(accountSlug), url.PathEscape(advocateToken), url.PathEscape(referralId))
	resp, body, err := c.request("DELETE", path, nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 206 {
		return newAPIError(resp.StatusCode, body)
	}
	return nil
}
