package chat

import (
	"fmt"

	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/gorilla/websocket"
)

type userInChat struct {
	user models.User
	conn *websocket.Conn
	chat *Chat
	send chan []byte
}

func (c *Chat) AddUserToChat(user models.User, connect *websocket.Conn) error {
	u := &userInChat{user, connect, c, make(chan []byte)}
	if activeUser, ok:=c.users[user.Id]; ok{
		activeUser.conn.Close()
	}
	c.users[user.Id] = u
	messages, err := u.chat.messageRepository.GetMessages()
	if err != nil {
		return err
	}

	for _, message := range messages {
		connect.WriteMessage(websocket.TextMessage, []byte(message.Message))
	}
	go u.Read()
	go u.Write()
	return nil
}

func (c *Chat) UserInChat() []models.User {
	var users []models.User
	for _, user := range c.users {
		users=append(users, user.user)
	}
	fmt.Println(c.users )
	return users
}

func (user userInChat) Read() {
	defer func(){
		user.conn.Close()
		delete(user.chat.users, user.user.Id)
	}()
	for {
		_, p, err := user.conn.ReadMessage()
		if err != nil {
			break
		}
		_, err = user.chat.messageRepository.CreateMessage(models.Message{UserId: user.user.Id, Message: string(p)})
		if err != nil {
			break
		}
		user.chat.broadcast <- p
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
