package main

import (
	"fmt"
	"net/http"

	"go-api/configs"
	"go-api/internal/auth"
	"go-api/internal/link"
	"go-api/pkg/db"
)

func main() {
	fmt.Println("Starting server on port 8081")
	config := configs.LoadConfig()

	_ = db.NewDb(config)

	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: config,
	})
	link.NewLinkHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	server.ListenAndServe()
}
