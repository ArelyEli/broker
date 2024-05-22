package main

import (
	"arely.dev/db"
	"arely.dev/server"
)

func main() {
	db.Init()
	server.Init()
}
