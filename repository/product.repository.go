package repository

import (
	"fmt"
	"gorm.io/gorm"
	"inventory-management/entity"
	"time"
)

type ProductRepository interface {
	GetAllProducts(page int, pageSize int) ([]entity.Product, int64, int, error)
	GetProductById(productId uint) (*entity.Product, error)
	CreateNewProduct(product *entity.Product) (*entity.Product, error)
	UpdateProduct(productId uint, reqProduct *entity.Product) error
	DeleteProductById(productId uint) error

	GetProductsByIDs(productIDs []uint) ([]entity.Product, error)
	UpdateProductsQuantities(updates map[uint]int) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetAllProducts(page int, pageSize int) ([]entity.Product, int64, int, error) {
	var products []entity.Product
	var total int64

	if err := r.db.Table("products").Count(&total).Error; err != nil {
		return nil, 0, 0, err
	}

	offset := (page - 1) * pageSize

	if err := r.db.Table("products").
		Limit(pageSize).
		Offset(offset).
		Find(&products).Error; err != nil {
		return nil, 0, 0, err
	}

	totalPages := int((total + int64(pageSize) - 1) / int64(pageSize))

	return products, total, totalPages, nil
}

func (r *productRepository) UpdateProductsQuantities(updates map[uint]int) error {
	tx := r.db.Begin()

	for productID, newQuantity := range updates {
		if err := tx.Table("products").Where("id = ?", productID).Update("stock", newQuantity).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update product %d: %w", productID, err)
		}
	}

	return tx.Commit().Error
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
