package entity

import "time"

type Product struct {
	Base
	Name        string    `gorm:"not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Price       int       `gorm:"not null" json:"price"`
	Stock       int       `gorm:"default:0" json:"stock"`
	LowStock    int       `gorm:"default:0" json:"low_stock"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
