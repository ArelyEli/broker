package models

import (
	"arely.dev/db"
	"gorm.io/gorm"
)

type Merchant struct {
	gorm.Model
	Name       string
	Commission float64
}

func (m *Merchant) CreateMerchant() {
	db.DB.Create(&m)
}
