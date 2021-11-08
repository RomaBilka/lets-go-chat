package user

import (
	"fmt"
	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/RomaBiliak/lets-go-chat/pkg/hasher"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

//CreateUser creates a new user and adds it to users
func (s *Service) CreateUser(user models.User) (*models.User, error) {

	_, ok := models.Users[user.Name]
	if ok {
		return &user, fmt.Errorf("%s", "User with that name already exists")
	}

	hashPassword, err := hasher.HashPassword(user.Password)
	if err != nil {
		return &user, err
	}

	models.UserId++
	user.Id = models.UserId
	user.Password = hashPassword
	models.Users[user.Name] = user

	return &user, nil
}
