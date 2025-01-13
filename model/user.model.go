package model

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	FullName string `json:"full_name" validate:"required,max=255"`
	Password string `json:"password" validate:"required,min=3,max=255"`
	Role     string `json:"role,omitempty" validate:"omitempty,oneof=admin chaser customer"`
}
