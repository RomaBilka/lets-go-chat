package models

import (
	"github.com/RomaBiliak/lets-go-chat/pkg/hasher"
)

type User struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	Password string `json:"-"`
}

//users inmemory store
var users = []User{}

var id uint64 = 0

func (user *User) Validate() bool {
	return len(user.Name) < 4 || len(user.Password) < 8
}

//CreateUser creates a new user and adds it to users
func (user *User) CreateUser() error {

	hashPassword, err := hasher.HashPassword(user.Password)
	if err != nil {
		return err
	}

	id++
	user.Id = id
	user.Password = hashPassword
	users = append(users, *user)

	return nil
}

func GetUserByName(name string) (User, bool) {
	for _, user := range users {
		if user.Name == name {
			return user, true
		}
	}
	return User{}, false
}
