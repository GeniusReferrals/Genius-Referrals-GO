package geniusreferrals

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// GetAdvocatesParam holds optional query parameters for listing advocates.
type GetAdvocatesParam struct {
	Page             int64
	Limit            int64
	Format           string
	Email            string
	Name             string
	Lastname         string
	Code             string
	Token            string
	Status           string
	Fraudulent       *bool
	IsEmailConfirmed *bool
	CanRefer         *bool
	UpdatedRange     string
	CreatedRange     string
	Fields           string
	Sort             string
}

// GetAdvocates lists advocates.
func (c *Client) GetAdvocates(accountSlug string, p GetAdvocatesParam) (map[string]interface{}, error) {
	q := url.Values{}
	if p.Page > 0 {
		q.Set("page", strconv.FormatInt(p.Page, 10))
	}
	if p.Limit > 0 {
		q.Set("limit", strconv.FormatInt(p.Limit, 10))
	}
	setOpt(&q, "format", p.Format)
	setOpt(&q, "email", p.Email)
	setOpt(&q, "name", p.Name)
	setOpt(&q, "lastname", p.Lastname)
	setOpt(&q, "code", p.Code)
	setOpt(&q, "token", p.Token)
	setOpt(&q, "status", p.Status)
	setBoolOpt(&q, "fraudulent", p.Fraudulent)
	setBoolOpt(&q, "is_email_confirmed", p.IsEmailConfirmed)
	setBoolOpt(&q, "can_refer", p.CanRefer)
	setOpt(&q, "updated_range", p.UpdatedRange)
	setOpt(&q, "created_range", p.CreatedRange)
	setOpt(&q, "fields", p.Fields)
	setOpt(&q, "sort", p.Sort)
	path := fmt.Sprintf("/accounts/%s/advocates", url.PathEscape(accountSlug))
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

func setOpt(q *url.Values, key, value string) {
	if value != "" {
		q.Set(key, value)
	}
}
func setBoolOpt(q *url.Values, key string, value *bool) {
	if value != nil {
		q.Set(key, strconv.FormatBool(*value))
	}
}

// GetAdvocate fetches a single advocate.
func (c *Client) GetAdvocate(accountSlug, advocateToken string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/advocates/%s", url.PathEscape(accountSlug), url.PathEscape(advocateToken))
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

// PostAdvocate creates an advocate.
func (c *Client) PostAdvocate(accountSlug string, form AdvocateForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/advocates", url.PathEscape(accountSlug))
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

// PutAdvocate updates an advocate.
func (c *Client) PutAdvocate(accountSlug, advocateToken string, form AdvocateForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/advocates/%s", url.PathEscape(accountSlug), url.PathEscape(advocateToken))
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

// PatchAdvocate partially updates an advocate.
func (c *Client) PatchAdvocate(accountSlug, advocateToken string, form AdvocatePatchForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/advocates/%s", url.PathEscape(accountSlug), url.PathEscape(advocateToken))
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

// DeleteAdvocate deletes a single advocate.
func (c *Client) DeleteAdvocate(accountSlug, advocateToken string) error {
	path := fmt.Sprintf("/accounts/%s/advocates/%s", url.PathEscape(accountSlug), url.PathEscape(advocateToken))
	resp, body, err := c.request("DELETE", path, nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 206 {
		return newAPIError(resp.StatusCode, body)
	}
	return nil
}

// DeleteAdvocates deletes multiple advocates.
func (c *Client) DeleteAdvocates(accountSlug string, advocateTokens []string) error {
	q := url.Values{}
	if len(advocateTokens) > 0 {
		q.Set("advocates", strings.Join(advocateTokens, ","))
	}
	path := fmt.Sprintf("/accounts/%s/advocates", url.PathEscape(accountSlug))
	resp, body, err := c.request("DELETE", path, q, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 206 {
		return newAPIError(resp.StatusCode, body)
	}
	return nil
}

// GetAdvocateShareLinks gets share links for an advocate.
func (c *Client) GetAdvocateShareLinks(accountSlug, advocateToken string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/advocates/%s/share-links", url.PathEscape(accountSlug), url.PathEscape(advocateToken))
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

// GetAdvocatePaymentMethods lists payment methods for an advocate.
func (c *Client) GetAdvocatePaymentMethods(accountSlug, advocateToken string, page, limit int64) (map[string]interface{}, error) {
	q := url.Values{}
	if page > 0 {
		q.Set("page", strconv.FormatInt(page, 10))
	}
	if limit > 0 {
		q.Set("limit", strconv.FormatInt(limit, 10))
	}
	path := fmt.Sprintf("/accounts/%s/advocates/%s/payment-methods", url.PathEscape(accountSlug), url.PathEscape(advocateToken))
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

// GetAdvocatePaymentMethod fetches a single payment method.
func (c *Client) GetAdvocatePaymentMethod(accountSlug, advocateToken, paymentMethodId string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/advocates/%s/payment-methods/%s", url.PathEscape(accountSlug), url.PathEscape(advocateToken), url.PathEscape(paymentMethodId))
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

// PostAdvocatePaymentMethod creates a payment method.
func (c *Client) PostAdvocatePaymentMethod(accountSlug, advocateToken string, form PaymentMethodForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/advocates/%s/payment-methods", url.PathEscape(accountSlug), url.PathEscape(advocateToken))
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

// PutAdvocatePaymentMethod updates a payment method.
func (c *Client) PutAdvocatePaymentMethod(accountSlug, advocateToken, paymentMethodId string, form PaymentMethodForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/advocates/%s/payment-methods/%s", url.PathEscape(accountSlug), url.PathEscape(advocateToken), url.PathEscape(paymentMethodId))
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
