package chat

import (
	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/gorilla/websocket"
)

var users = []models.User{}

type userRepository interface {
	GetUserById(id models.UserId) (models.User, error)
}

type Service struct {
	repository userRepository
}

func NewService(repository userRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s Service) UsersInChat() []models.User {
	return users
}

func (s Service) Reader(conn *websocket.Conn, userId models.UserId) error {

	user, err := s.repository.GetUserById(userId)
	if err != nil {
		return err
	}

	deleteUser(userId)
	users = append(users, user)
	defer func() {
		deleteUser(userId)
	}()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return err
		}

		if err := conn.WriteMessage(messageType, p); err != nil {
			return err
		}
	}
}

func deleteUser(userId models.UserId){
	for i := range users {
		if users[i].Id == userId {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}
}
