package auth

import (
	"encoding/json"
	"fmt"
	"go-api/configs"
	"net/http"
)

type AuthHandler struct {
	*configs.Config
}
type AuthHandlerDeps struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	authHandler := &AuthHandler{
		Config: deps.Config,
	}

	router.Handle("POST /auth/login", authHandler.Login())
	router.Handle("POST /auth/register", authHandler.Register())
}

func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(h.Config.Auth.Secret)
		fmt.Println("Login")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		res := LoginPayload{
			Token: "123",
		}
		json.NewEncoder(w).Encode(res)
	}
}

func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Register")
	}
}
