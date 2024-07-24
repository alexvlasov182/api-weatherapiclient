package client

import "fmt"

// APIError represents an error response from the API.
type APIError struct {
	Code    int
	Message string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error: code: %d, message %s", e.Code, e.Message)
}
