package statscollector

import (
	"log"

	"github.com/golang-encurtador-url/internal/infra/memrepository"
	"go.uber.org/fx"
)

func new(repo *memrepository.MemoryRepository, stats chan string, logger *log.Logger) *StatsCollector {
	return NewStatsCollector(repo, stats, logger)
}

var Module = fx.Provide(new)
