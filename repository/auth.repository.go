package repository

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"inventory-management/entity"
)

type AuthRepository interface {
	Login(username, password string) (*entity.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r authRepository) Login(username, password string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("username = ?", username).First(&user).Error

	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, err
	}

	return &user, nil
}
