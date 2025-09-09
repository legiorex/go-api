package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func Bearer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		token := strings.ReplaceAll(authorization, "Bearer", "")
		cleanToken := strings.TrimSpace(token)

		fmt.Println(cleanToken)

		next.ServeHTTP(w, r)
	})
}
