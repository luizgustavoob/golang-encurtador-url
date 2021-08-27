package memrepository

import (
	"github.com/golang-encurtador-url/internal/app/urlentities"
)

type MemoryRepository struct {
	urls   map[string]*urlentities.Url
	clicks map[string]int
}

func (mr *MemoryRepository) ExistsID(id string) bool {
	_, existe := mr.urls[id]
	return existe
}

func (mr *MemoryRepository) FindByID(id string) *urlentities.Url {
	return mr.urls[id]
}

func (mr *MemoryRepository) FindByURL(destino string) *urlentities.Url {
	for _, u := range mr.urls {
		if u.Destination == destino {
			return u
		}
	}
	return nil
}

func (mr *MemoryRepository) Save(url urlentities.Url) error {
	mr.urls[url.ID] = &url
	return nil
}

func (mr *MemoryRepository) AddLogClick(id string) {
	mr.clicks[id]++
}

func (mr *MemoryRepository) FindLogClicks(id string) int {
	return mr.clicks[id]
}

func NewMemoryRepository(urls map[string]*urlentities.Url, clicks map[string]int) *MemoryRepository {
	return &MemoryRepository{
		urls:   urls,
		clicks: clicks,
	}
}
