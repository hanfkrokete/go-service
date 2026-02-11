package middleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	status int
}

func (lw *loggingResponseWriter) WriteHeader(code int) {
	lw.status = code
	lw.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		//TODO :достань или сгенерируй request_id
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}
		//TODO :установи X-Request-Id в ответ
		w.Header().Set("X-Request-ID", requestID)

		lw := &loggingResponseWriter{
			ResponseWriter: w,
			status:         http.StatusOK, //TODO: подумать что поставить по умолчанию
		}

		next.ServeHTTP(lw, r)

		//TODO: если статус так и не выставлен, то 200
		if lw.status == 0 {
			lw.status = http.StatusOK
		}
		//TODO залогировать method, path, status, duration_ms, request_id
		slog.Info("handled request",
			"method", r.Method,
			"path", r.URL.Path,
			"status", lw.status,
			"duration_ms", time.Since(start).Milliseconds(),
			"request_id", requestID,
		)

		_ = start
	})
}
