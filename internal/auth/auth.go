package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RomaBiliak/lets-go-chat/pkg/hasher"
	"github.com/RomaBiliak/lets-go-chat/pkg/token"
	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/RomaBiliak/lets-go-chat/pkg/responses"
)

type loginResponse struct {
	Url string `json:"url"`
}

//Login returns token to join chat
func Login(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	ok := user.Validate()
	if !ok {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("%s", "Bad request, short username or password"))
		return
	}

	userInDb, ok := models.GetUserByName(user.Name)
	if !ok {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("%s", "Bad request, user not found"))
		return
	}

	ok = hasher.CheckPasswordHash(user.Password, userInDb.Password)
	if !ok {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("%s", "Invalid password"))
		return
	}

	tokenString, err := token.CreateToken(userInDb.Id)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusCreated, loginResponse{ Url: tokenString })
}
