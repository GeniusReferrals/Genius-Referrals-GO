package geniusreferrals

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// GetBonuses lists bonuses.
func (c *Client) GetBonuses(accountSlug string, page, limit int64, format, advocateToken, status, createdRange, updatedRange, fields, sort string) (map[string]interface{}, error) {
	q := url.Values{}
	if page > 0 {
		q.Set("page", strconv.FormatInt(page, 10))
	}
	if limit > 0 {
		q.Set("limit", strconv.FormatInt(limit, 10))
	}
	setOpt(&q, "format", format)
	setOpt(&q, "advocate_token", advocateToken)
	setOpt(&q, "status", status)
	setOpt(&q, "created_range", createdRange)
	setOpt(&q, "updated_range", updatedRange)
	setOpt(&q, "fields", fields)
	setOpt(&q, "sort", sort)
	path := fmt.Sprintf("/accounts/%s/bonuses", url.PathEscape(accountSlug))
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

// GetBonus fetches a single bonus.
func (c *Client) GetBonus(accountSlug, bonusId string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/bonuses/%s", url.PathEscape(accountSlug), url.PathEscape(bonusId))
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

// PostBonus creates a bonus.
func (c *Client) PostBonus(accountSlug string, form BonusesForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/bonuses", url.PathEscape(accountSlug))
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

// ForceBonus creates a forced bonus.
func (c *Client) ForceBonus(accountSlug string, form ForceBonusesForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/bonuses/force", url.PathEscape(accountSlug))
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

// DeleteBonus deletes a bonus.
func (c *Client) DeleteBonus(accountSlug, bonusId string) error {
	path := fmt.Sprintf("/accounts/%s/bonuses/%s", url.PathEscape(accountSlug), url.PathEscape(bonusId))
	resp, body, err := c.request("DELETE", path, nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 206 {
		return newAPIError(resp.StatusCode, body)
	}
	return nil
}

// CheckupBonuses runs the bonus checkup.
func (c *Client) CheckupBonuses(accountSlug string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/bonuses/checkup", url.PathEscape(accountSlug))
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

// GetBonusTraces lists bonus traces.
func (c *Client) GetBonusTraces(accountSlug string, page, limit int64) (map[string]interface{}, error) {
	q := url.Values{}
	if page > 0 {
		q.Set("page", strconv.FormatInt(page, 10))
	}
	if limit > 0 {
		q.Set("limit", strconv.FormatInt(limit, 10))
	}
	path := fmt.Sprintf("/accounts/%s/bonuses/traces", url.PathEscape(accountSlug))
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

// GetBonusTrace fetches a single bonus trace.
func (c *Client) GetBonusTrace(accountSlug, traceId string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/bonuses/traces/%s", url.PathEscape(accountSlug), url.PathEscape(traceId))
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

// PatchBonus partially updates a bonus.
func (c *Client) PatchBonus(accountSlug, bonusId string, form BonusPatchForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/bonuses/%s", url.PathEscape(accountSlug), url.PathEscape(bonusId))
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

// GetBonusTags lists tags for a bonus.
func (c *Client) GetBonusTags(accountSlug, bonusId string, page, limit int64) (map[string]interface{}, error) {
	q := url.Values{}
	if page > 0 {
		q.Set("page", strconv.FormatInt(page, 10))
	}
	if limit > 0 {
		q.Set("limit", strconv.FormatInt(limit, 10))
	}
	path := fmt.Sprintf("/accounts/%s/bonuses/%s/tags", url.PathEscape(accountSlug), url.PathEscape(bonusId))
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

// AddBonusTag adds a tag to a bonus.
func (c *Client) AddBonusTag(accountSlug, bonusId string, form TagForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/bonuses/%s/tags", url.PathEscape(accountSlug), url.PathEscape(bonusId))
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

// RemoveBonusTag removes a tag from a bonus.
func (c *Client) RemoveBonusTag(accountSlug, bonusId, tagSlug string) error {
	path := fmt.Sprintf("/accounts/%s/bonuses/%s/tags/%s", url.PathEscape(accountSlug), url.PathEscape(bonusId), url.PathEscape(tagSlug))
	resp, body, err := c.request("DELETE", path, nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 206 {
		return newAPIError(resp.StatusCode, body)
	}
	return nil
}
