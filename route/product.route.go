package route

import (
	"github.com/gin-gonic/gin"
	"inventory-management/controller"
	"inventory-management/database"
	"inventory-management/repository"
	"inventory-management/service"
)

func ProductRoute(ctx *gin.RouterGroup) {
	productRepository := repository.NewProductRepository(database.DB)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)
	product := ctx.Group("/products")
	{
		product.GET("/", productController.GetAllProducts)
	}
}
