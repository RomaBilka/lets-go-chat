package services

import (
	"fmt"

	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/RomaBiliak/lets-go-chat/pkg/hasher"
)

type userRepository interface {
	GetUserByName(name string) (models.User, error)
	CheckUserExists(name string) (bool, error)
	GetUserById(id models.UserId) (models.User, error)
	CreateUser(user models.User) (models.UserId, error)
}

type UserService struct {
	repository userRepository
}

func NewUserService(repository userRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

//CreateUser creates a new user and adds it to users
func (s *UserService) CreateUser(user models.User) (models.User, error) {
	exists, err := s.repository.CheckUserExists(user.Name)
	if err != nil {
		return user, err
	}

	if exists {
		return models.User{}, fmt.Errorf("%s", "User with that name already exists")
	}

	hashPassword, err := hasher.HashPassword(user.Password)
	if err != nil {
		return user, err
	}

	user.Password = hashPassword

	id, err := s.repository.CreateUser(user)
	if err != nil {
		return models.User{}, err
	}

	user, err = s.repository.GetUserById(id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
