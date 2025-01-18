package route

import (
	"github.com/gin-gonic/gin"
	"inventory-management/controller"
	"inventory-management/database"
	"inventory-management/repository"
	"inventory-management/service"
)

func TransactionRoute(ctx *gin.RouterGroup) {
	db := database.DB
	transRepo := repository.NewTransactionRepository(db)
	orderRepo := repository.NewOrderRepository(db)

	transService := service.NewTransactionService(transRepo, orderRepo)
	transController := controller.NewTransactionController(transService)

	transaction := ctx.Group("/transactions")
	{
		transaction.PUT("/:id", transController.UpdateTransactionStatus)
	}
}
