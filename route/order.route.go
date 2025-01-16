package route

import (
	"github.com/gin-gonic/gin"
	"inventory-management/controller"
	"inventory-management/database"
	"inventory-management/middleware"
	"inventory-management/repository"
	"inventory-management/service"
)

func OrderRoute(ctx *gin.RouterGroup) {
	var db = database.DB
	orderRepo := repository.NewOrderRepository(db)
	userRepo := repository.NewUserRepository(db)
	orderDetailRepo := repository.NewOrderDetailRepository(db)
	productRepo := repository.NewProductRepository(db)

	orderService := service.NewOrderService(orderRepo, userRepo, orderDetailRepo, productRepo)
	orderController := controller.NewOrderController(orderService)

	order := ctx.Group("/orders")
	{
		order.GET("", middleware.ProtectRoute("admin"), orderController.GetAllOrders)
		order.POST("", middleware.ProtectRoute("admin"), orderController.CreateOrder)
		order.GET("/:id", orderController.GetOrderDetailByID)
	}
}
