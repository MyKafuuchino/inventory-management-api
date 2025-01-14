package repository

import (
	"errors"
	"gorm.io/gorm"
	"inventory-management/entity"
	"time"
)

type ProductRepository interface {
	GetAllProducts() ([]entity.Product, error)
	GetProductById(productId uint) (*entity.Product, error)
	CreateNewProduct(product *entity.Product) (*entity.Product, error)
	UpdateProduct(productId uint, product *entity.Product) (*entity.Product, error)
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

func (r *productRepository) UpdateProduct(productId uint, product *entity.Product) (*entity.Product, error) {
	product.ID = productId
	product.UpdatedAt = time.Now()

	if err := r.db.Table("products").Save(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (r *productRepository) DeleteProductById(productID uint) error {
	product := &entity.Product{}
	if err := r.db.Table("products").Where("id = ?", productID).Delete(product).Error; err != nil {
		return errors.New("failed to delete product " + err.Error())
	}
	return nil
}

func (r *productRepository) GetProductsByIDs(productIDs []uint) ([]entity.Product, error) {
	var products []entity.Product
	if err := r.db.Where("id IN ?", productIDs).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
