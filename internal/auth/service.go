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

//GetToken returns token
func (s Service) GetToken(userName, password string) (string, error) {

	for _, user := range models.Users {
		if user.Name == userName {
			ok := hasher.CheckPasswordHash(password, user.Password)
			if !ok {
				return "", fmt.Errorf("%s", "Invalid password")
			}

			tokenString, err := token.CreateToken(user.Id)
			if err != nil {
				return "", err
			}

			return tokenString, nil
		}
	}

	return "", fmt.Errorf("%s", "Bad request, user not found")
}
