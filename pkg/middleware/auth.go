package middleware

import (
	"context"
	"go-api/pkg/jwt"
	"net/http"
	"strings"
)

type key string

const (
	ContextEmailKey key = "ContextEmailKey"
)

func writeUnauth(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func Auth(next http.Handler, j *jwt.JWT) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		if !strings.HasPrefix(authorization, "Bearer") {
			writeUnauth(w)
			return
		}

		token := strings.ReplaceAll(authorization, "Bearer", "")
		cleanToken := strings.TrimSpace(token)

		isValid, data := j.Parse(cleanToken)

		if !isValid {
			writeUnauth(w)
			return
		}

		ctx := context.WithValue(r.Context(), ContextEmailKey, data.Email)

		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
