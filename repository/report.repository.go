package repository

import (
	"gorm.io/gorm"
	"inventory-management/model"
)

type ReportRepository interface {
	GetTopProducts(limit int) ([]model.TopProduct, error)
	GetLowStockProducts() ([]model.LowStockProduct, error)
}

type reportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return &reportRepository{db: db}
}

func (r *reportRepository) GetTopProducts(limit int) ([]model.TopProduct, error) {
	var topProducts []model.TopProduct
	err := r.db.Table("products").
		Select("products.id AS product_id, products.name AS product_name, SUM(order_details.quantity) AS total_sold").
		Joins("JOIN order_details ON products.id = order_details.product_id").
		Group("products.id, products.name").
		Order("total_sold DESC").
		Limit(limit).
		Scan(&topProducts).Error
	return topProducts, err
}
func (r *reportRepository) GetLowStockProducts() ([]model.LowStockProduct, error) {
	var lowStockProducts []model.LowStockProduct
	err := r.db.Table("products").
		Select("id AS product_id, name AS product_name, stock, low_stock").
		Where("stock <= low_stock").
		Order("stock ASC").
		Scan(&lowStockProducts).Error
	return lowStockProducts, err
}
