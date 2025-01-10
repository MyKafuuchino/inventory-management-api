package entity

type ResponseSuccess[T interface{}] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

type ResponseError struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Errors  []string `json:"error,omitempty"`
}

func NewResponseSuccess[T interface{}](message string, data T) *ResponseSuccess[T] {
	return &ResponseSuccess[T]{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func NewResponseError(message string, errors ...string) *ResponseError {
	var errorList []string
	if len(errors) > 0 {
		errorList = errors
	}
	return &ResponseError{
		Success: false,
		Message: message,
		Errors:  errorList,
	}
}
