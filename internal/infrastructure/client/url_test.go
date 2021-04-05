package client_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/golang-encurtador-url/domain"
	"github.com/golang-encurtador-url/internal/infrastructure/client"
	"github.com/golang-encurtador-url/internal/infrastructure/storage"
	"github.com/stretchr/testify/assert"
)

func TestUrl_AllOperations(t *testing.T) {

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
	clicks["abc"] = 1
	clicks["def"] = 1

	t.Run("should check if find url", func(t *testing.T) {
		urlMock := client.UrlClientMock{
			FindFn: func(ID string) *domain.Url {
				return urls[ID]
			},
		}

		url := urlMock.Find("abc")
		assert.Equal(t, 1, urlMock.FindInvokedCount)
		assert.NotNil(t, url)

		url = urlMock.Find("asdasd")
		assert.Equal(t, 2, urlMock.FindInvokedCount)
		assert.Nil(t, url)
	})

	t.Run("should check add log click", func(t *testing.T) {
		urlMock := client.UrlClientMock{
			AddLogClickFn: func(ID string) {
				clicks[ID]++
			},
		}

		urlMock.AddLogClick("abc")
		assert.Equal(t, 1, urlMock.AddLogClickInvokedCount)
		assert.Equal(t, 2, clicks["abc"])

		urlMock.AddLogClick("asdasd")
		assert.Equal(t, 2, urlMock.AddLogClickInvokedCount)
		assert.Equal(t, 1, clicks["asdasd"])
	})

	t.Run("should check find or create URL", func(t *testing.T) {
		repMock := storage.RepositoryMock{
			FindByURLFn: func(destino string) *domain.Url {
				for _, u := range urls {
					if u.Destination == destino {
						return u
					}
				}
				return nil
			},
			SaveFn: func(url domain.Url) error {
				urls[url.ID] = &url
				return nil
			},
		}

		urlMock := client.UrlClientMock{
			RepositoryMock: repMock,
			FindOrCreateURLFn: func(repositoryMock *storage.RepositoryMock, destination string, invokedCount *int) (*domain.Url, bool, error) {
				if destination == "http://www.uol.com" || destination == "http://www.globo.com" {
					url := repositoryMock.FindByURL(destination)
					return url, false, nil
				}

				url := domain.Url{
					ID:          "ID" + strconv.Itoa(*invokedCount),
					CreatedAt:   time.Now(),
					Destination: destination,
				}
				repositoryMock.Save(url)
				return &url, true, nil
			},
		}

		url, created, err := urlMock.FindOrCreateURL("http://www.globo.com")
		assert.NotNil(t, url)
		assert.False(t, created)
		assert.NoError(t, err)

		url, created, err = urlMock.FindOrCreateURL("http://www.facebook.com")
		assert.NotNil(t, url)
		assert.True(t, created)
		assert.NoError(t, err)
	})
}
