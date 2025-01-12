package utils

import "fmt"

type CustomError struct {
	StatusCode int
	Message    string
	Errors     []string `json:"omitempty"`
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Status %d, Message %s\n", e.StatusCode, e.Message)
}

func NewCustomError(statusCode int, message string, errors ...string) *CustomError {
	var errorList []string
	if len(errors) > 0 {
		errorList = errors
	}
	return &CustomError{StatusCode: statusCode, Message: message, Errors: errorList}
}
