package stat

import (
	"go/adv-dev/configs"
	"go/adv-dev/pkg/middleware"
	"go/adv-dev/pkg/res"
	"net/http"
	"time"
)

const (
	GroupByDay   = "day"
	GroupByMonth = "month"
)

type StatHandler struct {
	StatRepository *StatRepository
	StatService    *StatService
	Config         *configs.Config
}

type StatHandlerDeps struct {
	StatRepository *StatRepository
	StatService    *StatService
	Config         *configs.Config
}

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := StatHandler{
		StatRepository: deps.StatRepository,
		Config:         deps.Config,
	}
	router.Handle("GET /stat", middleware.IsAuthed(handler.get(), handler.Config))
}

func (handler *StatHandler) get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		from, err := time.Parse(time.DateOnly, r.URL.Query().Get("from"))
		if err != nil {
			http.Error(w, "invalid from date", http.StatusBadRequest)
			return
		}
		to, err := time.Parse(time.DateOnly, r.URL.Query().Get("to"))
		if err != nil {
			http.Error(w, "invalid to date", http.StatusBadRequest)
			return
		}
		by := r.URL.Query().Get("by")
		if by != GroupByDay && by != GroupByMonth {
			http.Error(w, "invalid filter_by value", http.StatusBadRequest)
			return
		}
		stats := handler.StatRepository.GetStats(by, from, to)
		res.JsonResponse(w, stats, http.StatusOK)
	}
}
