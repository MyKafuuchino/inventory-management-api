package entity

import "time"

type Order struct {
	Base
	TotalPrice int       `gorm:"not null;" json:"total_price" validate:"required"`
	Status     string    `gorm:"default:pending" json:"status" validate:"omitempty,oneof=pending processed completed canceled"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	UserID   string    `gorm:"not null" json:"userId" validate:"required"`
	Products []Product `gorm:"many2many:order_detail;" json:"products;omitempty"`
}
