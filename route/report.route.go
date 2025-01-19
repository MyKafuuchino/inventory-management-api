package route

import (
	"github.com/gin-gonic/gin"
	"inventory-management/controller"
	"inventory-management/database"
	"inventory-management/middleware"
	"inventory-management/repository"
	"inventory-management/service"
)

func ReportRoute(ctx *gin.RouterGroup) {
	reportRepo := repository.NewReportRepository(database.DB)
	reportService := service.NewReportService(reportRepo)
	reportController := controller.NewReportController(reportService)

	report := ctx.Group("/reports")
	{
		report.GET("/top-products", middleware.ProtectRoute("chaser", "admin"), reportController.GetTopProducts)
		report.GET("/low-stock-products", middleware.ProtectRoute("chaser", "admin"), reportController.GetLowStockProducts)
	}
}
