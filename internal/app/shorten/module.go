package shorten

import (
	"log"

	app_handler "github.com/golang-encurtador-url/internal/infra/handler"
	"github.com/golang-encurtador-url/internal/infra/memrepository"

	"go.uber.org/fx"
)

func newService(repo *memrepository.MemoryRepository) *service {
	return NewService(repo)
}

func newHandler(service *service, logger *log.Logger) app_handler.HandlerResult {
	return app_handler.HandlerResult{
		Handler: NewHandler(service, logger, "http://localhost:8080"),
	}
}

var Module = fx.Provide(newService, newHandler)
