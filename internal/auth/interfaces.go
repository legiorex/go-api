package auth

import "go-api/internal/user"

type AuthServiceInterface interface {
	Register(email, name, password string) (string, error)
	Login(email, password string) (*user.User, error)
}
