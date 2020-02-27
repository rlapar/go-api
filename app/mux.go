package app

import (
	"net/http"

	"github.com/gorilla/mux"

	"go_api/app/controllers/convert"
	"go_api/app/controllers/ping"
)

func NewRouter() (*mux.Router, error) {
	router := mux.NewRouter()
	router.HandleFunc("/ping", ping.Get).Methods(http.MethodGet)
	router.HandleFunc("/convert", convert.Post).Methods(http.MethodPost)

	return router, nil
}
