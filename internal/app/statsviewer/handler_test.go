package statsviewer_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-encurtador-url/internal/app/mocks"
	"github.com/golang-encurtador-url/internal/app/statsviewer"
	"github.com/golang-encurtador-url/internal/app/urlentities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestViewerHandler(t *testing.T) {

	t.Run("should create handler", func(t *testing.T) {
		service := new(mocks.ViewerServiceMock)
		handler := statsviewer.NewHandler(service)

		assert.NotNil(t, handler)
		assert.Equal(t, "/api/stats/{short}", handler.GetPattern())
		assert.Equal(t, http.MethodGet, handler.GetMethod())
	})

	t.Run("should return 200", func(t *testing.T) {
		url := &urlentities.Url{
			ID:          "id",
			CreatedAt:   time.Now(),
			Destination: "destination",
		}
		stats := &urlentities.Statistics{
			URL:    url,
			Clicks: 2,
		}

		srvMock := new(mocks.ViewerServiceMock)
		srvMock.On("Find", mock.Anything).Return(url)
		srvMock.On("GetStatistics", mock.Anything).Return(stats)

		h := statsviewer.NewHandler(srvMock)
		res := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/stats/id", nil)

		h.ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)

		var target urlentities.Statistics
		err := json.NewDecoder(res.Body).Decode(&target)
		assert.Nil(t, err)
		assert.Equal(t, "id", target.URL.ID)
		assert.Equal(t, "destination", target.URL.Destination)
		assert.Equal(t, 2, target.Clicks)
	})

	t.Run("should return 404", func(t *testing.T) {
		srvMock := new(mocks.ViewerServiceMock)
		srvMock.On("Find", mock.Anything).Return(nil)

		h := statsviewer.NewHandler(srvMock)
		res := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/stats/id", nil)

		h.ServeHTTP(res, req)
		assert.Equal(t, http.StatusNotFound, res.Code)
	})
}
