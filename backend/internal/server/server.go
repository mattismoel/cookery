package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mattismoel/cookery/internal/service"
)

type server struct {
	addr string
	mux  *chi.Mux

	recipeSrv *service.RecipeService
}

// Returns a new server instance.
func New(addr string, recipeSrv *service.RecipeService) *server {
	mux := chi.NewRouter()

	return &server{
		mux:  mux,
		addr: addr,

		recipeSrv: recipeSrv,
	}
}

// Starts the server.
func (s server) Start() error {
	s.addRoutes()
	return http.ListenAndServe(s.addr, s.mux)
}
