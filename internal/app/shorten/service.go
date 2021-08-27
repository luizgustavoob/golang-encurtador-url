package shorten

import (
	"math/rand"
	"net/url"
	"time"

	"github.com/golang-encurtador-url/internal/app/urlentities"
)

type (
	Repository interface {
		ExistsID(id string) bool
		FindByURL(url string) *urlentities.Url
		Save(url urlentities.Url) error
	}

	service struct {
		repository Repository
	}
)

func (s *service) FindOrCreateURL(destination string) (*urlentities.Url, bool, error) {
	if url := s.repository.FindByURL(destination); url != nil {
		return url, false, nil
	}

	if _, err := url.ParseRequestURI(destination); err != nil {
		return nil, false, err
	}

	url := urlentities.Url{
		ID:          s.generateID(),
		CreatedAt:   time.Now(),
		Destination: destination,
	}

	s.repository.Save(url)
	return &url, true, nil
}

func (s *service) generateID() string {
	const (
		size    = 8
		simbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_-+"
	)

	rand.Seed(time.Now().UnixNano())

	newID := func() string {
		id := make([]byte, size)
		for i := range id {
			id[i] = simbols[rand.Intn(len(simbols))]
		}
		return string(id)
	}

	for {
		if id := newID(); !s.repository.ExistsID(id) {
			return id
		}
	}
}

func NewService(repo Repository) *service {
	return &service{
		repository: repo,
	}
}
