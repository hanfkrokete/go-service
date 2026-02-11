package httpserver

import (
	"github.com/hanfkrokete/go-service/internal/http/handler"

	"net/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handler.Health)

	return mux
}
