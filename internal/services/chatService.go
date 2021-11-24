package services

import (
	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/gorilla/websocket"
)

var users = make(map[models.UserId]models.User)

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

func (s ChatService) UsersInChat() map[models.UserId]models.User {
	return users
}

func (s ChatService) Reader(conn *websocket.Conn, userId models.UserId) error {
	user, err := s.repository.GetUserById(userId)
	if err != nil {
		return err
	}

	users[userId] = user
	defer func() {
		delete(users, userId)
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
