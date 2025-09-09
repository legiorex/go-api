package auth

import (
	"errors"
	"go-api/internal/user"
)

type AuthService struct {
	UserRepository user.UserRepositoryInterface
}

type AuthServiceDeps struct {
	UserRepository user.UserRepositoryInterface
}

func NewAuthService(deps AuthServiceDeps) AuthServiceInterface {
	return &AuthService{
		UserRepository: deps.UserRepository,
	}
}

func (service *AuthService) Register(email, name, password string) (string, error) {

	existedUser, _ := service.UserRepository.FindByEmail(email)

	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}

	newUser := &user.User{
		Name:     name,
		Password: password,
		Email:    email,
	}

	_, err := service.UserRepository.Create(newUser)

	if err != nil {
		return "", err
	}

	return email, nil
}
