package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"inventory-management/entity"
	"inventory-management/model"
	"inventory-management/service"
	"inventory-management/utils"
	"inventory-management/validation"
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
		err = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, utils.NewResponseSuccess[[]entity.Product]("Success get user", products))
}

func (c *ProductController) GetProductById(ctx *gin.Context) {
	userId := ctx.Param("id")
	product, err := c.productService.GetProductById(userId)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, utils.NewResponseSuccess[*entity.Product]("Success get user", product))
}

func (c *ProductController) CreateNewProduct(ctx *gin.Context) {
	productRequest := &model.CreateProductRequest{}
	if err := ctx.ShouldBindJSON(productRequest); err != nil {
		err = ctx.Error(utils.NewCustomError(http.StatusBadRequest, err.Error()))
		return
	}

	if err := validation.ValidationHandler(productRequest); err != nil {
		err = ctx.Error(err)
		return
	}

	createdProduct, err := c.productService.CreateNewProduct(productRequest)
	if err != nil {
		err = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, utils.NewResponseSuccess[*entity.Product]("Success create user", createdProduct))
}

func (c *ProductController) UpdateProduct(ctx *gin.Context) {
	productId := ctx.Param("id")
	product := &entity.Product{}

	if err := ctx.ShouldBindJSON(product); err != nil {
		err = ctx.Error(utils.NewCustomError(http.StatusBadRequest, "Invalid input "+err.Error()))
		return
	}

	if err := validation.ValidationHandler(product); err != nil {
		err = ctx.Error(err)
		return
	}

	updatedProduct, err := c.productService.UpdateProduct(productId, product)
	if err != nil {
		err = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, utils.NewResponseSuccess[*entity.Product]("Success update product", updatedProduct))
}

func (c *ProductController) DeleteProductById(ctx *gin.Context) {
	productId := ctx.Param("id")
	product, err := c.productService.DeleteProduct(productId)
	if err != nil {
		err = ctx.Error(utils.NewCustomError(http.StatusInternalServerError, "Failed to delete product "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewResponseSuccess[*entity.Product]("Success delete product", product))
}
