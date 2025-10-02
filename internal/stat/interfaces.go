package stat

import (
	"net/http"
	"time"
)

type Period struct {
	Period string
	Sum    uint
}

type StatData struct {
	From   time.Time
	To     time.Time
	Period string
}

type StatRepositoryInterface interface {
	AddClick(link uint)
	GetStats(data StatData) ([]StatPeriod, error)
}

type StatHandlerInterface interface {
	GetStat() http.HandlerFunc
}

type StatServiceInterface interface {
	AddClick()
}
