package auth_test

import (
	"fmt"
	"go-api/internal/auth"
	"go-api/internal/user"
	"testing"
)

type MockUserRepository struct{}

func (r *MockUserRepository) Create(u *user.User) (*user.User, error) {

	return &user.User{}, nil
}
func (r *MockUserRepository) FindByEmail(email string) (*user.User, error) {
	return &user.User{}, nil
}

func TestRegisterSuccess(t *testing.T) {

	authService := auth.NewAuthService(&MockUserRepository{})
	fmt.Println(authService)

}
