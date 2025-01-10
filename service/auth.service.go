package service

import (
	"inventory-management/entity"
	"inventory-management/repository"
)

type Login struct {
	Username string `validate:"required,gte=1,lte=255"`
	Password string `validate:"required,gte=1,lte=255"`
}

type AuthService interface {
	Login(body *Login) (*entity.User, error)
}

type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(authRepository repository.AuthRepository) AuthService {
	return &authService{authRepository: authRepository}
}

func (s *authService) Login(body *Login) (*entity.User, error) {
	return s.authRepository.Login(body.Username, body.Password)
}
