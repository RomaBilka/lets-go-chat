package chat

import (
	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/RomaBiliak/lets-go-chat/internal/repositories"
)

type Chat struct {
	broadcast         chan []byte
	messageRepository *repositories.MessageRepository
	users             map[models.UserId]*userInChat
}

func NewChat(messageRepository *repositories.MessageRepository) *Chat {
	return &Chat{
		broadcast:         make(chan []byte),
		messageRepository: messageRepository,
		users:             make(map[models.UserId]*userInChat),
	}
}

func (h *Chat) Run() {
	for {
		select {
		case message := <-h.broadcast:
			for _, user := range h.users {
				select {
				case user.send <- message:
				}
			}
		}
	}
}
