package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/mattismoel/cookery/internal/service"
)

func GetSingleRecipeById(recipeSrv *service.RecipeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			apiError(w, http.StatusBadRequest, "Invalid recipe ID format")
			log.Println(err)
		}

		recipe, err := recipeSrv.ById(ctx, int64(id))
		if err != nil {
			apiError(w, http.StatusBadRequest, "Could not find recipe.")
			log.Println(err)
		}

		err = json.NewEncoder(w).Encode(recipe)
		if err != nil {
			apiError(w, http.StatusInternalServerError, "Something went wrong.")
			log.Println(err)
		}
	}
}
