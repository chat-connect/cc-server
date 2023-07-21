package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/chat-connect/cc-server/domain/repository"
)

type transactionRepository struct {
	Conn *gorm.DB
}

func NewTransactionRepository(conn *gorm.DB) repository.TransactionRepository {
	return &transactionRepository{
		Conn: conn,
	}
}

func (transactionRepository *transactionRepository) Begin() (tx *gorm.DB, err error) {
	tx = transactionRepository.Conn.Begin()
	if err := tx.Error; err != nil {
		return tx, err
	}

	return tx, err
}

func (transactionRepository *transactionRepository) Commit(tx *gorm.DB) (err error) {
	tx.Commit()
	if err := tx.Error; err != nil {
		return err
	}

	return err
}

func (transactionRepository *transactionRepository) Rollback(tx *gorm.DB) (err error) {
	tx.Rollback()
	if err := tx.Error; err != nil {
		return err
	}

	return err
}