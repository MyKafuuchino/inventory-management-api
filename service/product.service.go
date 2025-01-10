package service

import (
	"inventory-management/entity"
	"inventory-management/repository"
)

type ProductService interface {
	GetAllProducts() ([]entity.Product, error)
	GetProductById(id string) (entity.Product, error)
	CreateNewProduct(product *entity.Product) (*entity.Product, error)
	UpdateProduct(userId string, product *entity.Product) (*entity.Product, error)
	DeleteProduct(id string) (*entity.Product, error)
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

func (s *productService) GetProductById(id string) (entity.Product, error) {
	return s.productRepository.GetProductById(id)
}

func (s *productService) CreateNewProduct(product *entity.Product) (*entity.Product, error) {
	return s.productRepository.CreateNewProduct(product)
}

func (s *productService) UpdateProduct(userid string, product *entity.Product) (*entity.Product, error) {
	return s.productRepository.UpdateProduct(userid, product)
}

func (s *productService) DeleteProduct(id string) (*entity.Product, error) {
	return s.productRepository.DeleteProductById(id)
}
