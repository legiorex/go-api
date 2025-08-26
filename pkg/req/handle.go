package req

import (
	"go-api/pkg/res"
	"net/http"
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
