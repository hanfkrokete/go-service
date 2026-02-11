package app

import (
	"context"

	"github.com/hanfkrokete/go-service/internal/config"
	httpserver "github.com/hanfkrokete/go-service/internal/http"
	"github.com/hanfkrokete/go-service/internal/http/middleware"
)

type App struct {
	HTTPServer *httpserver.Server
}

func New(cfg config.Config) *App {
	router := httpserver.NewRouter()
	routerWithMiddleware := middleware.Logging(router)
	server := httpserver.New(cfg.Port, routerWithMiddleware)

	return &App{
		HTTPServer: server,
	}
}

func (a *App) Run(ctx context.Context) error {
	return a.HTTPServer.Run()
}

func (a *App) Shutdown(ctx context.Context) error {
	return a.HTTPServer.Shutdown(ctx)
}
