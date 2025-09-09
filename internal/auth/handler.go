package auth

import (
	"fmt"
	"go-api/configs"
	"go-api/pkg/req"
	"go-api/pkg/res"
	"net/http"
)

type AuthHandler struct {
	*configs.Config
	AuthService AuthServiceInterface
}
type AuthHandlerDeps struct {
	*configs.Config
	AuthService AuthServiceInterface
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	authHandler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}

	router.Handle("POST /auth/login", authHandler.Login())
	router.Handle("POST /auth/register", authHandler.Register())
}

func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		payload, err := req.HandleBody[LoginRequest](&w, r)

		fmt.Println(payload)

		if err != nil {
			return
		}

		data := LoginPayload{
			Token: "123",
		}
		res.Json(w, http.StatusOK, data)
	}
}

func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := req.HandleBody[RegisterRequest](&w, r)

		fmt.Println(body)

		h.AuthService.Register(body.Email, body.Name, body.Password)

		if err != nil {
			return
		}

		data := RegisterPayload{
			Token: "123",
		}

		res.Json(w, http.StatusOK, data)

	}
}
