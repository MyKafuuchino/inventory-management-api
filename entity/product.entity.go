package entity

import "time"

type Product struct {
	Base
	Name        string    `gorm:"not null" json:"name" validate:"required,max=100"`
	Description string    `gorm:"type:text" json:"description" validate:"required"`
	Price       int       `gorm:"not null" json:"price" validate:"required"`
	Stock       int       `gorm:"default:0" json:"stock" validate:"gte=0"`
	LowStock    int       `gorm:"default:0" json:"low_stock" validate:"gte=0"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Orders []Order `gorm:"many2many:user_groups" json:"orders,omitempty"`
}
