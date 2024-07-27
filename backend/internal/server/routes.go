package server

import "github.com/mattismoel/cookery/internal/handler"

func (s server) addRoutes() {
	s.mux.Put("/recipe", handler.HandlePutRecipe())
	s.mux.Get("/recipe/{id}", handler.GetSingleRecipeById(s.recipeSrv))
}
