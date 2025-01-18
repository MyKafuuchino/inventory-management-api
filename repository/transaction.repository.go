package repository

import (
	"gorm.io/gorm"
	"inventory-management/entity"
)

type TransactionRepository interface {
	CreateTransaction(reqTrans *entity.Transaction) error
	GetTransactionById(id uint) (*entity.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return transactionRepository{db: db}

}

func (r transactionRepository) CreateTransaction(reqTrans *entity.Transaction) error {
	if err := r.db.Create(reqTrans).Error; err != nil {
		return err
	}
	return nil
}

func (r transactionRepository) GetTransactionById(id uint) (*entity.Transaction, error) {
	transaction := &entity.Transaction{}
	if err := r.db.First(transaction, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}
