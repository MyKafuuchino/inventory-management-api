package entity

import "time"

type Product struct {
	Base
	Name        string    `json:"name" validate:"required,max=100"`
	Description string    `json:"description" validate:"required"`
	Price       int       `json:"price" validate:"required"`
	Stock       int       `gorm:"default:0" json:"stock" validate:"gte=0"`
	LowStock    int       `gorm:"default:0" json:"low_stock" validate:"gte=0"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
