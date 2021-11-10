package auth

import (
	"fmt"

	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/RomaBiliak/lets-go-chat/pkg/hasher"
	"github.com/RomaBiliak/lets-go-chat/pkg/token"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

//Login returns token
func (s Service) Login(userName, password string) (string, error) {

	user, ok := models.Users[userName]
	if !ok {
		return "", fmt.Errorf("%s", "Bad request, user not found")
	}

	ok = hasher.CheckPasswordHash(password, user.Password)
	if !ok {
		return "", fmt.Errorf("%s", "Invalid password")
	}

	tokenString, err := token.CreateToken(user.Id)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
