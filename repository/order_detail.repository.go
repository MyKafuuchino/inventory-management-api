package repository

import (
	"fmt"
	"gorm.io/gorm"
	"inventory-management/entity"
	"inventory-management/model"
)

type OrderDetailRepository interface {
	GetOrderDetailsByOrderID(orderID uint) ([]entity.OrderDetail, error)
	CreateBulkOrderDetails(orderDetails []entity.OrderDetail) error
	GetOrderWithDetailsByID(orderID uint) (*entity.Order, []model.OrderDetailResponse, error)
}
type orderDetailRepository struct {
	db *gorm.DB
}

func NewOrderDetailRepository(db *gorm.DB) OrderDetailRepository {
	return &orderDetailRepository{db: db}
}

func (r *orderDetailRepository) GetOrderDetailsByOrderID(orderID uint) ([]entity.OrderDetail, error) {
	var orderDetails []entity.OrderDetail
	if err := r.db.Preload("Products").Where("order_id = ?", orderID).Find(&orderDetails).Error; err != nil {
		return nil, err
	}
	return orderDetails, nil
}

func (r *orderDetailRepository) CreateBulkOrderDetails(orderDetails []entity.OrderDetail) error {
	if err := r.db.Create(&orderDetails).Error; err != nil {
		return err
	}
	return nil
}

func (r *orderDetailRepository) GetOrderWithDetailsByID(orderID uint) (*entity.Order, []model.OrderDetailResponse, error) {
	var order entity.Order
	var orderDetails []model.OrderDetailResponse

	err := r.db.Table("orders").
		Select("orders.*, "+
			"order_details.order_id,"+
			"order_details.product_id, "+
			"order_details.quantity, "+
			"order_details.price, "+
			"products.name AS product_name").
		Joins("JOIN order_details ON order_details.order_id = orders.id").
		Joins("JOIN products ON products.id = order_details.product_id").
		Where("orders.id = ?", orderID).
		Find(&orderDetails).Error
	if err != nil {
		return nil, nil, err
	}
	fmt.Println(orderDetails)

	if err := r.db.First(&order, "id = ?", orderID).Error; err != nil {
		return nil, nil, err
	}

	return &order, orderDetails, nil
}
