package services

import (
	"fmt"

	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/RomaBiliak/lets-go-chat/pkg/hasher"
	"github.com/RomaBiliak/lets-go-chat/pkg/token"
)

type authRepository interface {
	GetUserByName(name string) (models.User, error)
}

type AuthService struct {
	repository authRepository
}

func NewAuthService(repository authRepository) *AuthService {
	return &AuthService{
		repository: repository,
	}
}

//Login returns token
func (s AuthService) Login(userName, password string) (string, error) {

	user, err := s.repository.GetUserByName(userName)
	if user.Id == 0 {
		return "", fmt.Errorf("%s", "Bad request, user not found")
	}

	ok := hasher.CheckPasswordHash(password, user.Password)
	if !ok {
		return "", fmt.Errorf("%s", "Invalid password")
	}

	tokenString, err := token.CreateToken(uint64(user.Id))
	if err != nil {
		return "", err
	}

	return tokenString, nil

}
