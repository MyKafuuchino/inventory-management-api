package entity

import (
	"time"
)

type User struct {
	Base
	Username  string    `gorm:"size:50;not null;unique" json:"username"`
	FullName  string    `gorm:"size:255;not null" json:"full_name"`
	Password  string    `gorm:"size:255;not null" json:"password"`
	Role      string    `gorm:"default:'customer'" json:"role,omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Orders []Order `gorm:"foreignKey:UserID" json:"orders,omitempty"`
}
