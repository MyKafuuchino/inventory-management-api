package entity

type OrderDetail struct {
	OrderID   string `gorm:"primaryKey;no null" json:"order_id" validate:"required"`
	ProductID string `gorm:"primaryKey;no null" json:"product_id" validate:"required"`
	Quantity  int    `gorm:"no null" json:"quantity" validate:"required"`
	Price     int    `gorm:"no null" json:"price" validate:"required"`
	CreatedAt int64  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64  `gorm:"autoUpdateTime" json:"updated_at"`
}
