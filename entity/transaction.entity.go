package entity

import "time"

type Transaction struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	OrderID       uint      `json:"order_id"`
	UserID        uint      `json:"user_id"`
	TotalPrice    int       `json:"total_price"`
	PaymentMethod string    `json:"payment_method"`
	TransactionAt time.Time `json:"transaction_date"`
}
