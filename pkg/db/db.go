package db

import (
	"go-api/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	db *gorm.DB
}

func NewDb(config *configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(config.Db.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Db{db}
}

func (db Db) GetDB() *gorm.DB {
	return db.db
}
