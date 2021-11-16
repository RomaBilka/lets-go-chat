package middleware

import (
	"net/http"

	"github.com/RomaBiliak/lets-go-chat/pkg/log"
)

func LogRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		logStdout := log.NewLogStdout("Request")
		logStdout.AddMessage("method", r.Method)
		logStdout.AddMessage("path", r.URL.EscapedPath())
		logStdout.AddMessage("ip", r.Header.Get("X-Real-Ip"))
		logStdout.AddMessage("User-Agent", r.Header.Get("User-Agent"))
		logStdout.Print()

		next.ServeHTTP(w, r)
	}
}
