package route

import (
	"github.com/gin-gonic/gin"
	"inventory-management/controller"
	"inventory-management/database"
	"inventory-management/middleware"
	"inventory-management/repository"
	"inventory-management/service"
)

func ProductRoute(ctx *gin.RouterGroup) {
	productRepository := repository.NewProductRepository(database.DB)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)
	product := ctx.Group("/products")
	{
		product.GET("", productController.GetAllProducts)
		product.GET("/:id", productController.GetProductById)
		product.POST("", middleware.ProtectRoute("admin"), productController.CreateNewProduct)
		product.PUT("/:id", middleware.ProtectRoute("admin"), productController.UpdateProduct)
		product.DELETE("/:id", middleware.ProtectRoute("admin"), productController.DeleteProductById)
	}
}
