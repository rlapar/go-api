package main

import (
	"log"
	"net/http"

	"go_api/app"
)

func main() {
	router, _ := app.NewRouter()
	log.Fatal(http.ListenAndServe(":8081", router))
}
