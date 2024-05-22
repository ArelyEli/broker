package db

import (
	"arely.dev/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() {
	dsn := "host=db user=arely password=password dbname=payments port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Merchant{})
}
