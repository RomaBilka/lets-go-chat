package token

import (
	"fmt"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Token struct {
	*jwt.Token
}

//CreateToken return jwt token
func CreateToken(userId uint64) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func ParseToken(tokenString string) (Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	return Token{token}, err
}

func (t *Token) IsExpired() (bool, error) {

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return false, nil
	}

	timestamp, err := strconv.ParseInt(fmt.Sprintf("%.0f", claims["exp"]), 10, 32)
	if err != nil {
		return false, err
	}

	return time.Now().Unix() > timestamp, nil
}

func (t *Token) UserId() (uint64, error) {

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return 0, nil
	}

	userId, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 64)
	if err != nil {
		return 0, err
	}

	return userId, nil
}
