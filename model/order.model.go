package model

import "time"

type CreateOrderRequest struct {
	UserID       uint                 `json:"user_id"`
	OrderStatus  string               `json:"order_status,omitempty" validate:"oneof=pending processed completed canceled"`
	OrderDetails []OrderDetailRequest `json:"order_details" validate:"required"`
}

type OrderDetailRequest struct {
	OrderID   uint `json:"order_id"`
	ProductID uint `json:"product_id" validate:"required"`
	Quantity  int  `json:"quantity" validate:"required"`
	Price     int  `json:"price"`
}

type OrderResponse struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	TotalPrice  int    `json:"total_price"`
	OrderStatus string `json:"order_status"`

	OrderDetail []OrderDetailResponse `json:"order_detail"`
}

type OrderDetailResponse struct {
	ProductName string `json:"product_name"`
	ProductID   uint   `json:"product_id"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
}

type GetAllOrdersResponse struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	TotalPrice  int       `json:"total_price"`
	OrderStatus string    `json:"order_status"`
	CreatedAt   time.Time `json:"created_at"`
}
