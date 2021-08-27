package statscollector_test

import (
	"bytes"
	"log"
	"testing"
	"time"

	"github.com/golang-encurtador-url/internal/app/statscollector"
	"github.com/golang-encurtador-url/internal/infra/memrepository"
	"github.com/stretchr/testify/assert"
)

func TestStatsCollector(t *testing.T) {

	t.Run("should collect stats", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		logger := log.New(buffer, "", log.LstdFlags)
		stats := make(chan string)
		clicks := make(map[string]int)
		repo := memrepository.NewMemoryRepository(nil, clicks)

		srv := statscollector.NewStatsCollector(repo, stats, logger)
		go srv.CollectStatistics()
		stats <- "id"
		time.Sleep(1 * time.Second)
		close(stats)

		assert.Contains(t, buffer.String(), "registrado com sucesso")
		assert.Equal(t, clicks["id"], 1)
	})
}
