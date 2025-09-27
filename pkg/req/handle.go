package req

import (
	"go-api/pkg/res"
	"net/http"

	"github.com/gorilla/schema"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {

	body, err := Decode[T](r.Body)

	if err != nil {
		res.Json(*w, http.StatusBadRequest, err.Error())
		return nil, err
	}

	err = Validate(body)

	if err != nil {
		res.Json(*w, http.StatusBadRequest, err.Error())
		return nil, err
	}
	return &body, nil
}

func HandleQuery[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	var decoder = schema.NewDecoder()
	var params T
	decoder.Decode(&params, r.URL.Query())

	err := Validate(params)
	if err != nil {
		res.Json(*w, http.StatusBadRequest, err.Error())
		return nil, err
	}

	return &params, nil
}
