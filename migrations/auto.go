package main

import (
	"go-api/configs"
	"go-api/internal/link"
	"go-api/pkg/db"
)

func main() {
	config := configs.LoadConfig()
	dataBase := db.NewDb(config)
	dataBase.GetDB().AutoMigrate(&link.Link{})
}
