package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/mattismoel/cookery/internal/model"
	"github.com/mattismoel/cookery/internal/service"
)

func HandleAddRecipe(rcpSrv *service.RecipeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		var recipe model.Recipe

		err := json.NewDecoder(r.Body).Decode(&recipe)
		if err != nil {
			apiError(w, http.StatusBadRequest, "could not read request body")
			log.Printf("could not decode request body to JSON: %v\n", err)
			return
		}

		_, err = rcpSrv.Add(ctx, recipe)
		if err != nil {
			apiError(w, http.StatusInternalServerError, "could not add recipe.")
			log.Printf("could not add recipe: %v\n", err)
			return
		}
	}
}
