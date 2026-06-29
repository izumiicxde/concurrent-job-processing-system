package middlewares

import "net/http"

type LoggingResponseWriter struct {
	http.ResponseWriter
	status int
}

func (l *LoggingResponseWriter) WriteHeader(status int) {
	l.status = status
	l.ResponseWriter.WriteHeader(status)
}
