package statscollector

type (
	Logger interface {
		Printf(format string, values ...interface{})
	}

	Repository interface {
		AddLogClick(id string)
	}

	StatsCollector struct {
		repository Repository
		logger     Logger
		stats      chan string
	}
)

func (s *StatsCollector) CollectStatistics() {
	for id := range s.stats {
		s.repository.AddLogClick(id)
		s.logger.Printf("Click registrado com sucesso para %s.", id)
	}
}

func NewStatsCollector(repo Repository, stats chan string, logger Logger) *StatsCollector {
	return &StatsCollector{
		repository: repo,
		stats:      stats,
		logger:     logger,
	}
}
