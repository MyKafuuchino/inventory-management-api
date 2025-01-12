package service

import (
	"inventory-management/entity"
	"inventory-management/repository"
)

type OrderService interface {
	GetAllOrder() ([]entity.Order, error)
	GetOrderById(orderID string) (*entity.Order, error)
	CreateNewOrder(order *entity.Order) (*entity.Order, error)
	UpdateOrderByID(orderID string, order *entity.Order) (*entity.Order, error)
	DeleteOrderById(orderID string) error
}

func (s *orderService) GetAllOrder() ([]entity.Order, error) {
	return s.orderRepository.GetAllOrder()
}

type orderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) OrderService {
	return &orderService{orderRepository: orderRepository}
}

func (s *orderService) GetOrderById(orderID string) (*entity.Order, error) {
	//TODO implement me
	panic("implement me")
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
