package statsviewer

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/golang-encurtador-url/internal/app/urlentities"
)

type (
	Service interface {
		Find(id string) *urlentities.Url
		GetStatistics(url *urlentities.Url) *urlentities.Statistics
	}

	handler struct {
		service Service
	}
)

func (h *handler) GetMethod() string {
	return http.MethodGet
}

func (h *handler) GetPattern() string {
	return "/api/stats/{short}"
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "short")

	if url := h.service.Find(id); url != nil {
		stats := h.service.GetStatistics(url)
		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, stats)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func NewHandler(service Service) *handler {
	return &handler{service: service}
}
