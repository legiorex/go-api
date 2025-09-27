package main

import (
	"go-api/configs"
	"go-api/internal/link"
	"go-api/internal/stat"
	"go-api/internal/user"
	"go-api/pkg/db"
)

func main() {
	config := configs.LoadConfig()
	dataBase := db.NewDb(config)
	dataBase.GetDB().AutoMigrate(&link.Link{}, &user.User{}, &stat.Stat{})
}
