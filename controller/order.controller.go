package controller

import (
	"github.com/gin-gonic/gin"
	"inventory-management/service"
	"inventory-management/utils"
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
		err = ctx.Error(utils.NewCustomError(400, "Failed to get all products "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, utils.NewResponseSuccess("Success get all products", products))
}
