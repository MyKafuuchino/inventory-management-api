package repository

import (
	"errors"
	"gorm.io/gorm"
	"inventory-management/entity"
	"time"
)

type ProductRepository interface {
	GetAllProducts() ([]entity.Product, error)
	GetProductById(productId string) (entity.Product, error)
	CreateNewProduct(product *entity.Product) (*entity.Product, error)
	UpdateProduct(productId string, product *entity.Product) (*entity.Product, error)
	DeleteProductById(productId string) (*entity.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetAllProducts() ([]entity.Product, error) {
	var products []entity.Product
	err := r.db.Table("products").Find(&products).Error
	return products, err
}

func (r *productRepository) GetProductById(productId string) (entity.Product, error) {
	var product entity.Product
	err := r.db.Table("products").Where("id = ?", productId).First(&product).Error
	return product, err
}

func (r *productRepository) CreateNewProduct(product *entity.Product) (*entity.Product, error) {
	var err error
	if err := r.db.Table("products").Create(&product).Error; err != nil {
		return nil, err
	}
	return product, err
}

func (r *productRepository) UpdateProduct(productId string, product *entity.Product) (*entity.Product, error) {
	existingProduct := &entity.Product{}
	query := r.db.Table("products").Where("id = ?", productId)
	if err := query.First(existingProduct).Error; err != nil {
		return nil, errors.New("product not found")
	}

	product.UpdatedAt = time.Now()

	if err := query.Updates(product).Error; err != nil {
		return nil, errors.New("failed to update product")
	}

	updatedProduct := &entity.Product{}
	if err := query.First(updatedProduct).Error; err != nil {
		return nil, errors.New("failed to retrieve updated product")
	}
	return updatedProduct, nil
}

func (r *productRepository) DeleteProductById(id string) (*entity.Product, error) {
	product := &entity.Product{}
	if err := r.db.Table("products").Where("id = ?", id).First(product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	if err := r.db.Table("products").Where("id = ?", product.ID).Delete(product).Error; err != nil {
		return nil, errors.New("failed to delete product " + err.Error())
	}
	return product, nil
}
