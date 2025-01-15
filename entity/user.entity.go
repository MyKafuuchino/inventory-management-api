package entity

import (
	"inventory-management/model"
)

type User struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	Username string `gorm:"size:50;not null;unique" json:"username"`
	FullName string `gorm:"size:255;not null" json:"full_name"`
	Password string `gorm:"size:255;not null" json:"password"`
	Role     string `gorm:"default:'customer'" json:"role,omitempty"`
	model.History
}
