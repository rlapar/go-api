package convert

import (
	"encoding/json"
	"net/http"

	"github.com/thedevsaddam/govalidator"
)

type Request struct {
	Foo     string `json:"foo"`
	FooInt  int64  `json:"foo_int"`
	FooBool bool   `json:"foo_bool"`
}

type Response struct {
	Pong string `json:"pong"`
}

func Post(w http.ResponseWriter, r *http.Request) {
	// -- validate request --
	requestBody := &Request{}

	defer r.Body.Close()

	validator := govalidator.New(govalidator.Options{
		Request:         r,
		Data:            requestBody,
		RequiredDefault: true,
		Rules: govalidator.MapData{
			"foo":      []string{"required"},
			"foo_int":  []string{"numeric_between:1,"},
			"foo_bool": []string{"bool"},
		},
	})
	e := validator.ValidateJSON()

	if len(e) > 0 {
		validationErrors := map[string]interface{}{"validation_errors": e}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(validationErrors)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(Response{Pong: "ok"})
}
