package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/bobrnor/hl-course/pkg/registration"
)

func Handler(r registration.Service) http.Handler {
	router := httprouter.New()

	router.POST("/user", registerUser(r))
	return router
}

func registerUser(s registration.Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		payload := Payload{}

		u, err := s.RegisterUser()
		if err != nil {
			log.Println("can't register new user:", err.Error())

			payload.Error = fmt.Sprintf("can't register new user: %s", err.Error())
			payload.Write(w, http.StatusInternalServerError)
			return
		}

		payload.Data = u
		payload.Write(w, http.StatusOK)
	}
}
