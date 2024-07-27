package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mattismoel/cookery/internal/service"
)

type server struct {
	addr string
	mux  *chi.Mux

	usrSrv    *service.UserService
	recipeSrv *service.RecipeService
}

// Returns a new server instance.
func New(addr string, usrSrv *service.UserService, recipeSrv *service.RecipeService) *server {
	mux := chi.NewRouter()

	return &server{
		mux:  mux,
		addr: addr,

		usrSrv:    usrSrv,
		recipeSrv: recipeSrv,
	}
}

// Starts the server.
func (s server) Start() error {
	s.addRoutes()
	return http.ListenAndServe(s.addr, s.mux)
}
