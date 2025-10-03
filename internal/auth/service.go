package auth

import (
	"errors"
	"go-api/internal/user"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository user.UserRepositoryInterface
}

func NewAuthService(userRepository user.UserRepositoryInterface) AuthServiceInterface {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (service *AuthService) Register(email, name, password string) (string, error) {

	existedUser, _ := service.UserRepository.FindByEmail(email)

	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	newUser := &user.User{
		Name:     name,
		Password: string(hashedPassword),
		Email:    email,
	}

	_, err = service.UserRepository.Create(newUser)

	if err != nil {
		return "", err
	}

	return email, nil
}

func (service *AuthService) Login(email, password string) (*user.User, error) {
	user, err := service.UserRepository.FindByEmail(email)

	if err != nil {
		return nil, err
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))

	if err != nil {
		return nil, err
	}

	return user, nil
}
