package memrepository

import (
	"github.com/golang-encurtador-url/internal/app/urlentities"
	"go.uber.org/fx"
)

func new() *MemoryRepository {
	urls := make(map[string]*urlentities.Url)
	clicks := make(map[string]int)
	return NewMemoryRepository(urls, clicks)
}

var Module = fx.Provide(new)
