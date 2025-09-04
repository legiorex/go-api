package link

import (
	"go-api/pkg/req"
	"go-api/pkg/res"
	"net/http"
)

type LinkHandler struct{}

func NewLinkHandler(router *http.ServeMux) {
	linkHandler := &LinkHandler{}

	router.Handle("POST /link/create", linkHandler.Create())
	router.Handle("GET /links", linkHandler.GetLinks())
	router.Handle("PATCH /link/update", linkHandler.Update())
	router.Handle("DELETE /link/delete", linkHandler.Delete())
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
func (h *LinkHandler) Update() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		res.Json(w, http.StatusNoContent, "update")
	}
}
func (h *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res.Json(w, http.StatusNoContent, "delete")
	}
}
