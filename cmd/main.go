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

	db := db.NewDb(config)

	// Repositories

	linkRepository := link.NewLinkRepository(db)

	router := http.NewServeMux()

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: config,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	server.ListenAndServe()
}
