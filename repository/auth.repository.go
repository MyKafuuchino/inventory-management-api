package repository

import (
	"gorm.io/gorm"
	"inventory-management/entity"
)

type AuthRepository interface {
	Login(username string) (*entity.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) Login(username string) (*entity.User, error) {

	var user entity.User
	err := r.db.Where("username = ?", username).First(&user).Error

	return &user, err
}
