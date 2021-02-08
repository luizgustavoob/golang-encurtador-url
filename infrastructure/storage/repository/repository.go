package storage

import "github.com/golang-encurtador-url/domain"

type memoryRepository struct {
	urls   map[string]*domain.Url
	clicks map[string]int
}

func NewMemoryRepository() *memoryRepository {
	return &memoryRepository{
		make(map[string]*domain.Url),
		make(map[string]int),
	}
}

func (mr *memoryRepository) ExistsID(id string) bool {
	_, existe := mr.urls[id]
	return existe
}

func (mr *memoryRepository) FindByID(id string) *domain.Url {
	return mr.urls[id]
}

func (mr *memoryRepository) FindByURL(destino string) *domain.Url {
	for _, u := range mr.urls {
		if u.Destination == destino {
			return u
		}
	}
	return nil
}

func (mr *memoryRepository) Save(url domain.Url) error {
	mr.urls[url.ID] = &url
	return nil
}

func (mr *memoryRepository) AddLogClick(id string) {
	mr.clicks[id]++
}

func (mr *memoryRepository) FindLogClicks(id string) int {
	return mr.clicks[id]
}
