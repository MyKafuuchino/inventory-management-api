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
	GetOrderDetailById(orderID string) (*model.OrderResponse, error)
	CreateNewOrder(order *model.CreateOrderRequest) (*model.OrderResponse, error)
	UpdateOrderService(orderID string, status *model.UpdateOrderStatusRequest) (*entity.Order, error)
}

type orderService struct {
	orderRepository   repository.OrderRepository
	productRepository repository.ProductRepository
	orderDetailRepo   repository.OrderDetailRepository
}

func NewOrderService(orderRepository repository.OrderRepository, productRepository repository.ProductRepository, orderDetailRepo repository.OrderDetailRepository) OrderService {
	return &orderService{orderRepository: orderRepository, productRepository: productRepository, orderDetailRepo: orderDetailRepo}
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

func (s *orderService) GetOrderDetailById(orderID string) (*model.OrderResponse, error) {
	order, err := s.orderDetailRepo.GetOrderDetailByID(orderID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := utils.NewCustomError(http.StatusNotFound, "order not found")
			return nil, err
		}
		err := utils.NewCustomError(http.StatusInternalServerError, err.Error())
		return nil, err
	}

	response := &model.OrderResponse{
		ID:         order.ID,
		UserID:     order.UserID,
		TotalPrice: order.TotalPrice,
		Status:     order.Status,
	}

	for _, detail := range order.OrderDetails {
		response.OrderDetails = append(response.OrderDetails, model.OrderDetailResponse{
			OrderID:   detail.OrderID,
			ProductID: detail.ProductID,
			Quantity:  detail.Quantity,
			Price:     detail.Price,
			Product:   detail.Product,
		})
	}

	return response, nil
}

func (s *orderService) CreateNewOrder(request *model.CreateOrderRequest) (*model.OrderResponse, error) {
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
	orderDetail := &entity.Order{}
	orderDetail, err = s.orderRepository.CreateNewOrder(order)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "failed to create order")
	}

	response := &model.OrderResponse{
		ID: orderDetail.ID,
	}

	for _, detail := range orderDetails {
		response.OrderDetails = append(response.OrderDetails, model.OrderDetailResponse{
			OrderID:   detail.OrderID,
			ProductID: detail.ProductID,
			Quantity:  detail.Quantity,
			Price:     detail.Price,
			Product:   detail.Product,
		})
	}
	return response, nil
}

func (s *orderService) UpdateOrderService(orderID string, orderRequest *model.UpdateOrderStatusRequest) (*entity.Order, error) {
	var err error
	var orderExist *entity.Order

	if orderExist, err = s.orderRepository.GetOrderById(orderID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewCustomError(http.StatusNotFound, "order not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "failed to fetch order")
	}

	orderExist.Status = orderRequest.Status

	var order *entity.Order
	if order, err = s.orderRepository.UpdateOrderStatus(orderExist); err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "failed to update order")
	}
	return order, nil
}
