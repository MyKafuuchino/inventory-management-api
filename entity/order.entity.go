package entity

import "inventory-management/model"

type Order struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	UserID      uint   `json:"user_id"`
	OrderStatus string `gorm:"default:pending" json:"order_status"`
	TotalPrice  int    `json:"total_price"`

	Product     []Product   `gorm:"many2many:order_details" json:"product,omitempty"`
	Transaction Transaction `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"transactions,omitempty"`
	model.History
}
