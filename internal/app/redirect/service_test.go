package redirect_test

import (
	"testing"
	"time"

	"github.com/golang-encurtador-url/internal/app/redirect"
	"github.com/golang-encurtador-url/internal/app/urlentities"
	"github.com/golang-encurtador-url/internal/infra/memrepository"
	"github.com/stretchr/testify/assert"
)

func TestRedirect(t *testing.T) {

	t.Run("should find url", func(t *testing.T) {
		urls := make(map[string]*urlentities.Url)
		repo := memrepository.NewMemoryRepository(urls, nil)
		createdAt := time.Now()
		repo.Save(urlentities.Url{
			ID:          "id",
			CreatedAt:   createdAt,
			Destination: "destination",
		})

		srv := redirect.NewService(repo, nil)

		url := srv.Find("id")
		assert.NotNil(t, url)
		assert.Equal(t, createdAt, url.CreatedAt)
		assert.Equal(t, "destination", url.Destination)
	})

	t.Run("should add statistc", func(t *testing.T) {
		url := &urlentities.Url{
			ID:          "id",
			CreatedAt:   time.Now(),
			Destination: "destination",
		}
		stats := make(chan string)
		srv := redirect.NewService(nil, stats)

		go srv.AddStatistics(url)
		id := <-stats
		close(stats)
		assert.Equal(t, "id", id)
	})
}
