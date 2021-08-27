package infra

import (
	"github.com/golang-encurtador-url/internal/infra/handler"
	"github.com/golang-encurtador-url/internal/infra/memrepository"
	"github.com/golang-encurtador-url/internal/infra/server"
	"go.uber.org/fx"
)

var Module = fx.Options(
	handler.Module,
	memrepository.Module,
	server.Module,
)
