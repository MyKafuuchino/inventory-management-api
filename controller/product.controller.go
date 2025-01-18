package controller

import (
	"github.com/gin-gonic/gin"
	"inventory-management/entity"
	"inventory-management/model"
	"inventory-management/service"
	"inventory-management/utils"
	"inventory-management/validation"
	"net/http"
	"strconv"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (c *ProductController) GetAllProducts(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))

	products, total, totalPages, err := c.productService.GetAllProducts(page, pageSize)
	if err != nil {
		err = ctx.Error(err)
		return
	}

	if len(products) == 0 {
		err = ctx.Error(err)
		return
	}

	response := utils.NewPaginatedResponse(
		"Success fetching products",
		products,
		total,
		totalPages,
		page,
		pageSize,
	)

	ctx.JSON(http.StatusOK, response)
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
	productRequest := &model.UpdateProductRequest{}

	if err := ctx.ShouldBindJSON(productRequest); err != nil {
		err = ctx.Error(utils.NewCustomError(http.StatusBadRequest, "Invalid input "+err.Error()))
		return
	}

	if err := validation.ValidationHandler(productRequest); err != nil {
		err = ctx.Error(err)
		return
	}

	updatedProduct, err := c.productService.UpdateProduct(productId, productRequest)
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
		err = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, utils.NewResponseSuccess[*entity.Product]("Success delete product", product))
}
