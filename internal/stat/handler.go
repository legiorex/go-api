package stat

import (
	"go-api/pkg/req"
	"go-api/pkg/res"
	"net/http"
	"time"
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

		params, err := req.HandleQuery[StatPayload](&w, r)

		if err != nil {
			return
		}

		const shortForm = "2006-01-02"

		from, err := time.Parse(shortForm, params.From)

		if err != nil {
			http.Error(w, "Invalid from params", http.StatusBadRequest)
			return
		}

		to, err := time.Parse(shortForm, params.To)

		if err != nil {
			http.Error(w, "Invalid to params", http.StatusBadRequest)
			return
		}

		response := StatGetResponse{
			From: from,
			To:   to,
			By:   params.By,
		}

		res.Json(w, http.StatusOK, response)
	}
}
