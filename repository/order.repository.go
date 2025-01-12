package repository

import (
	"errors"
	"gorm.io/gorm"
	"inventory-management/entity"
)

type OrderRepository interface {
	GetAllOrder() ([]entity.Order, error)
	GetOrderById(orderID string) (*entity.Order, error)
	CreateNewOrder(order *entity.Order) (*entity.Order, error)
	UpdateOrderByID(orderID string, order *entity.Order) (*entity.Order, error)
	DeleteOrderById(orderID string) error
}

type orderRepository struct {
	db    *gorm.DB
	query *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) QueryInit() {
	r.query = r.db.Table("orders")
}

func (r *orderRepository) GetAllOrder() ([]entity.Order, error) {
	var orders []entity.Order
	err := r.query.Find(&orders).Error

	if len(orders) == 0 {
		return nil, errors.New("no orders found")
	}

	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderRepository) GetOrderById(orderID string) (*entity.Order, error) {
	var order *entity.Order
	err := r.query.Where("id = ?", orderID).First(order).Error
	if err != nil {
		return nil, errors.New("order not found")
	}
	return order, nil
}

func (r *orderRepository) CreateNewOrder(order *entity.Order) (*entity.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (r *orderRepository) UpdateOrderByID(orderID string, order *entity.Order) (*entity.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (r *orderRepository) DeleteOrderById(orderID string) error {
	//TODO implement me
	panic("implement me")
}
