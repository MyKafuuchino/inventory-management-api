package repository

import (
	"gorm.io/gorm"
	"inventory-management/entity"
	"time"
)

type OrderRepository interface {
	GetAllOrders() ([]entity.Order, error)
	GetOrderByID(orderID uint) (*entity.Order, error)
	CreateOrderWithDetail(reqOrder *entity.Order, orderDetails []entity.OrderDetail) error
	UpdateOrderStatus(reqOrder *entity.Order) error
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

func (r orderRepository) UpdateOrderStatus(reqOrder *entity.Order) error {
	tx := r.db.Begin()

	if err := tx.Model(&entity.Order{}).
		Where("id = ?", reqOrder.ID).
		Updates(map[string]interface{}{
			"order_status": reqOrder.OrderStatus,
			"total_price":  reqOrder.TotalPrice,
		}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&entity.Transaction{}).
		Where("order_id = ?", reqOrder.ID).
		Updates(map[string]interface{}{
			"payment_status": reqOrder.Transaction.PaymentStatus,
			"payment_method": reqOrder.Transaction.PaymentMethod,
			"transaction_at": time.Now(),
		}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
