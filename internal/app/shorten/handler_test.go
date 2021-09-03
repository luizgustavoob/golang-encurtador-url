package shorten_test

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-encurtador-url/internal/app/mocks"
	"github.com/golang-encurtador-url/internal/app/shorten"
	"github.com/golang-encurtador-url/internal/app/urlentities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestShortenHandler(t *testing.T) {

	buffer := &bytes.Buffer{}
	logger := log.New(buffer, "", log.LstdFlags)

	t.Run("should create handler", func(t *testing.T) {
		service := new(mocks.ShortenServiceMock)
		handler := shorten.NewHandler(service, logger, "")

		assert.NotNil(t, handler)
		assert.Equal(t, "/api/encurtar", handler.GetPattern())
		assert.Equal(t, http.MethodPost, handler.GetMethod())
	})

	t.Run("should return 200", func(t *testing.T) {
		url := &urlentities.Url{
			ID:          "id",
			CreatedAt:   time.Now(),
			Destination: "destination",
		}
		srvMock := new(mocks.ShortenServiceMock)
		srvMock.On("FindOrCreateURL", mock.Anything).Return(url, false, nil)

		h := shorten.NewHandler(srvMock, logger, "baseUrl")
		res := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/encurtar", bytes.NewReader([]byte(`meencurte`)))

		h.ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Contains(t, res.Header().Get("Location"), "baseUrl")
		assert.Contains(t, res.Header().Get("Location"), "id")
		assert.Contains(t, res.Header().Get("Link"), "baseUrl")
		assert.Contains(t, res.Header().Get("Link"), "id")
		assert.Contains(t, buffer.String(), "sucesso")
	})

	t.Run("should return 201", func(t *testing.T) {
		url := &urlentities.Url{
			ID:          "id",
			CreatedAt:   time.Now(),
			Destination: "destination",
		}
		srvMock := new(mocks.ShortenServiceMock)
		srvMock.On("FindOrCreateURL", mock.Anything).Return(url, true, nil)

		h := shorten.NewHandler(srvMock, logger, "baseUrl")
		res := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/encurtar", bytes.NewReader([]byte(`meencurte`)))

		h.ServeHTTP(res, req)
		assert.Equal(t, http.StatusCreated, res.Code)
		assert.Contains(t, res.Header().Get("Location"), "baseUrl")
		assert.Contains(t, res.Header().Get("Location"), "id")
		assert.Contains(t, res.Header().Get("Link"), "baseUrl")
		assert.Contains(t, res.Header().Get("Link"), "id")
		assert.Contains(t, buffer.String(), "sucesso")
	})

	t.Run("should return 400", func(t *testing.T) {
		srvMock := new(mocks.ShortenServiceMock)
		srvMock.On("FindOrCreateURL", mock.Anything).Return(nil, false, errors.New("handler error"))

		h := shorten.NewHandler(srvMock, logger, "baseUrl")
		res := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/encurtar", bytes.NewReader([]byte(`meencurte`)))

		h.ServeHTTP(res, req)
		assert.Equal(t, http.StatusBadRequest, res.Code)

		responseBody, _ := io.ReadAll(res.Body)
		assert.Contains(t, string(responseBody), "handler error")
	})

}
