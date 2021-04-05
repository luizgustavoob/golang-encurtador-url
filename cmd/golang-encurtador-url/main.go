package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang-encurtador-url/domain/logger"
	"github.com/golang-encurtador-url/domain/url"
	"github.com/golang-encurtador-url/internal/infrastructure/client"
	"github.com/golang-encurtador-url/internal/infrastructure/server"
	"github.com/golang-encurtador-url/internal/infrastructure/storage"
)

var (
	port      *int
	activeLog *bool
)

func init() {
	port = flag.Int("p", 8888, "porta")
	activeLog = flag.Bool("l", true, "log ativo/inativo")
	flag.Parse()
}

func main() {
	stats := make(chan string)
	defer close(stats)

	logger.Configure(activeLog)

	repository := storage.NewMemoryRepository()
	urlClient := client.NewURLClient(repository)
	urlService := url.NewService(urlClient, stats)

	go urlService.CollectStatistics()

	handler := server.NewHandler(urlService, fmt.Sprintf("http://localhost:%d", *port), port)
	server := server.New(*port, handler)
	logger.Logar("Iniciando servidor na porta %d", *port)
	server.ListenAndServe()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	<-stopChan
	server.Shutdown()
}
