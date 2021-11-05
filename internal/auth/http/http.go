package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RomaBiliak/lets-go-chat/internal/auth"
	"github.com/RomaBiliak/lets-go-chat/pkg/response"
)

type AuthHTTP struct {
	authService *auth.Service
}

func NewAuthHttp(authService *auth.Service) *AuthHTTP {
	return &AuthHTTP{authService: authService}
}

type LoginRequest struct {
	UserName string
	Password string
}

func (r *LoginRequest) Validate() bool {
	return len(r.UserName) > 4 && len(r.Password) > 8
}

type LoginResponse struct {
	Url string
}

//Login returns token to join chat, status code of 201
func (h *AuthHTTP) Login(w http.ResponseWriter, r *http.Request) {
	loginRequest := &LoginRequest{}

	err := json.NewDecoder(r.Body).Decode(loginRequest)
	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	ok := loginRequest.Validate()
	if !ok {
		response.WriteERROR(w, http.StatusBadRequest, fmt.Errorf("%s", "Bad request, short username or password"))
		return
	}

	token, err := h.authService.GetToken(loginRequest.UserName, loginRequest.Password)
	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	response.WriteJSON(w, http.StatusCreated, LoginResponse{Url: token})
}
