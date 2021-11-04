package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/RomaBiliak/lets-go-chat/pkg/responses"
)

//CreateUser creates new user, status code of 201
func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	ok := user.Validate()
	if !ok {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("%s","Bad request, short username or password"))
		return
	}

	_, ok = models.GetUserByName(user.Name)
	if ok {
		responses.ERROR(w, http.StatusBadRequest, fmt.Errorf("%s", "User with that name already exists"))
		return
	}

	_ = user.CreateUser()
	user.Password = "";

	responses.JSON(w, http.StatusCreated, user)
}
