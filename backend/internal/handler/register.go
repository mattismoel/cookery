package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mattismoel/cookery/internal/service"
)

func HandleRegister(usrSrv *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		email := r.FormValue("email")
		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")

		password := r.FormValue("password")
		passwordRepeat := r.FormValue("passwordRepeat")

		err := usrSrv.Register(ctx, email, firstname, lastname, password, passwordRepeat)
		if err != nil {
			apiError(w, http.StatusUnauthorized, "Could not register user: %s", err.Error())
			log.Println(err)
		}

		fmt.Println("registered user.")
	}
}
