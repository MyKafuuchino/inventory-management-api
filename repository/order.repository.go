package repository

import (
	"gorm.io/gorm"
	"inventory-management/entity"
)

type OrderRepository interface {
	GetAllOrder() ([]entity.Order, error)
	CreateNewOrder(order *entity.Order) (*entity.Order, error)
	GetOrderById(orderId string) (*entity.Order, error)
	UpdateOrderStatus(order *entity.Order) (*entity.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) GetAllOrder() ([]entity.Order, error) {
	var orders []entity.Order
	if err := r.db.Table("orders").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderRepository) CreateNewOrder(order *entity.Order) (*entity.Order, error) {
	if err := r.db.Table("orders").Create(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *orderRepository) GetOrderById(orderId string) (*entity.Order, error) {
	order := &entity.Order{}
	if err := r.db.Table("orders").Where("id = ?", orderId).First(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *orderRepository) UpdateOrderStatus(order *entity.Order) (*entity.Order, error) {
	if err := r.db.Table("orders").Save(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}
