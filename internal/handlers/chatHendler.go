package handlers

import (
	"net/http"

	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/RomaBiliak/lets-go-chat/pkg/response"
	"github.com/RomaBiliak/lets-go-chat/pkg/token"
	"github.com/gorilla/websocket"
)

type chatService interface {
	Reader(conn *websocket.Conn, id models.UserId) error
	UsersInChat() map[models.UserId]models.User
	SetUser(models.User)
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

	upgrader := h.chatService.Upgrader()
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	err = h.chatService.Reader(ws, models.UserId(useId))

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

	type CreateUserResponse struct {
		Id       uint64 `json:"id"`
		UserName string `json:"userName"`
	}

	activeUsers := h.chatService.UsersInChat()
	users := make([]CreateUserResponse, len(activeUsers))

	for _, user := range activeUsers {
		users = append(users, CreateUserResponse{
			uint64(user.Id),
			user.Name,
		})
	}

	response.WriteJSON(w, http.StatusCreated, users)
}
