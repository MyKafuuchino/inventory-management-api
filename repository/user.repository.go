package repository

import (
	"errors"
	"gorm.io/gorm"
	"inventory-management/entity"
)

type UserRepository interface {
	GetAllUsers() ([]entity.User, error)
	GetUserById(id string) (*entity.User, error)
	CreateNewUser(body *entity.User) (*entity.User, error)
	DeleteUserByID(userID string) (int64, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAllUsers() ([]entity.User, error) {
	var users []entity.User
	if err := r.db.Table("users").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetUserById(id string) (*entity.User, error) {
	var user *entity.User
	if err := r.db.Table("users").Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepository) CreateNewUser(body *entity.User) (*entity.User, error) {
	if err := r.db.Table("users").Create(&body).Error; err != nil {
		return nil, err
	}
	return body, nil
}

func (r *userRepository) DeleteUserByID(userID string) (int64, error) {
	result := r.db.Table("users").Where("id = ?", userID).Delete(&entity.User{})
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
