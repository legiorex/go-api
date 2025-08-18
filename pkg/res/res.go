package res

import (
	"encoding/json"
	"net/http"
)

func Json[T any](w http.ResponseWriter, statusCode int, payload T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(payload)
}
