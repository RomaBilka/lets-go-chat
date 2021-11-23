package middleware

import (
	"net/http"
	"strconv"
)

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true

	return
}

func LogError(log logInterface, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		wrapped := wrapResponseWriter(w)

		next.ServeHTTP(wrapped, r)

		if wrapped.status >= 400 {
			log.Init("Error")
			log.AddMessage("status", strconv.Itoa(wrapped.status))
			log.AddMessage("method", r.Method)
			log.AddMessage("path", r.URL.EscapedPath())
			log.Print()
		}
	}
}
