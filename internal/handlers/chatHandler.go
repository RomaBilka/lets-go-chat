package handlers

import (
	"net/http"

	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/RomaBiliak/lets-go-chat/pkg/response"
	"github.com/RomaBiliak/lets-go-chat/pkg/token"
	"github.com/gorilla/websocket"
)

type chatService interface {
	GetUserById(id models.UserId) (models.User, error)
	AddUserToChat(user models.User, connect *websocket.Conn) error
	UsersInChat() []models.User
	Upgrader() websocket.Upgrader
}

type chatHTTP struct {
	chatService chatService
}

func NewChatHttp(chatService chatService) *chatHTTP {
	return &chatHTTP{chatService: chatService}
}

func (h *chatHTTP) Chat(w http.ResponseWriter, r *http.Request) {
	t, _ := token.ParseToken(r.URL.Query().Get("token"))

	useId, err := t.UserId()
	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.chatService.GetUserById(models.UserId(useId))
	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	upgrader := h.chatService.Upgrader()
	connect, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	err = h.chatService.AddUserToChat(user, connect)
	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}
}

func (h *chatHTTP) UsersInChat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.WriteERROR(w, http.StatusMethodNotAllowed, nil)
		return
	}

	var users []CreateUserResponse

	for _, user := range h.chatService.UsersInChat() {
		users = append(users, CreateUserResponse{
			uint64(user.Id),
			user.Name,
		})
	}

	response.WriteJSON(w, http.StatusCreated, users)
}
