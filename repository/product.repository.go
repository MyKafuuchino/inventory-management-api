package repository

import (
	"fmt"
	"gorm.io/gorm"
	"inventory-management/entity"
	"time"
)

type ProductRepository interface {
	GetAllProducts() ([]entity.Product, error)
	GetProductById(productId uint) (*entity.Product, error)
	CreateNewProduct(product *entity.Product) (*entity.Product, error)
	UpdateProduct(productId uint, reqProduct *entity.Product) error
	DeleteProductById(productId uint) error

	GetProductsByIDs(productIDs []uint) ([]entity.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetAllProducts() ([]entity.Product, error) {
	var products []entity.Product
	if err := r.db.Table("products").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) GetProductById(productId uint) (*entity.Product, error) {
	var product *entity.Product
	if err := r.db.Table("products").Where("id = ?", productId).First(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepository) CreateNewProduct(product *entity.Product) (*entity.Product, error) {
	if err := r.db.Table("products").Create(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepository) UpdateProduct(productId uint, reqProduct *entity.Product) error {
	reqProduct.ID = productId
	reqProduct.UpdatedAt = time.Now()

	if err := r.db.Table("products").Save(reqProduct).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepository) DeleteProductById(productID uint) error {
	if err := r.db.Debug().Table("products").Where("id = ?", productID).Delete(&entity.Product{}).Error; err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}
	return nil
}

func (r *productRepository) GetProductsByIDs(productIDs []uint) ([]entity.Product, error) {
	var products []entity.Product
	if len(productIDs) == 0 {
		return nil, fmt.Errorf("product IDs cannot be empty")
	}
	if err := r.db.Where("id IN ?", productIDs).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
