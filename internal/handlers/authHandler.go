package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/RomaBiliak/lets-go-chat/pkg/response"
)

type authService interface {
	Login(userName, password string) (string, error)
}

type AuthHTTP struct {
	authService authService
}

func NewAuthHttp(authService authService) *AuthHTTP {
	return &AuthHTTP{authService: authService}
}

type LoginRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func (r *LoginRequest) Validate() bool {
	return len(r.UserName) > 0 && len(r.Password) > 0
}

type LoginResponse struct {
	Url string `json:"url"`
}

//Login returns token to join chat, status code of 201
func (h *AuthHTTP) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.WriteERROR(w, http.StatusMethodNotAllowed, nil)
		return
	}

	loginRequest := &LoginRequest{}

	err := json.NewDecoder(r.Body).Decode(loginRequest)
	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	ok := loginRequest.Validate()
	if !ok {
		response.WriteERROR(w, http.StatusBadRequest, fmt.Errorf("%s", "Bad request, empty user name or password"))
		return
	}

	token, err := h.authService.Login(loginRequest.UserName, loginRequest.Password)
	if err != nil {
		response.WriteERROR(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("X-Rate-Limit", "999999")
	w.Header().Set("X-Expires-After", time.Now().Add(time.Hour*1).UTC().String())

	response.WriteJSON(w, http.StatusCreated, LoginResponse{Url: token})
}
