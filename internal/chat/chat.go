package chat

import (
	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/RomaBiliak/lets-go-chat/internal/repositories"
)

type Hub struct {
	broadcast         chan []byte
	messageRepository *repositories.MessageRepository
	unregister        chan userInChat
	users             map[models.UserId]*userInChat
}

func NewChat(messageRepository *repositories.MessageRepository) *Hub {
	return &Hub{
		broadcast:         make(chan []byte),
		messageRepository: messageRepository,
		unregister:        make(chan userInChat),
		users:             make(map[models.UserId]*userInChat),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case user := <-h.unregister:
			if _, ok := h.users[user.user.Id]; ok {
				delete(user.hub.users, user.user.Id)
				close(user.send)
			}
		case message := <-h.broadcast:
			for _, user := range h.users {
				select {
				case user.send <- message:
				default:
					close(user.send)
					delete(user.hub.users, user.user.Id)
				}
			}
		}
	}
}

func (c *Hub) UsersInChat() []models.User {
	var users []models.User
	for _, user := range c.users {
		users = append(users, user.user)
	}
	return users
}
