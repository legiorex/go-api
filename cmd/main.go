package main

import (
	"fmt"
	"net/http"

	"go-api/internal/auth"
	"go-api/internal/link"
	"go-api/internal/stat"
	"go-api/pkg/container"
	"go-api/pkg/middleware"
)

func App() http.Handler {
	fmt.Println("Starting server on port 8081")

	// Создаем все зависимости
	container := container.NewContainer()

	router := http.NewServeMux()

	// Инжектим готовые зависимости в handlers
	auth.NewAuthHandler(router, container.GetAuthHandlerDeps())
	link.NewLinkHandler(router, container.GetLinkHandlerDeps())
	stat.NewStatHandler(router, container.GetStatHandlerDeps())

	stack := middleware.Chain(
		middleware.Logging,
		middleware.CORS,
	)

	return stack(router)
}

func main() {
	app := App()

	server := http.Server{
		Addr:    ":8081",
		Handler: app,
	}

	server.ListenAndServe()
}
