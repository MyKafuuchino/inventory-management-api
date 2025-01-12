package model

type Login struct {
	Username string `validate:"required,gte=1,lte=255"`
	Password string `validate:"required,gte=1,lte=255"`
}
