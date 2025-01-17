package model

type CreateTransactionRequest struct {
	OrderID       int    `json:"orderId" validate:"required"`
	UserID        uint   `json:"user_id" validate:"required"`
	PaymentMethod string `json:"payment_method" validate:"required"`
	Status        string `json:"status" validate:"oneof=pending processed completed canceled" `
}
