package user

import (
	"fmt"
	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/RomaBiliak/lets-go-chat/pkg/hasher"
)

type UserRepository interface {
	GetUserByName(name string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
}

type Service struct{
	repository UserRepository
}

func NewService(repository UserRepository) *Service {
	return &Service{
		repository: repository,
	}
}

//CreateUser creates a new user and adds it to users
func (s *Service) CreateUser(user models.User) (models.User, error) {
	userInDb, err := s.repository.GetUserByName(user.Name)
	if err != nil {
		return user, err
	}

	if userInDb.Id > 0 {
		return user, fmt.Errorf("%s", "User with that name already exists")
	}

	hashPassword, err := hasher.HashPassword(user.Password)
	if err != nil {
		return user, err
	}

	user.Password = hashPassword
	return s.repository.CreateUser(user)
}
