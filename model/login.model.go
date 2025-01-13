package model

type LoginRequest struct {
	Username string `json:"username" validate:"required,gte=1,lte=255"`
	Password string `json:"password" validate:"required,gte=1,lte=255"`
}
