package main

import (
	"log"
	"net/http"
	"os"

	"github.com/pkg/errors"

	"github.com/go-redis/redis"

	"github.com/bobrnor/hl-course/pkg/authenticating"
	"github.com/bobrnor/hl-course/pkg/editing"

	"github.com/bobrnor/hl-course/pkg/http/rest"
	"github.com/bobrnor/hl-course/pkg/registration"
	"github.com/bobrnor/hl-course/pkg/storage/memory"
	redisstorage "github.com/bobrnor/hl-course/pkg/storage/redis"
)

func main() {
	if err := runWithRedisStorage(); err != nil {
		log.Println("app terminated due the error:", err.Error())
		os.Exit(1)
	}
}

func runWithMemoryStorage() error {
	storage := new(memory.Storage)

	registrator := registration.NewService(storage)
	editor := editing.NewService(storage)
	authenticator := authenticating.NewService(storage)

	router := rest.Handler(registrator, editor, authenticator)

	log.Println("The hl-course server is on tap now: http://0.0.0.0:8080")

	return http.ListenAndServe("0.0.0.0:8080", router)
}

func runWithRedisStorage() error {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "hl-redis-service:6379",
	})

	if _, err := redisClient.Ping().Result(); err != nil {
		return errors.WithStack(err)
	}

	storage := redisstorage.NewStorage(redisClient)

	registrator := registration.NewService(storage)
	editor := editing.NewService(storage)
	authenticator := authenticating.NewService(storage)

	router := rest.Handler(registrator, editor, authenticator)

	log.Println("The hl-course server is on tap now: http://0.0.0.0:8080")

	return http.ListenAndServe("0.0.0.0:8080", router)
}
