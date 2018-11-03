package main

import (
	"log"
	"os"

	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		log.Println("app terminated due the error:", err.Error())
		os.Exit(1)
	}
}

func run() error {
	return errors.New("nothing is implemented yet")
}
