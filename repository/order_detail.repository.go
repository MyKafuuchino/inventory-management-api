package repository

import (
	"gorm.io/gorm"
	"inventory-management/entity"
)

type OrderDetailRepository interface {
	GetOrderDetailByID(orderID string) (*entity.Order, error)
}

type orderDetailRepository struct {
	db *gorm.DB
}

func NewOrderDetailRepository(db *gorm.DB) OrderDetailRepository {
	return &orderDetailRepository{db: db}
}

func (r orderDetailRepository) GetOrderDetailByID(orderID string) (*entity.Order, error) {
	var order *entity.Order

	if err := r.db.Preload("OrderDetails.Product").Where("id = ?", orderID).First(&order).Error; err != nil {
		return nil, err
	}

	return order, nil
}
