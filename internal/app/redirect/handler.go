package redirect

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/golang-encurtador-url/internal/app/urlentities"
)

type (
	Service interface {
		AddStatistics(url *urlentities.Url)
		Find(id string) *urlentities.Url
	}

	handler struct {
		service Service
	}
)

func (h *handler) GetMethod() string {
	return http.MethodGet
}

func (h *handler) GetPattern() string {
	return "/r/{short}"
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "short")

	if url := h.service.Find(id); url != nil {
		http.Redirect(w, r, url.Destination, http.StatusMovedPermanently)
		h.service.AddStatistics(url)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}
