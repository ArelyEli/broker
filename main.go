package main

import (
	"flag"
	"fmt"

	"arely.dev/db"
	"arely.dev/models"
	"arely.dev/server"
)

func runMigrations() {
	err := db.DB.AutoMigrate(&models.Merchant{}, &models.Transaction{})
	if err != nil {
		panic(err)
	}
}

func main() {
	db.Init()
	runMigrations()

	migrateFlag := flag.Bool("migrate", false, "Run database migrations")
	flag.Parse()

	if *migrateFlag {
		fmt.Println("Migrations executed successfully.")
		return
	}

	server.Init()
}
