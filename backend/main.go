package main

import (
	"log"

	"github.com/mattismoel/cookery/internal/server"
	"github.com/mattismoel/cookery/internal/service"
	"github.com/mattismoel/cookery/internal/storage/sqlite"
)

func main() {
	storage := sqlite.New("database.db")
	err := storage.Start()
	if err != nil {
		log.Fatal(err)
	}

	usrSrv := service.NewUserService(storage)
	recipeSrv := service.NewRecipeService(storage)

	srv := server.New(":8080", usrSrv, recipeSrv)

	err = srv.Start()
	if err != nil {
		log.Fatal(err)
	}
}
