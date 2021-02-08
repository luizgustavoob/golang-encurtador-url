package server

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	server *http.Server
}

func New(port int, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:         ":" + strconv.Itoa(port),
			Handler:      handler,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 55 * time.Second,
		},
	}
}

func (s *Server) ListenAndServe() {
	go func() {
		log.Printf("golang-encurtador-url running on %s", s.server.Addr)
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("error when starting the service: %q", err)
		}
	}()
}

func (s *Server) Shutdown() {
	log.Printf("shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		log.Printf("unable to shut down the server in 60s: %q", err)
		return
	}
	log.Printf("server gracefully stopped")
}
