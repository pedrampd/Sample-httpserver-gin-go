package main

import (
	initproject "Desktop/go-crud/InitProject"
	"Desktop/go-crud/api"
	"log"
)

func main() {
	db, err := initproject.InitDatabase()
	if err != nil {
		log.Fatal("kir shodim")
		panic("db not connected")
	}
	server := initproject.InitServer()
	api.SetBookRoutes(server, db)
	server.Run(":8080")

}
