package redirect

import "github.com/golang-encurtador-url/internal/app/urlentities"

type (
	Repository interface {
		FindByID(id string) *urlentities.Url
	}

	service struct {
		stats      chan string
		repository Repository
	}
)

func (s *service) Find(id string) *urlentities.Url {
	return s.repository.FindByID(id)
}

func (s *service) AddStatistics(url *urlentities.Url) {
	s.stats <- url.ID
}

func NewService(repository Repository, stats chan string) *service {
	return &service{
		repository: repository,
		stats:      stats,
	}
}
