package model

import "inventory-management/entity"

type CreateOrderRequest struct {
	Status     string         `json:"status" validate:"omitempty,oneof=pending processed completed canceled"`
	UserID     string         `json:"userId"`
	Quantities map[string]int `json:"quantities" validate:"required"`
	ProductIDs []string       `json:"productIds" validate:"required,min=1,dive"`
}

type UpdateOrderStatusRequest struct {
	Status string `json:"status" validate:"omitempty,oneof=pending processed completed canceled"`
}

type OrderResponse struct {
	ID           string                `json:"id"`
	UserID       string                `json:"user_id"`
	TotalPrice   int                   `json:"total_price"`
	Status       string                `json:"status"`
	CreatedAt    string                `json:"created_at"`
	UpdatedAt    string                `json:"updated_at"`
	OrderDetails []OrderDetailResponse `json:"order_details"`
}

type OrderDetailResponse struct {
	OrderID   string         `json:"order_id"`
	ProductID string         `json:"product_id"`
	Quantity  int            `json:"quantity"`
	Price     int            `json:"price"`
	Product   entity.Product `json:"product"`
}
