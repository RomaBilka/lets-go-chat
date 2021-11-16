package middleware

import (
	"fmt"
	"net/http"

	"github.com/RomaBiliak/lets-go-chat/pkg/log"
)

func LogPanic(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if message:=recover(); message != nil {
				logStdout := log.NewLogStdout("Panic")
				logStdout.AddMessage("message", fmt.Sprintf("%s", message))
				logStdout.Print()
			}
		}()

		next.ServeHTTP(w, r)
	}
}
