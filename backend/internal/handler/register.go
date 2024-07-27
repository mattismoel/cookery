package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mattismoel/cookery/internal/service"
)

func HandleRegister(usrSrv *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("email")
		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")

		password := r.FormValue("password")
		passwordRepeat := r.FormValue("passwordRepeat")

		err := usrSrv.Register(email, firstname, lastname, password, passwordRepeat)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("registered user.")
	}
}
