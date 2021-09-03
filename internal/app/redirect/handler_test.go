package redirect_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-encurtador-url/internal/app/mocks"
	"github.com/golang-encurtador-url/internal/app/redirect"
	"github.com/golang-encurtador-url/internal/app/urlentities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRedirectHandler(t *testing.T) {

	t.Run("should create handler", func(t *testing.T) {
		service := new(mocks.RedirectServiceMock)
		handler := redirect.NewHandler(service)

		assert.NotNil(t, handler)
		assert.Equal(t, "/r/{short}", handler.GetPattern())
		assert.Equal(t, http.MethodGet, handler.GetMethod())

	})

	t.Run("should return 301", func(t *testing.T) {
		createdAt := time.Now()
		url := &urlentities.Url{
			ID:          "id",
			CreatedAt:   createdAt,
			Destination: "destination",
		}
		stats := make(map[string]int)
		srvMock := new(mocks.RedirectServiceMock)
		srvMock.On("Find", mock.Anything).Return(url)
		srvMock.AddStatisticsFn = func(urlParam *urlentities.Url) {
			stats[urlParam.ID] = 1
		}

		h := redirect.NewHandler(srvMock)
		res := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/r/id", nil)

		h.ServeHTTP(res, req)
		assert.Equal(t, http.StatusMovedPermanently, res.Code)
		assert.Contains(t, res.Header().Get("Location"), "destination")
		assert.Equal(t, 1, stats["id"])
	})

	t.Run("should return 404", func(t *testing.T) {
		srvMock := new(mocks.RedirectServiceMock)
		srvMock.On("Find", mock.Anything).Return(nil)

		h := redirect.NewHandler(srvMock)
		res := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/r/id", nil)

		h.ServeHTTP(res, req)
		assert.Equal(t, http.StatusNotFound, res.Code)
	})
}
