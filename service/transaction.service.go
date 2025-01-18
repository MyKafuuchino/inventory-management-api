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

type TransactionService interface {
	UpdateTransaction(transID int, reqTrans *model.UpdateTransactionRequest) (*entity.Order, error)
}

type transactionService struct {
	transRepo repository.TransactionRepository
	orderRepo repository.OrderRepository
}

func NewTransactionService(transRepo repository.TransactionRepository, orderRepo repository.OrderRepository) TransactionService {
	return transactionService{transRepo: transRepo, orderRepo: orderRepo}
}

func (s transactionService) UpdateTransaction(transID int, reqTrans *model.UpdateTransactionRequest) (*entity.Order, error) {
	order, err := s.orderRepo.GetOrderByID(uint(transID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewCustomError(http.StatusNotFound, "Order not found")
		}
		return nil, err
	}

	order.OrderStatus = reqTrans.OrderStatus
	order.Transaction.PaymentMethod = reqTrans.PaymentMethod
	order.Transaction.PaymentStatus = reqTrans.TransactionStatus

	if err := s.orderRepo.UpdateOrderStatus(order); err != nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, err.Error())
	}

	return order, nil
}
