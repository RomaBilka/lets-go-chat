package middleware

import (
	"net/http"

	"github.com/RomaBiliak/lets-go-chat/pkg/response"
	"github.com/RomaBiliak/lets-go-chat/pkg/token"
)

func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryToken := r.URL.Query().Get("token")

		if _, ok := token.RevokedTokens[queryToken]; ok {
			response.WriteERROR(w, http.StatusUnauthorized, nil)
			return
		}
		token.RevokedTokens[queryToken] = true

		t, err := token.ParseToken(queryToken)
		if err != nil {
			response.WriteERROR(w, http.StatusUnauthorized, nil)
			return
		}

		expired, err := t.IsExpired()
		if err != nil {
			response.WriteERROR(w, http.StatusUnauthorized, nil)
			return
		}

		if t.Valid && !expired {
			next(w, r)
		}

		return
	}
}
