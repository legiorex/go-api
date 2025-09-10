package middleware

import (
	"fmt"
	"go-api/pkg/jwt"
	"net/http"
	"strings"
)

func Auth(next http.Handler, j *jwt.JWT) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		token := strings.ReplaceAll(authorization, "Bearer", "")
		cleanToken := strings.TrimSpace(token)

		ok, email := j.Parse(cleanToken)

		fmt.Println(ok)
		fmt.Println(email)

		next.ServeHTTP(w, r)
	})
}
