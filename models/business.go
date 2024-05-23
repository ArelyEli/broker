package models

import (
	"arely.dev/db"
	"arely.dev/schemas"
	"gorm.io/gorm"
)

type Merchant struct {
	gorm.Model
	Name       string
	Commission float64
}

func (m *Merchant) CreateMerchant() error {
	return db.DB.Create(&m).Error
}

func (m *Merchant) UpdateMerchant(request schemas.UpdateMerchantRequest) error {
	if request.Name != "" {
		m.Name = request.Name
	}
	if request.Commission != 0 {
		m.Commission = request.Commission
	}
	return db.DB.Save(&m).Error
}

func (m *Merchant) GetMerchant(id string) error {
	return db.DB.First(&m, "id = ?", id).Error
}

func (m *Merchant) GetEarningsByMerchant(merchant_id string) float64 {
	transaction := Transaction{}
	var totalAmount float64
	db.DB.Model(&transaction).Select("sum(fee)").Where("merchant_id = ?", merchant_id).Scan(&totalAmount)
	return totalAmount
}
