package db

import (
	"gorm.io/gorm"
)

type DatabaseInterface interface {
	GetDB() *gorm.DB
}
