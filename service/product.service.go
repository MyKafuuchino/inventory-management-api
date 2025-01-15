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
	u, err := utils.ParseStringToUint(productId)
	if err != nil {
		err = utils.NewCustomError(http.StatusInternalServerError, err.Error())
	}
	product, err := s.productRepository.GetProductById(u)
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
	u, err := utils.ParseStringToUint(productId)
	if err != nil {
		err = utils.NewCustomError(http.StatusInternalServerError, err.Error())
	}
	existingProduct, err := s.productRepository.GetProductById(u)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = utils.NewCustomError(http.StatusNotFound, "Product not found")
			return nil, err
		}
		err = utils.NewCustomError(http.StatusInternalServerError, "Error while fetching product "+err.Error())
		return nil, err
	}

	utils.MapUpdateField(existingProduct, body)

	u, err = utils.ParseStringToUint(productId)

	if err != nil {
		err = utils.NewCustomError(http.StatusInternalServerError, err.Error())
	}

	updatedProduct, err := s.productRepository.UpdateProduct(u, existingProduct)

	if err != nil {
		err = utils.NewCustomError(http.StatusInternalServerError, "failed to update product")
		return nil, err
	}
	return updatedProduct, err
}

func (s *productService) DeleteProduct(productId string) (*entity.Product, error) {
	var err error
	product := &entity.Product{}
	u, err := utils.ParseStringToUint(productId)

	if err != nil {
		err = utils.NewCustomError(http.StatusInternalServerError, err.Error())
	}

	if product, err = s.productRepository.GetProductById(u); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = utils.NewCustomError(http.StatusNotFound, "Product not found")
			return nil, err
		}
		err = utils.NewCustomError(http.StatusInternalServerError, "Error while fetching product "+err.Error())
		return nil, err
	}

	if err := s.productRepository.DeleteProductById(u); err != nil {
		err = utils.NewCustomError(http.StatusInternalServerError, "Error while deleting product "+err.Error())
		return nil, err
	}

	return product, nil
}
