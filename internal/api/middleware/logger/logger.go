package logger

import (
	"log/slog"
	"net/http"
	"time"
)

func LoggerMiddleware(log *slog.Logger) func(next http.Handler) http.Handler {
	log.Info("logger middleware enabled")
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			ww := &responseWriterWithStatus{w, http.StatusOK}
			defer func() {
				status := ww.Status()
				elapsed := time.Since(startTime)
				log.Info("Request status", slog.String("Path", r.URL.Path), slog.String("Method", r.Method), slog.Int("status_code", status), slog.Duration("elapsed_time", elapsed))
			}()
			next.ServeHTTP(ww, r)
		}

		return http.HandlerFunc(fn)
	}
}

type responseWriterWithStatus struct {
	http.ResponseWriter
	status int
}

func (r *responseWriterWithStatus) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

func (r *responseWriterWithStatus) Status() int {
	return r.status
}
