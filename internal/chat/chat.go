package chat

import (
	"net/http"

	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/RomaBiliak/lets-go-chat/pkg/response"
	"github.com/RomaBiliak/lets-go-chat/pkg/token"
	"github.com/gorilla/websocket"
)

type chatService interface {
	Reader(conn *websocket.Conn, id models.UserId) error
	UsersInChat()[]models.User
}

type ChatHTTP struct {
	chatService chatService
}

func NewChatHttp(chatService chatService) *ChatHTTP {
	return &ChatHTTP{chatService: chatService}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *ChatHTTP) Chat(w http.ResponseWriter, r *http.Request) {
	t, _:= token.ParseToken(r.URL.Query().Get("token"))

	useId, err:=t.UserId()
	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

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

func (h *ChatHTTP) UserInChat(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		response.WriteERROR(w, http.StatusMethodNotAllowed, nil)
		return
	}

	type CreateUserResponse struct {
		Id       uint64 `json:"id"`
		UserName string `json:"userName"`
	}

	users := []CreateUserResponse{}

	for _, user := range h.chatService.UsersInChat(){
		users = append(users, CreateUserResponse{
			uint64(user.Id),
			user.Name,
		})
	}

	response.WriteJSON(w, http.StatusCreated, users)
}

