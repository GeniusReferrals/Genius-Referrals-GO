package geniusreferrals

import "fmt"

// APIError represents an error response from the Genius Referrals API.
type APIError struct {
	Message    string
	StatusCode int
	Body       []byte
}

func (e *APIError) Error() string {
	return fmt.Sprintf("api error %d: %s", e.StatusCode, e.Message)
}

// errorMessages maps common HTTP status codes to readable messages.
var errorMessages = map[int]string{
	400: "Bad request",
	401: "You are not authenticated",
	403: "User not authorized to perform the operation",
	404: "Resource not found",
	409: "Conflict",
	422: "Validation failed",
	429: "Rate limit exceeded",
}

func newAPIError(statusCode int, body []byte) *APIError {
	msg := errorMessages[statusCode]
	if msg == "" {
		msg = "HTTP Response Not OK"
	}
	return &APIError{Message: msg, StatusCode: statusCode, Body: body}
}
