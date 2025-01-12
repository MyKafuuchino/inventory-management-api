package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"inventory-management/entity"
	"inventory-management/model"
	"inventory-management/repository"
	"inventory-management/utils"
	"net/http"
)

type AuthService interface {
	Login(body *model.Login) (*entity.User, string, error)
}

type authService struct {
	authRepository repository.AuthRepository
	jwtSecret      []byte
}

func NewAuthService(authRepository repository.AuthRepository, jwtSecret []byte) AuthService {
	return &authService{authRepository: authRepository, jwtSecret: jwtSecret}
}

func (s *authService) Login(body *model.Login) (*entity.User, string, error) {

	user, err := s.authRepository.Login(body.Username)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", utils.NewCustomError(http.StatusUnauthorized, "invalid username or password")
		}
		return nil, "", err
	}

	jwtService := utils.NewJwtService(s.jwtSecret)

	token, err := jwtService.GenJwtToken(user.ID)

	if err != nil {
		return nil, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return nil, "", utils.NewCustomError(http.StatusUnauthorized, "invalid username or password")
	}

	return user, token, nil
}
