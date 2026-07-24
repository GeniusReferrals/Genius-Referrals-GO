package geniusreferrals

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// GetWidgetsPackages lists widgets packages.
func (c *Client) GetWidgetsPackages(accountSlug string, page, limit int64) (map[string]interface{}, error) {
	q := url.Values{}
	if page > 0 {
		q.Set("page", strconv.FormatInt(page, 10))
	}
	if limit > 0 {
		q.Set("limit", strconv.FormatInt(limit, 10))
	}
	path := fmt.Sprintf("/accounts/%s/widgets-packages", url.PathEscape(accountSlug))
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

// GetWidgetsPackage fetches a single widgets package.
func (c *Client) GetWidgetsPackage(accountSlug, widgetsPackageSlug string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/widgets-packages/%s", url.PathEscape(accountSlug), url.PathEscape(widgetsPackageSlug))
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

// CreateWidgetsPackage creates a widgets package.
func (c *Client) CreateWidgetsPackage(accountSlug string, form WidgetsPackageForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/widgets-packages", url.PathEscape(accountSlug))
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

// UpdateWidgetsPackage updates a widgets package.
func (c *Client) UpdateWidgetsPackage(accountSlug, widgetsPackageSlug string, form WidgetsPackageForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/widgets-packages/%s", url.PathEscape(accountSlug), url.PathEscape(widgetsPackageSlug))
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

// DeleteWidgetsPackage deletes a widgets package.
func (c *Client) DeleteWidgetsPackage(accountSlug, widgetsPackageSlug string) error {
	path := fmt.Sprintf("/accounts/%s/widgets-packages/%s", url.PathEscape(accountSlug), url.PathEscape(widgetsPackageSlug))
	resp, body, err := c.request("DELETE", path, nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 206 {
		return newAPIError(resp.StatusCode, body)
	}
	return nil
}

// GetWidgets lists widgets in a package.
func (c *Client) GetWidgets(accountSlug, widgetsPackageSlug string, page, limit int64) (map[string]interface{}, error) {
	q := url.Values{}
	if page > 0 {
		q.Set("page", strconv.FormatInt(page, 10))
	}
	if limit > 0 {
		q.Set("limit", strconv.FormatInt(limit, 10))
	}
	path := fmt.Sprintf("/accounts/%s/widgets-packages/%s/widgets", url.PathEscape(accountSlug), url.PathEscape(widgetsPackageSlug))
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

// GetWidget fetches a single widget.
func (c *Client) GetWidget(accountSlug, widgetsPackageSlug, widgetId string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/widgets-packages/%s/widgets/%s", url.PathEscape(accountSlug), url.PathEscape(widgetsPackageSlug), url.PathEscape(widgetId))
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

// CreateWidget creates a widget.
func (c *Client) CreateWidget(accountSlug, widgetsPackageSlug string, form WidgetForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/widgets-packages/%s/widgets", url.PathEscape(accountSlug), url.PathEscape(widgetsPackageSlug))
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

// PatchWidget updates a widget.
func (c *Client) PatchWidget(accountSlug, widgetsPackageSlug, widgetId string, form WidgetForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/widgets-packages/%s/widgets/%s", url.PathEscape(accountSlug), url.PathEscape(widgetsPackageSlug), url.PathEscape(widgetId))
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

// PutWidget replaces a widget.
func (c *Client) PutWidget(accountSlug, widgetsPackageSlug, widgetId string, form WidgetForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/widgets-packages/%s/widgets/%s", url.PathEscape(accountSlug), url.PathEscape(widgetsPackageSlug), url.PathEscape(widgetId))
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

// GetWidgetTranslations lists widget translations.
func (c *Client) GetWidgetTranslations(accountSlug, widgetsPackageSlug, widgetId string, page, limit int64) (map[string]interface{}, error) {
	q := url.Values{}
	if page > 0 {
		q.Set("page", strconv.FormatInt(page, 10))
	}
	if limit > 0 {
		q.Set("limit", strconv.FormatInt(limit, 10))
	}
	path := fmt.Sprintf("/accounts/%s/widgets-packages/%s/widgets/%s/translations", url.PathEscape(accountSlug), url.PathEscape(widgetsPackageSlug), url.PathEscape(widgetId))
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

// GetWidgetTranslation fetches a single widget translation.
func (c *Client) GetWidgetTranslation(accountSlug, widgetsPackageSlug, widgetId, locale string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/widgets-packages/%s/widgets/%s/translations/%s", url.PathEscape(accountSlug), url.PathEscape(widgetsPackageSlug), url.PathEscape(widgetId), url.PathEscape(locale))
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

// CreateWidgetTranslation creates a widget translation.
func (c *Client) CreateWidgetTranslation(accountSlug, widgetsPackageSlug, widgetId string, form WidgetTranslationForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/widgets-packages/%s/widgets/%s/translations", url.PathEscape(accountSlug), url.PathEscape(widgetsPackageSlug), url.PathEscape(widgetId))
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

// PatchWidgetTranslation updates a widget translation.
func (c *Client) PatchWidgetTranslation(accountSlug, widgetsPackageSlug, widgetId, locale string, form WidgetTranslationForm) (map[string]interface{}, error) {
	path := fmt.Sprintf("/accounts/%s/widgets-packages/%s/widgets/%s/translations/%s", url.PathEscape(accountSlug), url.PathEscape(widgetsPackageSlug), url.PathEscape(widgetId), url.PathEscape(locale))
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
