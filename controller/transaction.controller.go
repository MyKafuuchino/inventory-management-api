package controller

import (
	"github.com/gin-gonic/gin"
	"inventory-management/model"
	"inventory-management/service"
	"inventory-management/utils"
	"inventory-management/validation"
	"net/http"
	"strconv"
)

type TransactionController struct {
	transService service.TransactionService
}

func NewTransactionController(tranService service.TransactionService) *TransactionController {
	return &TransactionController{transService: tranService}
}

func (c *TransactionController) UpdateTransactionStatus(ctx *gin.Context) {
	var requestTransaction = &model.UpdateTransactionRequest{}
	orderID := ctx.Param("id")

	iOrderID, err := strconv.Atoi(orderID)

	if err != nil {
		err = ctx.Error(utils.NewCustomError(http.StatusBadRequest, "order_id is not a number"))
		return
	}

	if err := ctx.ShouldBindJSON(requestTransaction); err != nil {
		err = ctx.Error(utils.NewCustomError(http.StatusBadRequest, "Request body invalid."))
		return
	}

	if err := validation.ValidationHandler(requestTransaction); err != nil {
		err = ctx.Error(err)
		return
	}

	order, err := c.transService.UpdateTransaction(iOrderID, requestTransaction)

	ctx.JSON(http.StatusOK, utils.NewResponseSuccess("Success update transaction", order))
}
