package repository

import (
	"errors"
	"gorm.io/gorm"
	"inventory-management/entity"
)

type ProductRepository interface {
	GetAllProducts() ([]entity.Product, error)
	GetProductById(id string) (entity.Product, error)
	CreateNewProduct(product *entity.Product) (*entity.Product, error)
	UpdateProduct(userId string, product *entity.Product) (*entity.Product, error)
	DeleteProductById(id string) (*entity.Product, error)
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

func (r *productRepository) GetProductById(id string) (entity.Product, error) {
	var product entity.Product
	err := r.db.Table("products").Where("id = ?", id).First(&product).Error
	return product, err
}

func (r *productRepository) CreateNewProduct(product *entity.Product) (*entity.Product, error) {
	var err error
	if err := r.db.Table("products").Create(&product).Error; err != nil {
		return nil, err
	}
	return product, err
}

func (r *productRepository) UpdateProduct(userId string, product *entity.Product) (*entity.Product, error) {
	var existingProduct entity.Product
	if err := r.db.Table("products").Where("id = ?", userId).First(&existingProduct).Error; err != nil {
		return nil, errors.New("product not found")
	}
	if err := r.db.Table("products").Where("id = ?", userId).Updates(existingProduct).Error; err != nil {
		return nil, errors.New("failed to update product")
	}
	return &existingProduct, nil
}

func (r *productRepository) DeleteProductById(id string) (*entity.Product, error) {
	product := &entity.Product{}
	if err := r.db.Table("products").First(product, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	if err := r.db.Table("products").Delete(product, id).Error; err != nil {
		return nil, errors.New("failed to delete product " + err.Error())
	}
	return product, nil
}
