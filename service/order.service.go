package service

import (
	"errors"
	"gorm.io/gorm"
	"inventory-management/entity"
	"inventory-management/repository"
	"inventory-management/utils"
	"net/http"
)

type OrderService interface {
	GetAllOrder() ([]entity.Order, error)
	GetOrderById(orderID string) (*entity.Order, error)
	CreateNewOrder(order *entity.Order) (*entity.Order, error)
	UpdateOrderByID(orderID string, order *entity.Order) (*entity.Order, error)
	DeleteOrderById(orderID string) error
}

type orderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) OrderService {
	return &orderService{orderRepository: orderRepository}
}

func (s *orderService) GetAllOrder() ([]entity.Order, error) {
	orders, err := s.orderRepository.GetAllOrder()
	if len(orders) == 0 {
		return nil, utils.NewCustomError(http.StatusNotFound, "no orders found, create new order")
	}
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func (s *orderService) GetOrderById(orderID string) (*entity.Order, error) {
	order, err := s.orderRepository.GetOrderById(orderID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := utils.NewCustomError(http.StatusNotFound, "order not found")
			return nil, err
		}
		err := utils.NewCustomError(http.StatusInternalServerError, "failed to get order")
		return nil, err
	}
	return order, nil
}

func (s *orderService) CreateNewOrder(order *entity.Order) (*entity.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (s *orderService) UpdateOrderByID(orderID string, order *entity.Order) (*entity.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (s *orderService) DeleteOrderById(orderID string) error {
	//TODO implement me
	panic("implement me")
}
