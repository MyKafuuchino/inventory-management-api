package entity

import "inventory-management/model"

type Order struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	UserID     uint   `json:"user_id"`
	TotalPrice int    `json:"total_price"`
	Status     string `gorm:"default:'pending'" json:"status"`

	Product []Product `gorm:"many2many:order_details" json:"product,omitempty"`
	model.History
}
