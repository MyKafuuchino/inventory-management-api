package controller

import (
	"inventory-management/service"
)

type TransactionController struct {
	transService service.TransactionService
}

func NewTransactionController(tranService service.TransactionService) *TransactionController {
	return &TransactionController{transService: tranService}
}
