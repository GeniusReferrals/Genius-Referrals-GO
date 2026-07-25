package geniusreferrals

import "encoding/json"

// GetAuthentication verifies the API token.
func (c *Client) GetAuthentication() (map[string]interface{}, error) {
	resp, body, err := c.request("GET", "/test-authentication", nil, nil)
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
