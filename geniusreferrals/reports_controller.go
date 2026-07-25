package geniusreferrals

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// reportRequest performs a GET report request.
func (c *Client) reportRequest(path string, accountSlug, startDate, endDate, advocateToken string) (map[string]interface{}, error) {
	q := url.Values{}
	if accountSlug != "" {
		q.Set("account_slug", accountSlug)
	}
	if startDate != "" {
		q.Set("start_date", startDate)
	}
	if endDate != "" {
		q.Set("end_date", endDate)
	}
	if advocateToken != "" {
		q.Set("advocate_token", advocateToken)
	}
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

// GetBonusesDailyGiven report.
func (c *Client) GetBonusesDailyGiven(accountSlug, startDate, endDate, advocateToken string) (map[string]interface{}, error) {
	return c.reportRequest("/reports/bonuses-daily-given", accountSlug, startDate, endDate, advocateToken)
}

// GetBonusesSummaryPerOrigin report.
func (c *Client) GetBonusesSummaryPerOrigin(accountSlug, startDate, endDate, advocateToken string) (map[string]interface{}, error) {
	return c.reportRequest("/reports/bonuses-summary-per-origin", accountSlug, startDate, endDate, advocateToken)
}

// GetClickDailyParticipation report.
func (c *Client) GetClickDailyParticipation(accountSlug, startDate, endDate, advocateToken string) (map[string]interface{}, error) {
	return c.reportRequest("/reports/click-daily-participation", accountSlug, startDate, endDate, advocateToken)
}

// GetReferralDailyParticipation report.
func (c *Client) GetReferralDailyParticipation(accountSlug, startDate, endDate, advocateToken string) (map[string]interface{}, error) {
	return c.reportRequest("/reports/referral-daily-participation", accountSlug, startDate, endDate, advocateToken)
}

// GetReferralsSummaryPerOrigin report.
func (c *Client) GetReferralsSummaryPerOrigin(accountSlug, startDate, endDate, advocateToken string) (map[string]interface{}, error) {
	return c.reportRequest("/reports/referrals-summary-per-origin", accountSlug, startDate, endDate, advocateToken)
}

// GetShareDailyParticipation report.
func (c *Client) GetShareDailyParticipation(accountSlug, startDate, endDate, advocateToken string) (map[string]interface{}, error) {
	return c.reportRequest("/reports/share-daily-participation", accountSlug, startDate, endDate, advocateToken)
}

// GetTopAdvocates report.
func (c *Client) GetTopAdvocates(accountSlug, startDate, endDate, advocateToken string, limit int64) (map[string]interface{}, error) {
	q := url.Values{}
	if accountSlug != "" {
		q.Set("account_slug", accountSlug)
	}
	if startDate != "" {
		q.Set("start_date", startDate)
	}
	if endDate != "" {
		q.Set("end_date", endDate)
	}
	if advocateToken != "" {
		q.Set("advocate_token", advocateToken)
	}
	if limit > 0 {
		q.Set("limit", strconv.FormatInt(limit, 10))
	}
	resp, body, err := c.request("GET", "/reports/top-advocates", q, nil)
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

// GetRevenue report.
func (c *Client) GetRevenue(accountSlug, startDate, endDate string) (map[string]interface{}, error) {
	q := url.Values{}
	if accountSlug != "" {
		q.Set("account_slug", accountSlug)
	}
	if startDate != "" {
		q.Set("start_date", startDate)
	}
	if endDate != "" {
		q.Set("end_date", endDate)
	}
	resp, body, err := c.request("GET", "/reports/revenue", q, nil)
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

// Get1099TaxReport report.
func (c *Client) Get1099TaxReport(accountSlug, year string) (map[string]interface{}, error) {
	q := url.Values{}
	if accountSlug != "" {
		q.Set("account_slug", accountSlug)
	}
	if year != "" {
		q.Set("year", year)
	}
	resp, body, err := c.request("GET", "/reports/1099-tax-report", q, nil)
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
