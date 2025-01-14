package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"inventory-management/entity"
	"inventory-management/model"
	"inventory-management/service"
	"inventory-management/utils"
	"inventory-management/validation"
	"net/http"
)

type OrderController struct {
	orderService service.OrderService
}

func NewOrderController(orderService service.OrderService) *OrderController {
	return &OrderController{orderService}
}

func (c *OrderController) GetAllProducts(ctx *gin.Context) {
	products, err := c.orderService.GetAllOrder()
	if err != nil {
		err = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, utils.NewResponseSuccess("Success get all products", products))
}

func (c *OrderController) GetOrderById(ctx *gin.Context) {
	orderID := ctx.Param("id")
	order, err := c.orderService.GetOrderById(orderID)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, utils.NewResponseSuccess("Success get order by id", order))
}
func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var orderRequest *model.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&orderRequest); err != nil {
		err = ctx.Error(err)
		return
	}

	if err := validation.ValidationHandler(orderRequest); err != nil {
		err = ctx.Error(err)
		return
	}

	user, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	userData, ok := user.(*entity.User)

	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse user data"})
		return
	}

	orderRequest.UserID = userData.ID

	fmt.Println(orderRequest.UserID)

	order, err := c.orderService.CreateNewOrder(orderRequest)
	if err != nil {
		err = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, utils.NewResponseSuccess("Success create order", order))
}
