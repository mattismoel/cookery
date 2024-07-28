package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/mattismoel/cookery/internal/service"
)

func HandleLogin(usrSrv *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		email := r.FormValue("email")
		password := r.FormValue("password")

		usr, err := usrSrv.Login(ctx, email, password)
		if err != nil {
			apiError(w, http.StatusUnauthorized, "Invalid login information.")
			log.Printf("could not login user: %v\n", err)
			return
		}

		err = json.NewEncoder(w).Encode(usr)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
