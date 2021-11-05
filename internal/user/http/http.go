package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/RomaBiliak/lets-go-chat/internal/user"
	"github.com/RomaBiliak/lets-go-chat/pkg/response"
)

type  UserHTTP struct {
	userService *user.Service
}

func NewUserHttp(userService *user.Service) *UserHTTP{
	return &UserHTTP{userService: userService}
}

type CreateUserRequesr struct {
	UserName string
	Password string
}

func (r *CreateUserRequesr) Validate() bool {
	return len(r.UserName) > 4 && len(r.Password) > 8
}

type CreateUserResponse struct {
	Id uint64
	UserName string
}

//CreateUser creates new user, status code of 201
func (h *UserHTTP)CreateUser(w http.ResponseWriter, r *http.Request) {
	userRequesr := &CreateUserRequesr{}

	err := json.NewDecoder(r.Body).Decode(userRequesr)
	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	ok := userRequesr.Validate()
	if !ok {
		response.WriteERROR(w, http.StatusBadRequest, fmt.Errorf("%s","Bad request, short username or password"))
		return
	}

	userModel := models.User{
		Name: userRequesr.UserName,
		Password: userRequesr.Password,
	}
	newUser, err := h.userService.CreateUser(userModel)

	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	response.WriteJSON(w, http.StatusCreated, CreateUserResponse{Id: newUser.Id, UserName: newUser.Name})
}