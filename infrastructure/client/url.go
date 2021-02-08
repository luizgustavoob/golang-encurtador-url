package client

import (
	"math/rand"
	"net/url"
	"time"

	"github.com/golang-encurtador-url/domain"
	"github.com/golang-encurtador-url/infrastructure/storage"
)

const (
	size    = 8
	simbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_-+"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type urlClient struct {
	repository storage.Repository
}

func NewURLClient(repository storage.Repository) *urlClient {
	return &urlClient{repository}
}

func (u *urlClient) Find(ID string) *domain.Url {
	return u.repository.FindByID(ID)
}

func (u *urlClient) FindOrCreateURL(destination string) (*domain.Url, bool, error) {
	if url := u.repository.FindByURL(destination); url != nil {
		return url, false, nil
	}

	if _, err := url.ParseRequestURI(destination); err != nil {
		return nil, false, err
	}

	url := domain.Url{
		ID:          u.generateID(),
		CreatedAt:   time.Now(),
		Destination: destination,
	}
	u.repository.Save(url)
	return &url, true, nil
}

func (u *urlClient) AddLogClick(ID string) {
	u.repository.AddLogClick(ID)
}

func (u *urlClient) GetStatistics(url *domain.Url) *domain.Statistics {
	clicks := u.repository.FindLogClicks(url.ID)
	return &domain.Statistics{
		URL:    url,
		Clicks: clicks,
	}
}

func (u *urlClient) generateID() string {
	newID := func() string {
		id := make([]byte, size, size)
		for i := range id {
			id[i] = simbols[rand.Intn(len(simbols))]
		}
		return string(id)
	}

	for {
		if id := newID(); !u.repository.ExistsID(id) {
			return id
		}
	}
}
