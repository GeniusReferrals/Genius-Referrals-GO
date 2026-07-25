package geniusreferrals

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// GetProducts lists products.
func (c *Client) GetProducts(accountSlug string, page, limit int64) (map[string]interface{}, error) {
	q := url.Values{}
	if page > 0 {
		q.Set("page", strconv.FormatInt(page, 10))
	}
	if limit > 0 {
		q.Set("limit", strconv.FormatInt(limit, 10))
	}
	path := fmt.Sprintf("/accounts/%s/products", url.PathEscape(accountSlug))
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

// GetProduct fetches a single product.
func (c *Client) GetProduct(accountSlug, productSlug string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/products/%s", url.PathEscape(accountSlug), url.PathEscape(productSlug))
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

// CreateProduct creates a product.
func (c *Client) CreateProduct(accountSlug string, form ProductForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/products", url.PathEscape(accountSlug))
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

// UpdateProduct updates a product.
func (c *Client) UpdateProduct(accountSlug, productSlug string, form ProductForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/products/%s", url.PathEscape(accountSlug), url.PathEscape(productSlug))
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

// DeleteProduct deletes a product.
func (c *Client) DeleteProduct(accountSlug, productSlug string) error {
	path := fmt.Sprintf("/accounts/%s/products/%s", url.PathEscape(accountSlug), url.PathEscape(productSlug))
	resp, body, err := c.request("DELETE", path, nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 206 {
		return newAPIError(resp.StatusCode, body)
	}
	return nil
}

// GetProductVariants lists variants for a product.
func (c *Client) GetProductVariants(accountSlug, productSlug string, page, limit int64) (map[string]interface{}, error) {
	q := url.Values{}
	if page > 0 {
		q.Set("page", strconv.FormatInt(page, 10))
	}
	if limit > 0 {
		q.Set("limit", strconv.FormatInt(limit, 10))
	}
	path := fmt.Sprintf("/accounts/%s/products/%s/variants", url.PathEscape(accountSlug), url.PathEscape(productSlug))
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

// CreateProductVariant creates a product variant.
func (c *Client) CreateProductVariant(accountSlug, productSlug string, form ProductVariantForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/products/%s/variants", url.PathEscape(accountSlug), url.PathEscape(productSlug))
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

// PatchProductVariant updates a product variant.
func (c *Client) PatchProductVariant(accountSlug, productSlug, variantId string, form ProductVariantForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/products/%s/variants/%s", url.PathEscape(accountSlug), url.PathEscape(productSlug), url.PathEscape(variantId))
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

// DeleteProductVariant deletes a product variant.
func (c *Client) DeleteProductVariant(accountSlug, productSlug, variantId string) error {
	path := fmt.Sprintf("/accounts/%s/products/%s/variants/%s", url.PathEscape(accountSlug), url.PathEscape(productSlug), url.PathEscape(variantId))
	resp, body, err := c.request("DELETE", path, nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 206 {
		return newAPIError(resp.StatusCode, body)
	}
	return nil
}

// PatchProductVariantPrice updates a variant price.
func (c *Client) PatchProductVariantPrice(accountSlug, productSlug, variantId, currencyCode string, form ProductVariantPriceForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/products/%s/variants/%s/prices/%s", url.PathEscape(accountSlug), url.PathEscape(productSlug), url.PathEscape(variantId), url.PathEscape(currencyCode))
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
