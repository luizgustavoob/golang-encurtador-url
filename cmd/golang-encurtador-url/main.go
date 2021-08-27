package main

import (
	"context"
	"log"
	"os"

	"github.com/golang-encurtador-url/internal/app"
	"github.com/golang-encurtador-url/internal/infra"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Options(
			fx.Provide(
				newLogger,
				newStatsChannel,
			),
			fx.Invoke(hookCloseChannel),
		),
		app.Module,
		infra.Module,
	).Run()
}

func newLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags)
}

func newStatsChannel() chan string {
	return make(chan string)
}

func hookCloseChannel(lc fx.Lifecycle, stats chan string) {
	lc.Append(fx.Hook{
		OnStop: func(c context.Context) error {
			close(stats)
			return nil
		},
	})
}
