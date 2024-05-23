package models

import (
	"arely.dev/db"
	"arely.dev/schemas"
	"gorm.io/gorm"
)

type Business struct {
	gorm.Model
	Name       string
	Commission float64
}

func (m *Business) CreateBusiness() error {
	return db.DB.Create(&m).Error
}

func (m *Business) UpdateBusiness(request schemas.UpdateBusinessRequest) error {
	if request.Name != "" {
		m.Name = request.Name
	}
	if request.Commission != 0 {
		m.Commission = request.Commission
	}
	return db.DB.Save(&m).Error
}

func (m *Business) GetBusiness(id string) error {
	return db.DB.First(&m, "id = ?", id).Error
}

func (m *Business) GetEarningsByBusiness(business_id string) float64 {
	transaction := Transaction{}
	var totalAmount float64
	db.DB.Model(&transaction).Select("sum(fee)").Where("business_id = ?", business_id).Scan(&totalAmount)
	return totalAmount
}
