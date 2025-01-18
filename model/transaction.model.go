package model

type CreateTransactionRequest struct {
	OrderID       int    `json:"orderId" validate:"required"`
	UserID        uint   `json:"user_id" validate:"required"`
	PaymentMethod string `json:"payment_method" validate:"required"`
	OrderStatus   string `json:"order_status" validate:"oneof=pending processed completed canceled" `
}

type UpdateTransactionRequest struct {
	PaymentMethod     string `json:"payment_method" validate:"required"`
	OrderStatus       string `json:"order_status" validate:"oneof=pending processed completed canceled"`
	TransactionStatus string `json:"transaction_status" validate:"oneof=paid unpaid"`
}
