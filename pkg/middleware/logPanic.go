package middleware

import (
	"fmt"
	"net/http"
)

func LogPanic(log logInterface, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if message:=recover(); message != nil {
				log.Init("Panic")
				log.AddMessage("message", fmt.Sprintf("%s", message))
				log.Print()
			}
		}()

		next.ServeHTTP(w, r)
	}
}
