package repository

import (
	"fmt"
	"gorm.io/gorm"
	"inventory-management/entity"
)

type OrderRepository interface {
	GetAllOrder() ([]entity.Order, error)
	GetOrderById(orderID string) (*entity.Order, error)
	CreateNewOrder(order *entity.Order) (*entity.Order, error)
	UpdateOrderByID(orderID string, order *entity.Order) (*entity.Order, error)
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

func (r *orderRepository) GetOrderById(orderID string) (*entity.Order, error) {
	order := &entity.Order{}
	var err error
	if err = r.db.Table("orders").Where("id = ?", orderID).First(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *orderRepository) CreateNewOrder(order *entity.Order) (*entity.Order, error) {
	if err := r.db.Table("orders").Create(order).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return order, nil
}

func (r *orderRepository) UpdateOrderByID(orderID string, order *entity.Order) (*entity.Order, error) {
	if err := r.db.Save(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}
