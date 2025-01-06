package entity

import "fmt"

type CustomError struct {
	StatusCode int
	Message    string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Status %d, Message %s\n", e.StatusCode, e.Message)
}

func NewCustomError(statusCode int, message string) *CustomError {
	return &CustomError{StatusCode: statusCode, Message: message}
}
