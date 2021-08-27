package shorten_test

import (
	"testing"
	"time"

	"github.com/golang-encurtador-url/internal/app/shorten"
	"github.com/golang-encurtador-url/internal/app/urlentities"
	"github.com/golang-encurtador-url/internal/infra/memrepository"
	"github.com/stretchr/testify/assert"
)

func TestShorten(t *testing.T) {

	t.Run("should create url", func(t *testing.T) {
		urls := make(map[string]*urlentities.Url)
		repo := memrepository.NewMemoryRepository(urls, nil)

		srv := shorten.NewService(repo)
		url, created, err := srv.FindOrCreateURL("http://www.google.com.br")

		assert.NotNil(t, url)
		assert.True(t, created)
		assert.Nil(t, err)
		assert.NotEmpty(t, url.ID)
		assert.NotNil(t, url.CreatedAt)
		assert.Equal(t, "http://www.google.com.br", url.Destination)
		assert.Equal(t, urls[url.ID].Destination, url.Destination) //então salvou
	})

	t.Run("should not create because url already exists", func(t *testing.T) {
		urls := make(map[string]*urlentities.Url)
		repo := memrepository.NewMemoryRepository(urls, nil)
		repo.Save(urlentities.Url{
			ID:          "id",
			CreatedAt:   time.Now(),
			Destination: "http://www.google.com.br",
		})

		srv := shorten.NewService(repo)
		url, created, err := srv.FindOrCreateURL("http://www.google.com.br")

		assert.NotNil(t, url)
		assert.False(t, created)
		assert.Nil(t, err)
	})

	t.Run("should not create because destination is invalid", func(t *testing.T) {
		urls := make(map[string]*urlentities.Url)
		repo := memrepository.NewMemoryRepository(urls, nil)

		srv := shorten.NewService(repo)
		url, created, err := srv.FindOrCreateURL("bug")

		assert.Nil(t, url)
		assert.False(t, created)
		assert.NotNil(t, err)
	})
}
