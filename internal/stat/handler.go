package stat

import (
	"fmt"
	"go-api/pkg/res"
	"net/http"
)

type StatHandler struct {
	StatRepository StatRepositoryInterface
}

type StatHandlerDeps struct {
	StatRepository StatRepositoryInterface
}

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	statHandler := &StatHandler{
		StatRepository: deps.StatRepository,
	}
	router.HandleFunc("GET /stat", statHandler.GetStat())
}

func (h *StatHandler) GetStat() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		from := r.URL.Query().Get("from")

		to := r.URL.Query().Get("to")

		by := r.URL.Query().Get("by")

		response := StatPayload{
			From: from,
			To:   to,
			By:   by,
		}

		res.Json(w, http.StatusOK, response)
	}
}
