package app

import (
	"encoding/json"
	"net/http"
)

type PingResponse struct {
	pong string
}

func Get() (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(PingResponse{pong: "ok"})

	return
}
