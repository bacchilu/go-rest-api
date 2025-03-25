package main

import (
	"github.com/bacchilu/rest-api/db"
	"github.com/bacchilu/rest-api/interactor"
	controller "github.com/bacchilu/rest-api/server"
)

func main() {
	store := db.NewSQLiteEventRepository()
	app := interactor.NewApplication(store)

	server := controller.NewServer(app)
	server.Run()
}
