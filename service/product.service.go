package service

import (
	"errors"
	"gorm.io/gorm"
	"inventory-management/entity"
	"inventory-management/model"
	"inventory-management/repository"
	"inventory-management/utils"
	"net/http"
)

type ProductService interface {
	GetAllProducts() ([]entity.Product, error)
	GetProductById(productId string) (*entity.Product, error)
	CreateNewProduct(body *model.CreateProductRequest) (*entity.Product, error)
	UpdateProduct(productId string, body *model.UpdateProductRequest) (*entity.Product, error)
	DeleteProduct(productId string) (*entity.Product, error)
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &productService{productRepository: productRepository}
}

func (s *productService) GetAllProducts() ([]entity.Product, error) {
	products, err := s.productRepository.GetAllProducts()
	if err != nil {
		err = utils.NewCustomError(http.StatusInternalServerError, "Error while fetching products "+err.Error())
	}
	if len(products) == 0 {
		err = utils.NewCustomError(http.StatusNotFound, "Products not found, please create new product")
	}
	return products, err
}

func (s *productService) GetProductById(productId string) (*entity.Product, error) {
	product, err := s.productRepository.GetProductById(productId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = utils.NewCustomError(http.StatusNotFound, "Product not found")
			return nil, err
		}
		err = utils.NewCustomError(http.StatusInternalServerError, "Error while fetching product "+err.Error())
		return nil, err
	}
	return product, err
}

func (s *productService) CreateNewProduct(body *model.CreateProductRequest) (*entity.Product, error) {
	newProduct := &entity.Product{
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
		Stock:       body.Stock,
		LowStock:    body.LowStock,
	}
	productBody, err := s.productRepository.CreateNewProduct(newProduct)
	if err != nil {
		err = utils.NewCustomError(http.StatusInternalServerError, "Error while creating product "+err.Error())
	}
	return productBody, nil
}

func (s *productService) UpdateProduct(productId string, body *model.UpdateProductRequest) (*entity.Product, error) {
	product := &entity.Product{
		Name:        body.Name,
		Description: "",
		Price:       0,
		Stock:       0,
		LowStock:    0,
	}
	updatedProduct, err := s.productRepository.UpdateProduct(productId, product)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = utils.NewCustomError(http.StatusNotFound, "Product not found")
			return nil, err
		}
		err = utils.NewCustomError(http.StatusInternalServerError, "Error while updating product "+err.Error())
		return nil, err
	}
	return updatedProduct, err
}

func (s *productService) DeleteProduct(productId string) (*entity.Product, error) {
	return s.productRepository.DeleteProductById(productId)
}
