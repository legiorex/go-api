package auth_test

import (
	"go-api/internal/auth"
	"go-api/internal/user"
	"testing"
)

type MockUserRepository struct{}

func (r *MockUserRepository) Create(u *user.User) (*user.User, error) {

	return &user.User{
		Email: "w@w.ru",
	}, nil
}
func (r *MockUserRepository) FindByEmail(email string) (*user.User, error) {
	return nil, nil
}

func TestRegisterSuccess(t *testing.T) {

	const initEmail = "a@a.ru"

	authService := auth.NewAuthService(&MockUserRepository{})

	email, err := authService.Register(initEmail, "Jon", "12345")

	if err != nil {
		t.Fatal(err)
	}

	if email != initEmail {
		t.Fatalf("Email %s do not math %s", email, initEmail)
	}

}
