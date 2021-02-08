package storage

import "github.com/golang-encurtador-url/domain"

type Repository interface {
	ExistsID(id string) bool
	FindByID(id string) *domain.Url
	FindByURL(url string) *domain.Url
	Save(url domain.Url) error
	AddLogClick(id string)
	FindLogClicks(id string) int
}
