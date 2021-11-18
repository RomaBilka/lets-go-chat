package services

import (
	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/gorilla/websocket"
)

var users = []models.User{}

type chatRepository interface {
	GetUserById(id models.UserId) (models.User, error)
}

type ChatService struct {
	repository chatRepository
}

func NewChatService(repository chatRepository) *ChatService {
	return &ChatService{
		repository: repository,
	}
}

func (s ChatService) UsersInChat() []models.User {
	return users
}

func (s ChatService) Reader(conn *websocket.Conn, userId models.UserId) error {

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
