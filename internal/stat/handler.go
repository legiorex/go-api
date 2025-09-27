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

	fmt.Println("get stat")

	return func(w http.ResponseWriter, r *http.Request) {
		res.Json(w, http.StatusOK, "")
	}
}
