package entity

import (
	"time"
)

type Order struct {
	Base
	UserID     string    `json:"user_id"`
	TotalPrice int       `json:"total_price"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	OrderDetails []OrderDetail `gorm:"foreignKey:OrderID" json:"order_details,omitempty"`
}
