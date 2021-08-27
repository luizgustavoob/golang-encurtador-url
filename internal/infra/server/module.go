package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.uber.org/fx"
)

func startServer(lc fx.Lifecycle, handler http.Handler, logger *log.Logger) {
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Println("Starting server at port 8080")
			go srv.ListenAndServe()
			return nil
		},

		OnStop: func(ctx context.Context) error {
			logger.Println("Stopping server...")
			return srv.Shutdown(ctx)
		},
	})
}

var Module = fx.Invoke(startServer)
