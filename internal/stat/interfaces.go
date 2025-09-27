package stat

import "net/http"

type StatRepositoryInterface interface {
	AddClick(link uint)
}

type StatHandlerInterface interface {
	GetStat() http.HandlerFunc
}

type StatServiceInterface interface {
	AddClick()
}
