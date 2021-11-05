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

	_, ok := getUserByName(user.Name)
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
	models.Users = append(models.Users, user)

	return &user, nil
}

func getUserByName(name string) (models.User, bool) {
	for _, user := range models.Users {
		if user.Name == name {
			return user, true
		}
	}
	return models.User{}, false
}
