package main

import (
	"arely.dev/db"
	"arely.dev/models"
	"arely.dev/server"
)

func main() {
	db.Init()
	db.DB.AutoMigrate(&models.Business{})

	server.Init()
}
