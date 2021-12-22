package chat

import "github.com/RomaBiliak/lets-go-chat/internal/repositories"

type Chat struct {
	broadcast chan []byte
	messageRepository  *repositories.MessageRepository
	users map[*userInChat]bool
}

func NewChat(messageRepository *repositories.MessageRepository) *Chat {
	return &Chat{
		broadcast:  make(chan []byte),
		messageRepository:  messageRepository,
		users:    make(map[*userInChat]bool),
	}
}

func (h *Chat) Run() {
	for {
		select {
		case message := <-h.broadcast:
			for user := range h.users {
				select {
				case user.send <- message:
				}
			}
		}
	}
}
