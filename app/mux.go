package app

import (
	"net/http"

	"github.com/gorilla/mux"

	"go_api/app/controllers/ping"
)

func NewRouter() (*mux.Router, error) {
	router := mux.NewRouter()
	router.HandleFunc("/ping", ping.Get).Methods(http.MethodGet)

	return router, nil
}
