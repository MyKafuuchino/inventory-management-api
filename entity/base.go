package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID string `json:"id"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return
}
