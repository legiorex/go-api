package user

import (
	"go-api/pkg/db"
)

type UserRepository struct {
	Database db.DatabaseInterface
}

func NewUserRepository(database db.DatabaseInterface) UserRepositoryInterface {
	return &UserRepository{
		Database: database,
	}
}

func (repo *UserRepository) Create(user *User) (*User, error) {

	result := repo.Database.GetDB().Create(user)

	if result.Error != nil {
		return nil, result.Error

	}

	return user, nil

}

func (repo *UserRepository) FindByEmail(email string) (*User, error) {

	var user User

	result := repo.Database.GetDB().Take(&user, "email = ?", email)

	if result.Error != nil {
		return nil, result.Error

	}

	return &user, nil

}
