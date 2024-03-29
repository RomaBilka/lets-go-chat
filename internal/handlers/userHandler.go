package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/RomaBiliak/lets-go-chat/pkg/response"
)

type userService interface {
	CreateUser(user models.User) (models.User, error)
}

type userHTTP struct {
	userService userService
}

func NewUserHttp(userService userService) *userHTTP {
	return &userHTTP{userService: userService}
}

type CreateUserRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func (r *CreateUserRequest) Validate() bool {
	return len(r.UserName) > 4 && len(r.Password) > 8
}

type CreateUserResponse struct {
	Id       uint64 `json:"id"`
	UserName string `json:"userName"`
}

//CreateUser creates new user, status code of 201
func (h *userHTTP) CreateUser(w http.ResponseWriter, r *http.Request) {
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

	response.WriteJSON(w, http.StatusCreated, CreateUserResponse{Id: uint64(newUser.Id), UserName: newUser.Name})
}
