package service

import (
	"inventory-management/entity"
	"inventory-management/repository"
)

type ProductService interface {
	GetAllProducts() ([]entity.Product, error)
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &productService{productRepository: productRepository}
}

func (r *productService) GetAllProducts() ([]entity.Product, error) {
	return r.productRepository.GetAllProducts()
}
