package entity

import (
	"time"
)

type User struct {
	Base
	Username  string    `json:"username"`
	FullName  string    `json:"full_name"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time ` json:"updated_at"`
}
