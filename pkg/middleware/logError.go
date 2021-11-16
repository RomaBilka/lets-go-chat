package middleware

import (
	"net/http"
	"strconv"

	"github.com/RomaBiliak/lets-go-chat/pkg/log"
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

func LogError(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		wrapped := wrapResponseWriter(w)

		next.ServeHTTP(wrapped, r)

		if wrapped.status >= 400 {
			logStdout := log.NewLogStdout("Error")
			logStdout.AddMessage("status", strconv.Itoa(wrapped.status))
			logStdout.AddMessage("method", r.Method)
			logStdout.AddMessage("path", r.URL.EscapedPath())
			logStdout.Print()
		}
	}
}
