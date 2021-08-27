package statsviewer

import "github.com/golang-encurtador-url/internal/app/urlentities"

type (
	Repository interface {
		FindByID(id string) *urlentities.Url
		FindLogClicks(id string) int
	}

	service struct {
		repository Repository
	}
)

func (s *service) GetStatistics(url *urlentities.Url) *urlentities.Statistics {
	clicks := s.repository.FindLogClicks(url.ID)
	return &urlentities.Statistics{
		URL:    url,
		Clicks: clicks,
	}
}

func (s *service) Find(id string) *urlentities.Url {
	return s.repository.FindByID(id)
}

func NewService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}
