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

type OrderController struct {
	orderService service.OrderService
}

func NewOrderController(orderService service.OrderService) *OrderController {
	return &OrderController{orderService: orderService}
}

func (c *OrderController) GetAllOrders(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))

	orders, total, totalPages, err := c.orderService.GetAllOrders(page, pageSize)
	if err != nil {
		err = ctx.Error(err)
		return
	}

	response := utils.NewPaginatedResponse(
		"Success get all orders",
		orders,
		total,
		totalPages,
		page,
		pageSize,
	)

	ctx.JSON(http.StatusOK, response)
}

func (c *OrderController) GetOrderDetailByID(ctx *gin.Context) {
	orderID := ctx.Param("id")
	order, err := c.orderService.GetOrderDetailById(orderID)
	if err != nil {
		err = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, utils.NewResponseSuccess("Success get order detail", order))
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var err error
	var orderRequest model.CreateOrderRequest

	if err := ctx.ShouldBindJSON(&orderRequest); err != nil {
		err = ctx.Error(utils.NewCustomError(http.StatusBadRequest, err.Error()))
		return
	}

	if err := validation.ValidationHandler(&orderRequest); err != nil {
		err = ctx.Error(err)
		return
	}

	value, exists := ctx.Get("user")
	if !exists {
		err = ctx.Error(utils.NewCustomError(http.StatusBadRequest, "userID not found in context"))
		return
	}

	user := value.(*entity.User)

	orderRequest.UserID = user.ID

	createdOrder, err := c.orderService.CreateOrderWithDetail(&orderRequest)
	if err != nil {
		err = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, utils.NewResponseSuccess("Success create order", createdOrder))
}
