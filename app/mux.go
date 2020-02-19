package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	ping "https://github.com/rlapar/go-api/ping"
)

func NewRouter() (*mux.Router, error) {
	// TODO init
	// TODO setup logger
	// TODO vault client init
	// TODO fetch creds from vault
	// TODO setup sentry

	router := mux.NewRouter()
	router.HandleFunc("/ping", ping.Get)

	return router, nil
}

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
