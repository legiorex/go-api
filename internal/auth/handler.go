package auth

import (
	"encoding/json"
	"fmt"
	"go-api/configs"
	"go-api/pkg/res"
	"net/http"

	"github.com/go-playground/validator/v10"
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
		var payload LoginRequest

		err := json.NewDecoder(r.Body).Decode(&payload)

		if err != nil {
			res.Json(w, http.StatusBadRequest, err.Error())
			return
		}

		validate := validator.New()

		err = validate.Struct(payload)

		if err != nil {
			res.Json(w, http.StatusBadRequest, err.Error())
			return
		}
		// if payload.Email == "" {
		// 	res.Json(w, http.StatusBadRequest, "Email is required")
		// 	return
		// }

		data := LoginPayload{
			Token: "123",
		}
		res.Json(w, http.StatusOK, data)
	}
}

func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Register")
	}
}
