package repository

import (
	"gorm.io/gorm"
	"inventory-management/entity"
)

type OrderRepository interface {
	GetAllOrders() ([]entity.Order, error)
	GetOrderByID(orderID uint) (*entity.Order, error)
	CreateOrderWithDetail(reqOrder *entity.Order, orderDetails []entity.OrderDetail) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r orderRepository) GetAllOrders() ([]entity.Order, error) {
	var orders []entity.Order
	if err := r.db.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r orderRepository) GetOrderByID(orderID uint) (*entity.Order, error) {
	var order entity.Order
	if err := r.db.First(&order, "id = ?", orderID).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r orderRepository) CreateOrderWithDetail(reqOrder *entity.Order, orderDetails []entity.OrderDetail) error {
	tx := r.db.Begin()

	if err := tx.Create(reqOrder).Error; err != nil {
		tx.Rollback()
		return err
	}

	for i := range orderDetails {
		orderDetails[i].OrderID = reqOrder.ID
	}

	if err := tx.Create(orderDetails).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
