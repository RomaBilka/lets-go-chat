package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

//CreateToken return jwt token
func CreateToken(userId uint64) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("54daA5cSDH56"))
}
