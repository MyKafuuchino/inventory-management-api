package route

import (
	"github.com/gin-gonic/gin"
	"inventory-management/controller"
	"inventory-management/database"
	"inventory-management/repository"
	"inventory-management/service"
)

func ReportRoute(ctx *gin.RouterGroup) {
	reportRepo := repository.NewReportRepository(database.DB)
	reportService := service.NewReportService(reportRepo)
	reportController := controller.NewReportController(reportService)

	report := ctx.Group("/reports")
	{
		report.GET("/top-products", reportController.GetTopProducts)
		report.GET("/low-stock-products", reportController.GetLowStockProducts)
	}
}
