package model

type CreateProductRequest struct {
	Name        string `json:"name" validate:"required,max=100"`
	Description string `json:"description" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Stock       int    `json:"stock" validate:"gte=0"`
	LowStock    int    `json:"low_stock" validate:"gte=0"`
}

type UpdateProductRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Price       *int    `json:"price,omitempty"`
	Stock       *int    `json:"stock,omitempty"`
	LowStock    *int    `json:"low_stock,omitempty"`
}
