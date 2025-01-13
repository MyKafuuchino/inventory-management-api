package route

import (
	"github.com/gin-gonic/gin"
	"inventory-management/controller"
	"inventory-management/database"
	"inventory-management/repository"
	"inventory-management/service"
)

func OrderRoute(ctx *gin.RouterGroup) {
	orderRepository := repository.NewOrderRepository(database.DB)
	orderService := service.NewOrderService(orderRepository)
	orderController := controller.NewOrderController(orderService)

	order := ctx.Group("/orders")
	{
		order.GET("", orderController.GetAllProducts)
		order.GET("/:id", orderController.GetOrderById)
	}
}
