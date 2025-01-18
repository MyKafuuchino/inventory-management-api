package entity

import "time"

type Transaction struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	OrderID       uint      `gorm:"unique" json:"order_id"`
	UserID        uint      `json:"user_id"`
	PaymentStatus string    `gorm:"default:unpaid" json:"payment_status"`
	PaymentMethod string    `json:"payment_method"`
	TransactionAt time.Time `json:"transaction_date"`
}
