package errors

import (
	"fmt"
)

type FileNotProvidedError struct{}

func (err FileNotProvidedError) Error() string {
	return "No file is provided"
}

type InvalidHTTPResponse struct {
	Url        string
	StatusCode int
	Message    string
}

func (err InvalidHTTPResponse) Error() string {
	return fmt.Sprintf(
		"Error when making request to %s\nResponse status code: %d\nResponse message: %v",
		err.Url,
		err.StatusCode,
		err.Message,
	)
}
