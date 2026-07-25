package geniusreferrals

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// GetVouchers lists vouchers.
func (c *Client) GetVouchers(accountSlug string, page, limit int64) (map[string]interface{}, error) {
	q := url.Values{}
	if page > 0 {
		q.Set("page", strconv.FormatInt(page, 10))
	}
	if limit > 0 {
		q.Set("limit", strconv.FormatInt(limit, 10))
	}
	path := fmt.Sprintf("/accounts/%s/vouchers", url.PathEscape(accountSlug))
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

// GetVoucher fetches a single voucher.
func (c *Client) GetVoucher(accountSlug, voucherId string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/vouchers/%s", url.PathEscape(accountSlug), url.PathEscape(voucherId))
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

// CreateVoucher creates a voucher.
func (c *Client) CreateVoucher(accountSlug string, form VoucherForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/vouchers", url.PathEscape(accountSlug))
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

// PatchVoucher updates a voucher.
func (c *Client) PatchVoucher(accountSlug, voucherId string, form VoucherForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/vouchers/%s", url.PathEscape(accountSlug), url.PathEscape(voucherId))
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

// DeleteVoucher deletes a voucher.
func (c *Client) DeleteVoucher(accountSlug, voucherId string) error {
	path := fmt.Sprintf("/accounts/%s/vouchers/%s", url.PathEscape(accountSlug), url.PathEscape(voucherId))
	resp, body, err := c.request("DELETE", path, nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 206 {
		return newAPIError(resp.StatusCode, body)
	}
	return nil
}

// GetVoucherDenominations gets voucher denominations for a currency.
func (c *Client) GetVoucherDenominations(accountSlug, currencyCode string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/vouchers-denominations/%s", url.PathEscape(accountSlug), url.PathEscape(currencyCode))
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
