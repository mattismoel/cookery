package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/mattismoel/cookery/internal/service"
)

func GetSingleRecipeById(recipeSrv *service.RecipeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello", r.URL)
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			log.Fatal(err)
		}

		recipe, err := recipeSrv.ById(int64(id))
		if err != nil {
			log.Fatal(err)
		}

		err = json.NewEncoder(w).Encode(recipe)
		if err != nil {
			log.Fatal(err)
		}
	}
}
