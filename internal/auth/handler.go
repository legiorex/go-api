package auth

import (
	"go-api/configs"
	"go-api/pkg/jwt"
	"go-api/pkg/req"
	"go-api/pkg/res"
	"net/http"
)

type AuthHandler struct {
	*configs.Config
	AuthService AuthServiceInterface
	*jwt.JWT
}
type AuthHandlerDeps struct {
	*configs.Config
	AuthService AuthServiceInterface
	*jwt.JWT
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	authHandler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
		JWT:         deps.JWT,
	}

	router.Handle("POST /auth/login", authHandler.Login())
	router.Handle("POST /auth/register", authHandler.Register())
}

func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](&w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := h.AuthService.Login(body.Email, body.Password)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		token, err := h.JWT.Create(user.Email)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response := LoginPayload{
			Token: token,
			Email: user.Email,
			Name:  user.Name,
		}

		res.Json(w, http.StatusOK, response)
	}
}

func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](&w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = h.AuthService.Register(body.Email, body.Name, body.Password)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, err := h.JWT.Create(body.Email)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response := RegisterPayload{
			Token: token,
			Email: body.Email,
		}

		res.Json(w, http.StatusOK, response)

	}
}
