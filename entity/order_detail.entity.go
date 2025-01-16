package entity

import "inventory-management/model"

type OrderDetail struct {
	OrderID   uint `gorm:"primaryKey" json:"order_id"`
	ProductID uint `gorm:"primaryKey" json:"product_id"`
	Quantity  int  `json:"quantity"`
	Price     int  `json:"price"`
	model.History
}

func (OrderDetail) TableName() string {
	return "order_details"
}
