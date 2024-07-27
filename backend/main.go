package main

import (
	"log"

	"github.com/mattismoel/cookery/internal/server"
	"github.com/mattismoel/cookery/internal/service"
	"github.com/mattismoel/cookery/internal/storage/memory"
)

func main() {
	storage := memory.NewMemoryStorage()

	usrSrv := service.NewUserService(storage)
	recipeSrv := service.NewRecipeService(storage)

	srv := server.New(":8080", usrSrv, recipeSrv)

	err := srv.Start()
	if err != nil {
		log.Fatal(err)
	}
}
