package convert

import (
	"encoding/json"
	"net/http"

	"github.com/thedevsaddam/govalidator"

	"go_api/app/converter"
)

type Request struct {
	HtmlString string `json:"html_string"`
}

type Response struct {
	Content []byte `json:"content"`
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
			"html_string": []string{"required"},
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

	content, err := converter.HtmlToPDF(requestBody.HtmlString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(Response{
		Content: content,
		// Content: base64.StdEncoding.EncodeToString(content),
		// TODO convert to base64 String
	})
}
