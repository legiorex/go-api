package user

type UserRepositoryInterface interface {
	Create(user *User) (*User, error)
	FindByEmail(email string) (*User, error)
}
