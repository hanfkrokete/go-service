package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hanfkrokete/go-service/internal/app"
	"github.com/hanfkrokete/go-service/internal/config"
)

func main() {
	slog.Info("App is running")

	cfg := config.MustLoad()
	application := app.New(cfg)

	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)

	defer stop()

	go func() {
		if err := application.Run(ctx); err != nil {
			slog.Error("failed to run application", "error", err)
			stop()
		}
	}()

	<-ctx.Done()
	slog.Info("shutdown signal recived")

	shutDownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := application.Shutdown(shutDownCtx); err != nil {
		slog.Error("failed to shutdown application", "error", err)
	}

	slog.Info("app stopped")

}
