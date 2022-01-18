package chat

import (
	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/gorilla/websocket"
)

type userInChat struct {
	user models.User
	conn *websocket.Conn
	hub  *Hub
	send chan []byte
}

func (h *Hub) AddUserToChat(user models.User, connect *websocket.Conn) error {
	u := &userInChat{user, connect, h, make(chan []byte)}
	if activeUser, ok := h.users[user.Id]; ok {
		activeUser.conn.Close()
	}
	h.users[user.Id] = u
	messages, err := u.hub.messageRepository.GetMessages()
	if err != nil {
		return err
	}

	for _, message := range messages {
		err := connect.WriteMessage(websocket.TextMessage, []byte(message.Message))

		if err != nil {
			return err
		}
	}
	go u.Read()
	go u.Write()
	return nil
}

func (user userInChat) Read() {
	defer func() {
		user.hub.unregister <- user
		user.conn.Close()
	}()

	for {
		_, p, err := user.conn.ReadMessage()
		if err != nil {
			break
		}
		_, err = user.hub.messageRepository.CreateMessage(models.Message{UserId: user.user.Id, Message: string(p)})
		if err != nil {
			break
		}

		user.hub.broadcast <- p
	}
}
func (user userInChat) Write() {
	for {
		select {
		case messahe := <-user.send:
			user.conn.WriteMessage(websocket.TextMessage, messahe)
		}
	}
}
