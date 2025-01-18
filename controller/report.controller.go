package controller

import (
	"github.com/gin-gonic/gin"
	"inventory-management/service"
	"inventory-management/utils"
	"strconv"
)

type ReportController struct {
	reportService service.ReportService
}

func NewReportController(reportService service.ReportService) *ReportController {
	return &ReportController{reportService: reportService}
}

func (c *ReportController) GetTopProducts(ctx *gin.Context) {
	limit := 10
	if l := ctx.Query("limit"); l != "" {
		if val, err := strconv.Atoi(l); err == nil {
			limit = val
		}
	}

	topProducts, err := c.reportService.GetTopProducts(limit)
	if err != nil {
		ctx.JSON(500, gin.H{"success": false, "message": "Failed to get top products", "error": err.Error()})
		return
	}

	ctx.JSON(200, utils.NewResponseSuccess("Top products fetched successfully", topProducts))
}

func (c *ReportController) GetLowStockProducts(ctx *gin.Context) {
	lowStockProducts, err := c.reportService.GetLowStockProducts()
	if err != nil {
		ctx.JSON(500, gin.H{"success": false, "message": "Failed to get low stock products", "error": err.Error()})
		return
	}

	ctx.JSON(200, utils.NewResponseSuccess("Low stock products fetched successfully", lowStockProducts))
}
