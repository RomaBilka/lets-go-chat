package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/RomaBiliak/lets-go-chat/internal/user"
	"github.com/RomaBiliak/lets-go-chat/pkg/response"
)

type UserHTTP struct {
	userService *user.Service
}

func NewUserHttp(userService *user.Service) *UserHTTP {
	return &UserHTTP{userService: userService}
}

type CreateUserRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func (r *CreateUserRequest) Validate() bool {
	return len(r.UserName) > 4 && len(r.Password) > 8
}

type CreateUserResponse struct {
	Id       uint64 `json:"is"`
	UserName string `json:"userName"`
}

//CreateUser creates new user, status code of 201
func (h *UserHTTP) CreateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		response.WriteERROR(w, http.StatusMethodNotAllowed, nil)
		return
	}

	userRequest := &CreateUserRequest{}

	err := json.NewDecoder(r.Body).Decode(userRequest)
	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	ok := userRequest.Validate()
	if !ok {
		response.WriteERROR(w, http.StatusBadRequest, fmt.Errorf("%s", "Bad request, short user name or password"))
		return
	}

	userModel := models.User{
		Name:     userRequest.UserName,
		Password: userRequest.Password,
	}

	newUser, err := h.userService.CreateUser(userModel)
	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	response.WriteJSON(w, http.StatusCreated, CreateUserResponse{Id: newUser.Id, UserName: newUser.Name})
}
