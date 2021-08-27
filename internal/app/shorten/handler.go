package shorten

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-encurtador-url/internal/app/urlentities"
)

type (
	Service interface {
		FindOrCreateURL(destination string) (*urlentities.Url, bool, error)
	}

	Logger interface {
		Printf(format string, values ...interface{})
	}

	handler struct {
		service Service
		logger  Logger
		baseUrl string
	}
)

func (h *handler) GetMethod() string {
	return http.MethodPost
}

func (h *handler) GetPattern() string {
	return "/api/encurtar"
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u, created, err := h.service.FindOrCreateURL(h.extractURL(r))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := make(map[string]string)
		error["error"] = err.Error()
		js, _ := json.Marshal(error)
		w.Write(js)
		return
	}

	var status int
	if created {
		status = http.StatusCreated
	} else {
		status = http.StatusOK
	}

	shortURL := fmt.Sprintf("%s/r/%s", h.baseUrl, u.ID)

	h.logger.Printf("URL %s encurtada com sucesso: %s", u.Destination, shortURL)

	w.Header().Add("Location", shortURL)
	w.Header().Add("Link", fmt.Sprintf("<%s/api/stats/%s>; rel=\"stats\"", h.baseUrl, u.ID))
	w.WriteHeader(status)
}

func (h *handler) extractURL(r *http.Request) string {
	u := make([]byte, r.ContentLength)
	r.Body.Read(u)
	return string(u)
}

func NewHandler(service Service, logger Logger, baseUrl string) *handler {
	return &handler{
		service: service,
		logger:  logger,
		baseUrl: baseUrl,
	}
}
