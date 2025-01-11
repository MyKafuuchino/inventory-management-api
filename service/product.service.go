package service

import (
	"inventory-management/entity"
	"inventory-management/repository"
)

type ProductService interface {
	GetAllProducts() ([]entity.Product, error)
	GetProductById(productId string) (entity.Product, error)
	CreateNewProduct(product *entity.Product) (*entity.Product, error)
	UpdateProduct(productId string, product *entity.Product) (*entity.Product, error)
	DeleteProduct(productId string) (*entity.Product, error)
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &productService{productRepository: productRepository}
}

func (s *productService) GetAllProducts() ([]entity.Product, error) {
	return s.productRepository.GetAllProducts()
}

func (s *productService) GetProductById(productId string) (entity.Product, error) {
	return s.productRepository.GetProductById(productId)
}

func (s *productService) CreateNewProduct(product *entity.Product) (*entity.Product, error) {
	return s.productRepository.CreateNewProduct(product)
}

func (s *productService) UpdateProduct(productId string, product *entity.Product) (*entity.Product, error) {
	return s.productRepository.UpdateProduct(productId, product)
}

func (s *productService) DeleteProduct(productId string) (*entity.Product, error) {
	return s.productRepository.DeleteProductById(productId)
}
