package main

import (
	"arely.dev/db"
	"arely.dev/models"
	"arely.dev/server"
)

func main() {
	db.Init()
	err := db.DB.AutoMigrate(&models.Merchant{})
	if err != nil {
		panic(err)
	}
	err = db.DB.AutoMigrate(&models.Transaction{})
	if err != nil {
		panic(err)
	}

	server.Init()
}
