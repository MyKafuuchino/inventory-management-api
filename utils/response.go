package utils

type ResponseSuccess[T interface{}] struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	Total      int64  `json:"total,omitempty"`
	TotalPages int    `json:"totalPages,omitempty"`
	Page       int    `json:"page,omitempty"`
	PageSize   int    `json:"pageSize,omitempty"`
	Data       T      `json:"data,omitempty"`
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

func NewPaginatedResponse[T interface{}](
	message string,
	data T,
	total int64,
	totalPages int,
	page int,
	pageSize int,
) *ResponseSuccess[T] {
	return &ResponseSuccess[T]{
		Success:    true,
		Message:    message,
		Data:       data,
		Total:      total,
		TotalPages: totalPages,
		Page:       page,
		PageSize:   pageSize,
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
