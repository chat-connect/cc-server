package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/game-connect/gc-server/domain/repository"
)

type transactionDao struct {
	Conn *gorm.DB
}

func NewTransactionDao(conn *gorm.DB) repository.TransactionRepository {
	return &transactionDao{
		Conn: conn,
	}
}

func (transactionDao *transactionDao) Begin() (tx *gorm.DB, err error) {
	tx = transactionDao.Conn.Begin()
	if err := tx.Error; err != nil {
		return tx, err
	}

	return tx, err
}

func (transactionDao *transactionDao) Commit(tx *gorm.DB) (err error) {
	tx.Commit()
	if err := tx.Error; err != nil {
		return err
	}

	return err
}

func (transactionDao *transactionDao) Rollback(tx *gorm.DB) (err error) {
	tx.Rollback()
	if err := tx.Error; err != nil {
		return err
	}

	return err
}
