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
	GetAllOrders() ([]entity.Order, error)
	GetOrderDetailById(orderID string) (*model.OrderResponse, error)
	CreateOrderWithDetail(reqOrder *model.CreateOrderRequest) (*model.OrderResponse, error)
}

type orderService struct {
	orderRepo       repository.OrderRepository
	userRepo        repository.UserRepository
	orderDetailRepo repository.OrderDetailRepository
	productRepo     repository.ProductRepository
	transactionRepo repository.TransactionRepository
}

func NewOrderService(orderRepo repository.OrderRepository, userRepo repository.UserRepository, orderDetailRepo repository.OrderDetailRepository, productRepo repository.ProductRepository, transactionRepo repository.TransactionRepository) OrderService {
	return &orderService{orderRepo: orderRepo, userRepo: userRepo, orderDetailRepo: orderDetailRepo, productRepo: productRepo, transactionRepo: transactionRepo}
}

func (s *orderService) GetAllOrders() ([]entity.Order, error) {
	orders, err := s.orderRepo.GetAllOrders()
	if err != nil {
		return nil, utils.NewCustomError(http.StatusBadRequest, err.Error())
	}
	if len(orders) == 0 {
		return nil, utils.NewCustomError(http.StatusBadRequest, "no orders found")
	}
	return orders, nil
}

func (s *orderService) GetOrderDetailById(orderId string) (*model.OrderResponse, error) {
	uOrderID, err := utils.ParseStringToUint(orderId)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to parse order id")
	}

	order, orderDetails, err := s.orderDetailRepo.GetOrderWithDetailsByID(uOrderID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewCustomError(http.StatusNotFound, "Order not found")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to get order details")
	}

	var orderDetailResponse []model.OrderDetailResponse
	for _, od := range orderDetails {
		orderDetailResponse = append(orderDetailResponse, model.OrderDetailResponse{
			ProductName: od.ProductName,
			ProductID:   od.ProductID,
			Quantity:    od.Quantity,
			Price:       od.Price,
		})
	}

	response := &model.OrderResponse{
		ID:          order.ID,
		UserID:      order.UserID,
		TotalPrice:  order.TotalPrice,
		OrderDetail: orderDetailResponse,
	}

	return response, nil
}

func (s *orderService) CreateOrderWithDetail(reqOrder *model.CreateOrderRequest) (*model.OrderResponse, error) {
	if _, err := s.userRepo.GetUserById(reqOrder.UserID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewCustomError(http.StatusNotFound, "User not found")
		}
		return nil, utils.NewCustomError(http.StatusBadRequest, "Failed to validate user")
	}

	var totalPrice int
	var newOrderDetail []entity.OrderDetail

	productIDs := make([]uint, len(reqOrder.OrderDetails))
	for i, od := range reqOrder.OrderDetails {
		productIDs[i] = od.ProductID
	}

	products, err := s.productRepo.GetProductsByIDs(productIDs)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to get products")
	}

	productMap := make(map[uint]*entity.Product)
	for _, product := range products {
		productMap[product.ID] = &product
	}

	for _, od := range reqOrder.OrderDetails {
		product, exists := productMap[od.ProductID]
		if !exists {
			return nil, utils.NewCustomError(http.StatusNotFound, "Product not found")
		}

		odPrice := od.Quantity * product.Price
		totalPrice += odPrice

		newOrderDetail = append(newOrderDetail, entity.OrderDetail{
			ProductID: od.ProductID,
			Quantity:  od.Quantity,
			Price:     odPrice,
			History:   model.History{},
		})
	}

	newOrder := &entity.Order{
		ID:          0,
		UserID:      reqOrder.UserID,
		OrderStatus: reqOrder.OrderStatus,
		TotalPrice:  totalPrice,
	}

	err = s.orderRepo.CreateOrderWithDetail(newOrder, newOrderDetail)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to create order")
	}

	orderDetailResponse := make([]model.OrderDetailResponse, len(newOrderDetail))
	for i, od := range newOrderDetail {
		for _, p := range products {
			orderDetailResponse[i] = model.OrderDetailResponse{
				ProductName: p.Name,
				ProductID:   od.ProductID,
				Quantity:    od.Quantity,
				Price:       od.Price,
			}
			break
		}
	}

	var createTransaction = &entity.Transaction{
		OrderID:       newOrder.ID,
		UserID:        reqOrder.UserID,
		PaymentMethod: "",
	}

	if err := s.transactionRepo.CreateTransaction(createTransaction); err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to create transaction "+err.Error())
	}

	orderResponse := &model.OrderResponse{
		ID:          newOrder.ID,
		UserID:      newOrder.UserID,
		TotalPrice:  newOrder.TotalPrice,
		OrderStatus: newOrder.OrderStatus,
		OrderDetail: orderDetailResponse,
	}

	return orderResponse, nil
}
