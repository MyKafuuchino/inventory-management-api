package entity

import (
	"time"
)

type User struct {
	Base
	Username  string    `gorm:"size:50;not null;unique" json:"username" validate:"required,min=3,max=50"`
	FullName  string    `gorm:"size:255;not null" json:"full_name" validate:"required,max=255"`
	Password  string    `gorm:"size:255;not null" json:"password" validate:"required,min=3,max=255"`
	Role      string    `gorm:"default:'customer'" json:"role,omitempty" validate:"omitempty,oneof=admin chaser customer"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
