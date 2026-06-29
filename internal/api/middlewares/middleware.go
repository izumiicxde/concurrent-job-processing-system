package middlewares

import (
	"concurrent-job-processing-system/internal/logger"
	"net/http"
	"time"
)

func Logging(logger *logger.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wrappedWriter := &LoggingResponseWriter{status: 200, ResponseWriter: w}
			start := time.Now()
			next.ServeHTTP(wrappedWriter, r)
			logger.HTTPRequest(r.Method, r.URL.Path, wrappedWriter.status, time.Since(start), r.RemoteAddr)
		})
	}
}
