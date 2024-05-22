package models

import "gorm.io/gorm"

type Merchant struct {
	gorm.Model
	Name       string
	Commission float64
}
