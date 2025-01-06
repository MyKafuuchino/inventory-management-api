package entity

import "time"

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username" validate:"required,min=3,max=50"`
	FullName  string    `json:"fullName" validate:"required,min=3,max=255"`
	Password  string    `json:"password" validate:"required,min=8"`
	Role      string    `json:"role" validate:"required,oneof=admin chaser customer"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
