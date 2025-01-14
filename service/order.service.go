package service

import (
	"errors"
	"gorm.io/gorm"
	"inventory-management/entity"
	"inventory-management/model"
	"inventory-management/repository"
	"inventory-management/utils"
	"net/http"
)

type OrderService interface {
	GetAllOrder() ([]entity.Order, error)
	GetOrderById(orderID string) (*entity.Order, error)
	CreateNewOrder(order *model.CreateOrderRequest) (*entity.Order, error)
	UpdateOrderByID(orderID string, body *entity.Order) (*entity.Order, error)
	DeleteOrderById(orderID string) error
}

type orderService struct {
	orderRepository   repository.OrderRepository
	productRepository repository.ProductRepository
}

func NewOrderService(orderRepository repository.OrderRepository, productRepository repository.ProductRepository) OrderService {
	return &orderService{orderRepository: orderRepository, productRepository: productRepository}
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

func (s *orderService) CreateNewOrder(request *model.CreateOrderRequest) (*entity.Order, error) {
	products, err := s.productRepository.GetProductByIDs(request.ProductIDs)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "failed to fetch products")
	}

	if len(products) != len(request.ProductIDs) {
		return nil, utils.NewCustomError(http.StatusBadRequest, "some products are invalid or do not exist")
	}

	var totalPrice int
	var orderDetails []entity.OrderDetail
	for _, product := range products {
		quantity := request.Quantities[product.ID]
		if quantity <= 0 {
			return nil, utils.NewCustomError(http.StatusBadRequest, "invalid quantity for product: "+product.ID)
		}

		totalPrice += product.Price * quantity

		orderDetails = append(orderDetails, entity.OrderDetail{
			ProductID: product.ID,
			Quantity:  quantity,
			Price:     product.Price,
		})
	}

	order := &entity.Order{
		TotalPrice:   totalPrice,
		Status:       request.Status,
		UserID:       request.UserID,
		OrderDetails: orderDetails,
	}

	orderDetail, err := s.orderRepository.CreateNewOrder(order)
	return orderDetail, err
}

func (s *orderService) UpdateOrderByID(orderID string, body *entity.Order) (*entity.Order, error) {
	panic("implement me")
}

func (s *orderService) DeleteOrderById(orderID string) error {
	//TODO implement me
	panic("implement me")
}
