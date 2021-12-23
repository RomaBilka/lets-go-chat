package services

import (
	"net/http"

	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/gorilla/websocket"
)

type chatRepository interface {
	GetUserById(id models.UserId) (models.User, error)
}

type chat interface {
	AddUserToChat(user models.User, connect *websocket.Conn) error
	UsersInChat() []models.User
}

type ChatService struct {
	repository chatRepository
	upgrader   websocket.Upgrader
	chat       chat
}

func NewChatService(repository chatRepository, chat chat) *ChatService {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	return &ChatService{
		repository: repository,
		chat:       chat,
		upgrader:   upgrader,
	}
}

func (s ChatService) Upgrader() websocket.Upgrader {
	return s.upgrader
}

func (s ChatService) GetUserById(id models.UserId) (models.User, error) {
	return s.repository.GetUserById(id)
}

func (s ChatService) AddUserToChat(user models.User, connect *websocket.Conn) error {
	return s.chat.AddUserToChat(user, connect)
}

func (s ChatService) UsersInChat() []models.User {
	return s.chat.UsersInChat()
}
