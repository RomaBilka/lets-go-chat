package services

import (
	"net/http"

	"github.com/RomaBiliak/lets-go-chat/internal/chat"
	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/gorilla/websocket"
)

type chatRepository interface {
	GetUserById(id models.UserId) (models.User, error)
}

//type chat interface {
//	GetUserById(id models.UserId) (models.User, error)
//}

type ChatService struct {
	repository chatRepository
	upgrader   websocket.Upgrader
	chat       *chat.Chat
}

func NewChatService(repository chatRepository, chat *chat.Chat) *ChatService {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	return &ChatService{
		repository: repository,
		chat: chat,
		upgrader:   upgrader,
	}
}

func (s ChatService) Upgrader() websocket.Upgrader {
	return s.upgrader
}

func (s ChatService) GetUserById(id models.UserId) (models.User, error) {
	return s.repository.GetUserById(id)
}

func (s ChatService) AddUserToChat(user models.User, connect *websocket.Conn) {
	s.chat.AddUserToChat(user, connect)
}

//func (s ChatService) UsersInChat() map[models.UserId]models.User {
//	return users
//}
//
//func (s ChatService) SetUser(user models.User) {
//	users[user.Id] = user
//}

func (s ChatService) Reader(conn *websocket.Conn, userId models.UserId) error {
	//	user, err := s.repository.GetUserById(userId)
	//	if err != nil {
	//		return err
	//	}
	//	s.SetUser(user)

	defer func() {
		//		delete(users, userId)
		conn.Close()
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
