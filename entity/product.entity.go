package entity

import "inventory-management/model"

type Product struct {
	ID          uint   `gorm:"primarykey" json:"id"`
	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	Price       int    `gorm:"not null" json:"price"`
	Stock       int    `gorm:"default:0" json:"stock"`
	LowStock    int    `gorm:"default:0" json:"low_stock"`
	model.History
}
