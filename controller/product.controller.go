package controller

import (
	"github.com/gin-gonic/gin"
	"inventory-management/entity"
	"inventory-management/service"
	"net/http"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (c *ProductController) GetAllProducts(ctx *gin.Context) {
	products, err := c.productService.GetAllProducts()
	if err != nil {
		err = ctx.Error(entity.NewCustomError(http.StatusInternalServerError, err.Error()))
		return
	}
	if len(products) == 0 {
		err = ctx.Error(entity.NewCustomError(http.StatusNotFound, "No product found, please create new product"))
		return
	}
	ctx.JSON(http.StatusOK, entity.NewResponseSuccess[[]entity.Product]("Success get user", products))
}
