package statsviewer_test

import (
	"testing"
	"time"

	"github.com/golang-encurtador-url/internal/app/statsviewer"
	"github.com/golang-encurtador-url/internal/app/urlentities"
	"github.com/golang-encurtador-url/internal/infra/memrepository"
	"github.com/stretchr/testify/assert"
)

func TestViewer(t *testing.T) {

	t.Run("should find url", func(t *testing.T) {
		urls := make(map[string]*urlentities.Url)
		repo := memrepository.NewMemoryRepository(urls, nil)
		createdAt := time.Now()
		repo.Save(urlentities.Url{
			ID:          "id",
			CreatedAt:   createdAt,
			Destination: "destination",
		})

		srv := statsviewer.NewService(repo)

		url := srv.Find("id")
		assert.NotNil(t, url)
		assert.Equal(t, createdAt, url.CreatedAt)
		assert.Equal(t, "destination", url.Destination)
	})

	t.Run("should find log clicks", func(t *testing.T) {
		urls := make(map[string]*urlentities.Url)
		clicks := make(map[string]int)
		clicks["id"] = 2
		repo := memrepository.NewMemoryRepository(urls, clicks)

		srv := statsviewer.NewService(repo)
		url := &urlentities.Url{
			ID:          "id",
			CreatedAt:   time.Now(),
			Destination: "destination",
		}
		stats := srv.GetStatistics(url)
		assert.Equal(t, stats.URL.ID, "id")
		assert.Equal(t, stats.Clicks, 2)
	})
}
