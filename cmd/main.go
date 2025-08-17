package main

import (
	"fmt"
	"net/http"

	"go-api/internal/hello"
)

func main() {
	fmt.Println("Starting server on port 8081")

	router := http.NewServeMux()
	hello.NewHelloHandler(router)
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	server.ListenAndServe()
}
