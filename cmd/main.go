package main

import (
	"fmt"
	"net/http"

	"go-api/internal/auth"
	"go-api/internal/link"
	"go-api/pkg/container"
)

func main() {
	fmt.Println("Starting server on port 8081")

	// Создаем все зависимости
	container := container.NewContainer()

	router := http.NewServeMux()

	// Инжектим готовые зависимости в handlers
	auth.NewAuthHandler(router, container.GetAuthHandlerDeps())
	link.NewLinkHandler(router, container.GetLinkHandlerDeps())

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	server.ListenAndServe()
}
