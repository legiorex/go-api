package auth

type AuthServiceInterface interface {
	Register(email, name, password string) (string, error)
}
