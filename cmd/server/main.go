package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bobrnor/hl-course/pkg/http/rest"
	"github.com/bobrnor/hl-course/pkg/registration"
	"github.com/bobrnor/hl-course/pkg/storage/memory"
)

func main() {
	if err := run(); err != nil {
		log.Println("app terminated due the error:", err.Error())
		os.Exit(1)
	}
}

func run() error {
	storage := new(memory.Storage)
	registrator := registration.NewService(storage)
	router := rest.Handler(registrator)

	log.Println("The hl-course server is on tap now: http://0.0.0.0:8080")

	return http.ListenAndServe("0.0.0.0:8080", router)
}
