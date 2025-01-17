package repository

import (
	"gorm.io/gorm"
	"inventory-management/entity"
	"time"
)

type TransactionRepository interface {
	CreateTransaction(reqTrans *entity.Transaction) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return transactionRepository{db: db}

}
func (t transactionRepository) CreateTransaction(reqTrans *entity.Transaction) error {
	reqTrans.TransactionAt = time.Now()
	if err := t.db.Create(&reqTrans).Error; err != nil {
		return err
	}
	return nil
}
