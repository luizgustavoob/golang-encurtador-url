package app

import (
	"context"

	"github.com/golang-encurtador-url/internal/app/redirect"
	"github.com/golang-encurtador-url/internal/app/shorten"
	"github.com/golang-encurtador-url/internal/app/statscollector"
	"github.com/golang-encurtador-url/internal/app/statsviewer"
	"go.uber.org/fx"
)

func startStatsCollector(lc fx.Lifecycle, collector *statscollector.StatsCollector) {
	lc.Append(fx.Hook{
		OnStart: func(c context.Context) error {
			go collector.CollectStatistics()
			return nil
		},
	})
}

var Module = fx.Options(
	shorten.Module,
	redirect.Module,
	statscollector.Module,
	statsviewer.Module,
	fx.Invoke(startStatsCollector),
)
