package service

import (
	"golang.org/x/crypto/bcrypt"
	"inventory-management/entity"
	"inventory-management/repository"
	"inventory-management/utils"
	"net/http"
)

type UserService interface {
	GetAllUsers() ([]entity.User, error)
	GetUserByID(userID string) (*entity.User, error)
	CreateNewUser(user *entity.User) (*entity.User, error)
	DeleteUserByID(userID string) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {

	return &userService{userRepo: userRepo}
}

func (s *userService) GetAllUsers() ([]entity.User, error) {
	users, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, utils.NewCustomError(http.StatusBadRequest, err.Error())
	}

	if len(users) == 0 {
		return nil, utils.NewCustomError(http.StatusBadRequest, "No users found")
	}
	return users, nil
}

func (s *userService) GetUserByID(userID string) (*entity.User, error) {
	user, err := s.userRepo.GetUserById(userID)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusBadRequest, err.Error())
	}
	return user, nil
}

func (s *userService) CreateNewUser(body *entity.User) (*entity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	body.Password = string(hashedPassword)

	userBody, err := s.userRepo.CreateNewUser(body)
	if err != nil {
		return nil, utils.NewCustomError(http.StatusBadRequest, err.Error())
	}
	return userBody, nil
}

func (s *userService) DeleteUserByID(userID string) error {
	rowsAffected, err := s.userRepo.DeleteUserByID(userID)

	if err != nil {
		return utils.NewCustomError(http.StatusBadRequest, err.Error())
	}

	if rowsAffected == 0 {
		return utils.NewCustomError(http.StatusNotFound, "User not found")
	}
	return nil
}
