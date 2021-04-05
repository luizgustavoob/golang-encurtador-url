package storage_test

import (
	"errors"
	"testing"
	"time"

	"github.com/golang-encurtador-url/domain"
	"github.com/golang-encurtador-url/internal/infrastructure/storage"
	"github.com/stretchr/testify/assert"
)

func TestStorage_AllOperations(t *testing.T) {

	// setup
	urls := make(map[string]*domain.Url)
	clicks := make(map[string]int)

	urls["abc"] = &domain.Url{
		ID:          "abc",
		CreatedAt:   time.Now(),
		Destination: "http://www.globo.com",
	}
	urls["def"] = &domain.Url{
		ID:          "def",
		CreatedAt:   time.Now(),
		Destination: "http://www.uol.com",
	}
	clicks["abc"] = 5
	clicks["def"] = 4

	t.Run("not should save url", func(t *testing.T) {
		repositoryMock := storage.RepositoryMock{
			SaveFn: func(url domain.Url) error {
				return errors.New("Ooops, error!")
			},
		}

		err := repositoryMock.Save(domain.Url{
			ID:          "ghi",
			CreatedAt:   time.Now(),
			Destination: "http://www.youtube.com",
		})

		assert.Error(t, err)
		assert.Equal(t, 1, repositoryMock.SaveInvokedCount)
		assert.Equal(t, 2, len(urls))
	})

	t.Run("should save url", func(t *testing.T) {
		repositoryMock := storage.RepositoryMock{
			SaveFn: func(url domain.Url) error {
				urls[url.ID] = &url
				return nil
			},
		}

		err := repositoryMock.Save(domain.Url{
			ID:          "ghi",
			CreatedAt:   time.Now(),
			Destination: "http://www.youtube.com",
		})

		assert.NoError(t, err)
		assert.Equal(t, 1, repositoryMock.SaveInvokedCount)
		assert.Equal(t, 3, len(urls))
	})

	t.Run("should check if ID exists", func(t *testing.T) {
		repositoryMock := storage.RepositoryMock{
			ExistsIDFn: func(id string) bool {
				_, existe := urls[id]
				return existe
			},
		}

		exists := repositoryMock.ExistsID("abc")
		assert.True(t, exists)

		exists = repositoryMock.ExistsID("6540")
		assert.False(t, exists)
	})

	t.Run("should check if map return domain.URL", func(t *testing.T) {
		repositoryMock := storage.RepositoryMock{
			FindByIdFn: func(id string) *domain.Url {
				return urls[id]
			},
		}

		url := repositoryMock.FindByID("abc")
		assert.Equal(t, 1, repositoryMock.FindByIdInvokedCount)
		assert.NotNil(t, url)

		url = repositoryMock.FindByID("a5sd0")
		assert.Equal(t, 2, repositoryMock.FindByIdInvokedCount)
		assert.Nil(t, url)
	})

	t.Run("should check if map return by destination", func(t *testing.T) {
		repositoryMock := storage.RepositoryMock{
			FindByURLFn: func(destino string) *domain.Url {
				for _, u := range urls {
					if u.Destination == destino {
						return u
					}
				}
				return nil
			},
		}

		url := repositoryMock.FindByURL("http://www.globo.com")
		assert.Equal(t, 1, repositoryMock.FindByURLInvokedCount)
		assert.NotNil(t, url)

		url = repositoryMock.FindByURL("http://www.google.com")
		assert.Equal(t, 2, repositoryMock.FindByURLInvokedCount)
		assert.Nil(t, url)
	})

	t.Run("should check add log click", func(t *testing.T) {
		repositoryMock := storage.RepositoryMock{
			AddLogClickFn: func(id string) {
				clicks[id]++
			},
		}

		repositoryMock.AddLogClick("abc")
		assert.Equal(t, 1, repositoryMock.AddLogClickInvokedCount)
		assert.Equal(t, 6, clicks["abc"])

		repositoryMock.AddLogClick("def")
		assert.Equal(t, 2, repositoryMock.AddLogClickInvokedCount)
		assert.Equal(t, 5, clicks["def"])

		repositoryMock.AddLogClick("qwe")
		assert.Equal(t, 3, repositoryMock.AddLogClickInvokedCount)
		assert.Equal(t, 1, clicks["qwe"])
	})

	t.Run("should check find log clicks", func(t *testing.T) {
		repositoryMock := storage.RepositoryMock{
			FindLogClicksFn: func(id string) int {
				return clicks[id]
			},
		}

		log := repositoryMock.FindLogClicks("abc")
		assert.Equal(t, 1, repositoryMock.FindLogClicksInvokedCount)
		assert.Equal(t, 6, log)

		log = repositoryMock.FindLogClicks("def")
		assert.Equal(t, 2, repositoryMock.FindLogClicksInvokedCount)
		assert.Equal(t, 5, log)

		log = repositoryMock.FindLogClicks("qwe")
		assert.Equal(t, 3, repositoryMock.FindLogClicksInvokedCount)
		assert.Equal(t, 1, log)
	})
}
