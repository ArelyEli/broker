package main

import (
	"arely.dev/db"
	"arely.dev/models"
	"arely.dev/server"
)

func main() {
	db.Init()
	err := db.DB.AutoMigrate(&models.Business{})
	if err != nil {
		panic(err)
	}

	server.Init()
}
