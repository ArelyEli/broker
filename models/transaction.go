package models

import (
	"arely.dev/db"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Amount     float64
	Fee        float64
	BusinessID uint
	Business   Business `gorm:"constraint:OnDelete:CASCADE;"`
}

func (t *Transaction) GetTransaction(id string) error {
	return db.DB.First(&t, "id = ?", id).Error
}

func (t *Transaction) CreateTransaction() error {
	return db.DB.Create(&t).Error
}

func (t *Transaction) GetEarnings() float64 {
	var totalAmount float64
	db.DB.Model(&t).Select("sum(fee)").Scan(&totalAmount)
	return totalAmount
}
