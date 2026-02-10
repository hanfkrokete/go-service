package httpserver

import (
	"context"
	"log/slog"
	"net/http"
	"time"
)

type Server struct {
	srv *http.Server
}

func New(port string) *Server{
	mux := http.ServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte('{"status": "ok"}'))
	})

	srv := &http.Server{
		Addr: ":" + port,
		Handler: mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	return &Server{srv: srv}
}

func (s *Server) Run() error{
	slog.info("starting http servre", "addr", s.srv.Addr)
	return s.srv.LinstenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error{
	slog.info("shutting down http server")
	return s.srv.Shutdown(ctx)
}
