package link

import (
	"go-api/pkg/req"
	"go-api/pkg/res"
	"net/http"
)

type LinkHandler struct {
	LinkRepository LinkRepositoryInterface
}
type LinkHandlerDeps struct {
	LinkRepository LinkRepositoryInterface
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	linkHandler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}

	router.Handle("POST /link", linkHandler.Create())
	router.Handle("GET /links", linkHandler.GetLinks())
	router.Handle("GET /{alias}", linkHandler.GoTo())
	router.Handle("PATCH /link/{id}", linkHandler.Update())
	router.Handle("DELETE /link/{id}", linkHandler.Delete())
}

func (h *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		payload, err := req.HandleBody[LinkRequest](&w, r)

		if err != nil {
			return
		}

		res.Json(w, http.StatusCreated, payload)
	}
}

func (h *LinkHandler) GetLinks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res.Json(w, http.StatusOK, "links")
	}
}
func (h *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res.Json(w, http.StatusOK, "go to")
	}
}
func (h *LinkHandler) Update() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		res.Json(w, http.StatusOK, id)
	}
}
func (h *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := r.PathValue("id")

		res.Json(w, http.StatusOK, id)
	}
}
