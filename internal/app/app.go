package app

import (
	"context"
	"github.com/hanfkrokete/go-service/internal/config"
	httpserver "github.com/hanfkrokete/go-service/internal/http"
)

type App struct {
	HTTPServer *httpserver.Server
}

func New(cfg config.Config) *App {
	return &App{
		HTTPServer: httpserver.New(cfg.Port),
	}
}

func (a *App) Run(ctx context.Context) error {
	return a.HTTPServer.Run()
}

func (a *App) Shutdown(ctx context.Context) error {
	return a.HTTPServer.Shutdown(ctx)
}
