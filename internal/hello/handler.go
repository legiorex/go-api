package hello

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func NewHelloHandler(router *http.ServeMux) {
	helloHandler := &HelloHandler{}
	router.Handle("/hello", helloHandler.Hello())
}

func (h *HelloHandler) Hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello, World!")
	}
}
