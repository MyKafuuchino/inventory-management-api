package repository

import (
	"gorm.io/gorm"
	"inventory-management/entity"
)

type UserRepository interface {
	GetAllUsers() ([]entity.User, error)
	GetUserById(id string) (entity.User, error)
	CreateNewUser(body entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAllUsers() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Table("users").Find(&users).Error
	return users, err
}

func (r *userRepository) GetUserById(id string) (entity.User, error) {
	var user entity.User
	err := r.db.Table("users").Where("id = ?", id).First(&user).Error
	return user, err
}

func (r *userRepository) CreateNewUser(body entity.User) error {
	return r.db.Create(&body).Error
}
