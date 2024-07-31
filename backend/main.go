package main

import (
	"log"

	"github.com/mattismoel/cookery/internal/filemanager/s3fm"
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

	fileManager, err := s3fm.New("eu-north-1", "cookery-application")
	if err != nil {
		log.Fatal(err)
	}

	usrSrv := service.NewUserService(storage)
	recipeSrv := service.NewRecipeService(storage, fileManager)

	srv := server.New(":8080", usrSrv, recipeSrv)

	err = srv.Start()
	if err != nil {
		log.Fatal(err)
	}
}
