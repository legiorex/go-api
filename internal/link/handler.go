package link

import (
	"fmt"
	"go-api/pkg/event"
	"go-api/pkg/jwt"
	"go-api/pkg/middleware"
	"go-api/pkg/req"
	"go-api/pkg/res"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type LinkHandler struct {
	LinkRepository LinkRepositoryInterface
	EventBus       event.EventBusInterface
	*jwt.JWT
}
type LinkHandlerDeps struct {
	LinkRepository LinkRepositoryInterface
	EventBus       event.EventBusInterface

	*jwt.JWT
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	linkHandler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
		EventBus:       deps.EventBus,
	}

	router.HandleFunc("POST /link", linkHandler.Create())
	router.HandleFunc("GET /links", linkHandler.GetLinks())
	router.HandleFunc("GET /{hash}", linkHandler.GoTo())
	router.Handle("PATCH /link/{id}", middleware.Auth(linkHandler.Update(), deps.JWT))
	router.HandleFunc("DELETE /link/{id}", linkHandler.Delete())
}

func (h *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := req.HandleBody[LinkRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		link := NewLink(body.Url)
		for {
			existedLink, _ := h.LinkRepository.GetByHash(link.Hash)

			if existedLink == nil {
				break
			}
			link.GenerateHash()
		}

		createdLink, err := h.LinkRepository.Create(link)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, http.StatusCreated, createdLink)
	}
}

func (h *LinkHandler) GetLinks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		page := r.URL.Query().Get("page")
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset, err := strconv.Atoi(r.URL.Query().Get("offset"))

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("page", page)
		fmt.Println("limit", limit)
		fmt.Println("offset", offset)

		links, err := h.LinkRepository.GetAllPagination(limit, offset)
		total := h.LinkRepository.GetCount()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := LinksPaginationPayload{
			Links: links,
			Total: total,
		}

		res.Json(w, http.StatusOK, response)

	}
}
func (h *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		hash := r.PathValue("hash")

		link, err := h.LinkRepository.GetByHash(hash)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		go h.EventBus.Publish(event.Event{
			Type: event.TypeLinkVisitedEvent,
			Data: link.ID,
		})

		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}
func (h *LinkHandler) Update() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		idInt, err := strconv.Atoi(id)

		email, ok := r.Context().Value(middleware.ContextEmailKey).(string)

		fmt.Println(email)
		fmt.Println(ok)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		body, err := req.HandleBody[LinkUpdateRequest](&w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		link, err := h.LinkRepository.Update(&Link{
			Model: gorm.Model{ID: uint(idInt)},
			Url:   body.Url,
			Hash:  body.Hash,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, http.StatusOK, link)
	}
}
func (h *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := r.PathValue("id")

		idInt, err := strconv.Atoi(id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = h.LinkRepository.GetByID(uint(idInt))

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		err = h.LinkRepository.Delete(uint(idInt))

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, http.StatusOK, "success")
	}
}
