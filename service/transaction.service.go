package service

import (
	"inventory-management/repository"
)

type TransactionService interface {
}

type transactionService struct {
	transRepo repository.TransactionRepository
	orderRepo repository.OrderRepository
}

func NewTransactionService(transRepo repository.TransactionRepository, orderRepo repository.OrderRepository) TransactionService {
	return transactionService{transRepo: transRepo, orderRepo: orderRepo}
}
