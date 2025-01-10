package repository

import (
	"gorm.io/gorm"
	"inventory-management/entity"
)

type ProductRepository interface {
	GetAllProducts() ([]entity.Product, error)
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
