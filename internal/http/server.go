package httpserver

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"
)

type Server struct {
	srv *http.Server
}

func New(port string, h http.Handler) *Server {
	return &Server{
		srv: &http.Server{
			Addr:              ":" + port,
			Handler:           h,
			ReadHeaderTimeout: 5 * time.Second,
			IdleTimeout:       60 * time.Second,
		},
	}
}

func (s *Server) Run() error {
	slog.Info("starting http servre", "addr", s.srv.Addr)
	err := s.srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	slog.Info("shutting down http server")
	return s.srv.Shutdown(ctx)
}
