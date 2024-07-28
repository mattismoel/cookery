package server

import "github.com/mattismoel/cookery/internal/handler"

func (s server) addRoutes() {
	s.mux.Put("/recipe", handler.HandlePutRecipe())
	s.mux.Put("/register", handler.HandleRegister(s.usrSrv))

	s.mux.Get("/recipe/{id}", handler.GetSingleRecipeById(s.recipeSrv))
	s.mux.Post("/login", handler.HandleLogin(s.usrSrv))
}
