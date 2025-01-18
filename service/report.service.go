package service

import (
	"inventory-management/model"
	"inventory-management/repository"
)

type ReportService interface {
	GetTopProducts(limit int) ([]model.TopProduct, error)
	GetLowStockProducts() ([]model.LowStockProduct, error)
}

type reportService struct {
	reportRepository repository.ReportRepository
}

func NewReportService(reportRepository repository.ReportRepository) ReportService {
	return reportService{reportRepository: reportRepository}
}

func (r reportService) GetTopProducts(limit int) ([]model.TopProduct, error) {
	topProducts, err := r.reportRepository.GetTopProducts(limit)
	if err != nil {
		return nil, err
	}
	return topProducts, nil
}

func (r reportService) GetLowStockProducts() ([]model.LowStockProduct, error) {
	lowStockProducts, err := r.reportRepository.GetLowStockProducts()
	if err != nil {
		return nil, err
	}
	return lowStockProducts, nil
}
