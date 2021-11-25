package middleware

import (
	"net/http"
)

func LogRequest(log logInterface, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Init("Request")
		log.AddMessage("method", r.Method)
		log.AddMessage("path", r.URL.EscapedPath())
		log.AddMessage("ip", r.Header.Get("X-Real-Ip"))
		log.AddMessage("User-Agent", r.Header.Get("User-Agent"))
		log.Print()

		next.ServeHTTP(w, r)
	}
}
