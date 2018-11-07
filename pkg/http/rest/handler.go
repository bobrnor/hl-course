package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sanity-io/litter"

	"github.com/bobrnor/hl-course/pkg/authenticating"

	"github.com/bobrnor/hl-course/pkg/editing"

	"github.com/julienschmidt/httprouter"

	"github.com/bobrnor/hl-course/pkg/registration"
)

const (
	UserIDKey = "user-id"
)

func Handler(r registration.Service, e editing.Service, a authenticating.Service) http.Handler {
	router := httprouter.New()

	router.GET("/healthz", liveness())
	router.POST("/user", registerUser(r))
	router.PATCH("/profile/:token", authMiddleware(editProfile(e), a))

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

		log.Println("registered:", litter.Sdump(u))

		payload.Data = u
		payload.Write(w, http.StatusOK)
	}
}

func editProfile(s editing.Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var payload Payload

		userID, ok := r.Context().Value(UserIDKey).(int)
		if !ok {
			log.Println("can't fetch user id from request context")

			payload.Error = fmt.Sprint("can't fetch user id from request context")
			payload.Write(w, http.StatusInternalServerError)
			return
		}

		var profile editing.Profile
		if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
			log.Println("can't decode request body: ", err.Error())

			payload.Error = fmt.Sprintf("can't decode request body: %s", err.Error())
			payload.Write(w, http.StatusBadRequest)
			return
		}

		if err := s.EditProfile(userID, profile); err != nil {
			log.Println("can't register new user:", err.Error())

			payload.Error = fmt.Sprintf("can't register new user: %s", err.Error())
			payload.Write(w, http.StatusInternalServerError)
			return
		}

		log.Println("edited:", userID, litter.Sdump(profile))

		payload.Write(w, http.StatusOK)
	}
}

func authMiddleware(next httprouter.Handle, a authenticating.Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var payload Payload

		token := p.ByName("token")

		id, err := a.Authenticate(token)
		if err != nil {
			log.Println("can't authenticate user: ", err.Error())

			payload.Error = fmt.Sprintf("can't authenticate user: %s", err.Error())
			payload.Write(w, http.StatusInternalServerError)
			return
		}

		log.Println("auth:", id, token)

		ctx := context.WithValue(r.Context(), UserIDKey, id)
		next(w, r.WithContext(ctx), p)
	}
}

func liveness() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("X-Healthz-Header", "Awesome")
		w.WriteHeader(http.StatusOK)
	}
}
