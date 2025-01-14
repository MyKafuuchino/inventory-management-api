package entity

type OrderDetail struct {
	OrderID   string `gorm:"primaryKey;index" json:"order_id"`
	ProductID string `gorm:"primaryKey;index" json:"product_id"`
	Quantity  int    `json:"quantity"`
	Price     int    `json:"price"`

	Product Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Order   Order   `gorm:"foreignKey:OrderID" json:"order,omitempty"`
}
