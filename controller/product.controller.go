package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"inventory-management/entity"
	"inventory-management/service"
	"inventory-management/utils"
	"net/http"
)

var validate = validator.New()

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

func (c *ProductController) GetProductById(ctx *gin.Context) {
	userId := ctx.Param("id")
	product, err := c.productService.GetProductById(userId)
	if err != nil {
		err = ctx.Error(entity.NewCustomError(http.StatusNotFound, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, entity.NewResponseSuccess[entity.Product]("Success get user", product))
}

func (c *ProductController) CreateNewProduct(ctx *gin.Context) {
	product := entity.Product{}
	if err := ctx.ShouldBindJSON(&product); err != nil {
		err = ctx.Error(entity.NewCustomError(http.StatusBadRequest, err.Error()))
		return
	}
	if err := validate.Struct(&product); err != nil {
		err = ctx.Error(entity.NewCustomError(http.StatusBadRequest, "Validation Failed", utils.GetErrorValidationMessages(err)...))
		return
	}
	createdProduct, err := c.productService.CreateNewProduct(&product)
	if err != nil {
		err = ctx.Error(entity.NewCustomError(http.StatusInternalServerError, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, entity.NewResponseSuccess[*entity.Product]("Success create user", createdProduct))
}

func (c *ProductController) UpdateProduct(ctx *gin.Context) {
	userId := ctx.Param("id")
	product := &entity.Product{}
	if err := ctx.ShouldBindJSON(product); err != nil {
		err = ctx.Error(entity.NewCustomError(http.StatusBadRequest, err.Error()))
		return
	}

	fmt.Printf("Request ID: %s, Product: %+v\n", userId, product)

	//if err := validate.Struct(&product); err != nil {
	//	err = ctx.Error(entity.NewCustomError(http.StatusBadRequest, err.Error()))
	//	return
	//}
	updatedProduct, err := c.productService.UpdateProduct(userId, product)
	if err != nil {
		err = ctx.Error(entity.NewCustomError(http.StatusInternalServerError, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, entity.NewResponseSuccess[*entity.Product]("Success update user", updatedProduct))
}
