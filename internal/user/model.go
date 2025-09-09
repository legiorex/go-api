package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"index"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"uniqueIndex"`
}
