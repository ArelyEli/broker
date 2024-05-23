package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func Init() {
	host := getEnv("DB_HOST", "db")
	user := getEnv("DB_USER", "arely")
	password := getEnv("DB_PASSWORD", "password")
	dbname := getEnv("DB_NAME", "payments")
	port := getEnv("DB_PORT", "5432")
	sslmode := getEnv("DB_SSLMODE", "disable")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db
}

func ClearDB() {
	DB.Exec("DELETE FROM merchants")

	DB.Exec("DELETE FROM transactions")
}
