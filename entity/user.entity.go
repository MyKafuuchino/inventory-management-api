package entity

import (
	"time"
)

type User struct {
	Base
	Username  string    `json:"username" validate:"required,min=3,max=50"`
	FullName  string    `json:"full_name" validate:"required,max=255"`
	Password  string    `json:"password" validate:"required,min=8"`
	Role      string    `gorm:"default:'customer'" json:"role,omitempty" validate:"omitempty,oneof=admin chaser customer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
