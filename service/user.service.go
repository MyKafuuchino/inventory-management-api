package service

import (
	"inventory-management/entity"
	"inventory-management/repository"
)

type UserService interface {
	GetAllUsers() ([]entity.User, error)
	GetUserById(id string) (entity.User, error)
	CreateNewUser(user *entity.User) (*entity.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {

	return &userService{userRepo: userRepo}
}

func (s *userService) GetAllUsers() ([]entity.User, error) {
	return s.userRepo.GetAllUsers()
}

func (s *userService) GetUserById(id string) (entity.User, error) {
	return s.userRepo.GetUserById(id)
}

func (s *userService) CreateNewUser(body *entity.User) (*entity.User, error) {
	return s.userRepo.CreateNewUser(body)
}
