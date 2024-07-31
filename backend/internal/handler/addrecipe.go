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
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			apiError(w, http.StatusBadRequest, "could not parse form data")
			log.Printf("could not parse multipart form: %v\n", err)
			return
		}

		var rcp model.Recipe

		rcpJSONStr := r.FormValue("recipe")
		if rcpJSONStr == "" {
			apiError(w, http.StatusBadRequest, "could not get recipe JSON")
			log.Printf("could not read form recipe\n")
			return
		}

		err = json.Unmarshal([]byte(rcpJSONStr), &rcp)
		if err != nil {
			apiError(w, http.StatusBadRequest, "could not read recipe JSON")
			log.Printf("could not unmarshal recipe JSON: %v\n", err)
			return
		}

		_, fh, err := r.FormFile("bannerFile")
		if err != nil {
			apiError(w, http.StatusInternalServerError, "Something went wrong reading banner file.")
			log.Printf("could not get form file: %v\n", err)
			return
		}

		rcpId, err := rcpSrv.Add(ctx, rcp)
		if err != nil {
			apiError(w, http.StatusInternalServerError, "Something went wrong adding recipe.")
			log.Printf("could not add recipe: %v\n", err)
			return
		}

		_, err = rcpSrv.SetBannerImage(ctx, rcpId, fh)
		if err != nil {
			apiError(w, http.StatusInternalServerError, "Something went wrong adding banner image.")
			log.Printf("could not add banner image: %v", err)
			return
		}
	}
}

// func requestToRecipe(r *http.Request) (model.Recipe, error) {
// 	var rcp model.Recipe
//
// 	var title string
// 	var cookMinutes, totalMinutes int
//
// 	title, err := formValueStr(r, "title")
// 	if err != nil {
// 		return model.Recipe{}, fmt.Errorf("could not get title: %v", err)
// 	}
//
// 	cookMinutes, err = formValueInt(r, "cookMinutes")
// 	if err != nil {
// 		return model.Recipe{}, fmt.Errorf("could not get cook minutes: %v", err)
// 	}
//
// 	totalMinutes, err = formValueInt(r, "totalMinutes")
// 	if err != nil {
// 		return model.Recipe{}, fmt.Errorf("could not get total minutes: %v", err)
// 	}
//
// 	rcp.Title = title
// 	rcp.CookMinutes = cookMinutes
// 	rcp.TotalMinutes = totalMinutes
//
// 	for key, val := range r.Form {
// 		if !strings.HasPrefix(key, "sub-") {
// 			continue
// 		}
// 		var subRcp model.SubRecipe
// 	}
//
// 	return rcp, nil
// }
//
// func formValueStr(r *http.Request, key string) (string, error) {
// 	val := r.FormValue(key)
// 	if val == "" {
// 		return "", fmt.Errorf("could not find key %q", key)
// 	}
//
// 	return val, nil
// }
//
// func formValueInt(r *http.Request, key string) (int, error) {
// 	valStr, err := formValueStr(r, key)
// 	if err != nil {
// 		return 0, err
// 	}
//
// 	val, err := strconv.Atoi(valStr)
// 	if err != nil {
// 		return 0, fmt.Errorf("value for key %q is not an integer: %v", err)
// 	}
//
// 	return val, nil
// }
