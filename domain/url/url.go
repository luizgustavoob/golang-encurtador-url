package url

import (
	"github.com/golang-encurtador-url/domain"
	"github.com/golang-encurtador-url/domain/logger"
)

type service struct {
	client domain.URLClient
	stats  chan string
}

func NewService(client domain.URLClient, stats chan string) *service {
	return &service{client, stats}
}

func (s *service) Find(ID string) *domain.Url {
	return s.client.Find(ID)
}

func (s *service) FindOrCreateURL(destination string) (*domain.Url, bool, error) {
	return s.client.FindOrCreateURL(destination)
}

func (s *service) GetStatistics(url *domain.Url) *domain.Statistics {
	return s.client.GetStatistics(url)
}

func (s *service) AddStatistics(url *domain.Url) {
	s.stats <- url.ID
}

func (s *service) CollectStatistics() {
	for id := range s.stats {
		s.client.AddLogClick(id)
		logger.Logar("Click registrado com sucesso para %s.", id)
	}
}
