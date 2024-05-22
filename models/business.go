package models

import (
	"arely.dev/db"
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
