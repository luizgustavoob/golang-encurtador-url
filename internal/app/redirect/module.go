package redirect

import (
	app_handler "github.com/golang-encurtador-url/internal/infra/handler"
	"github.com/golang-encurtador-url/internal/infra/memrepository"
	"go.uber.org/fx"
)

func newService(repository *memrepository.MemoryRepository, stats chan string) *service {
	return NewService(repository, stats)
}

func newHandler(service *service) app_handler.HandlerResult {
	return app_handler.HandlerResult{
		Handler: NewHandler(service),
	}
}

var Module = fx.Provide(newService, newHandler)
