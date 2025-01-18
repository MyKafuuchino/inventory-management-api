package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"inventory-management/entity"
	"inventory-management/model"
	"inventory-management/repository"
	"inventory-management/utils"
	"net/http"
)

type OrderService interface {
	GetAllOrders(page, pageSize int) ([]model.GetAllOrdersResponse, int64, int, error)
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

func (s *orderService) GetAllOrders(page, pageSize int) ([]model.GetAllOrdersResponse, int64, int, error) {
	orders, total, totalPages, err := s.orderRepo.GetAllOrders(page, pageSize)
	if err != nil {
		return nil, 0, 0, utils.NewCustomError(http.StatusBadRequest, err.Error())
	}
	if len(orders) == 0 {
		return nil, 0, 0, utils.NewCustomError(http.StatusBadRequest, "no orders found")
	}
	response := make([]model.GetAllOrdersResponse, len(orders))
	for i, order := range orders {
		response[i] = model.GetAllOrdersResponse{
			ID:          order.ID,
			UserID:      order.UserID,
			TotalPrice:  order.TotalPrice,
			OrderStatus: order.OrderStatus,
			CreatedAt:   order.CreatedAt,
		}
	}
	return response, total, totalPages, nil
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
		OrderStatus: order.OrderStatus,
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

	updates := make(map[uint]int)
	for _, od := range newOrderDetail {
		product, exists := productMap[od.ProductID]
		if !exists {
			return nil, utils.NewCustomError(http.StatusNotFound, "Product not found")
		}

		updatedQuantity := product.Stock - od.Quantity
		if updatedQuantity < 0 {
			return nil, utils.NewCustomError(http.StatusBadRequest, fmt.Sprintf("Product %d has insufficient stock", product.ID))
		}

		updates[od.ProductID] = updatedQuantity
	}

	if err := s.productRepo.UpdateProductsQuantities(updates); err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "Failed to update product quantities")
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
