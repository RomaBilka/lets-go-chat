//Home task (retraining program)
package main

import (
	"github.com/RomaBiliak/lets-go-chat/internal/user"
	"github.com/RomaBiliak/lets-go-chat/internal/auth"
	userHttp "github.com/RomaBiliak/lets-go-chat/internal/user/http"
	authHttp "github.com/RomaBiliak/lets-go-chat/internal/auth/http"
	httpServer "github.com/RomaBiliak/lets-go-chat/pkg/http"
	"net/http"
)


func main() {

	userService := user.NewService()
	uHttp:= userHttp.NewUserHttp(userService)
	http.HandleFunc("/v1/user", uHttp.CreateUser)

	authService := auth.NewService()
	aHttp:= authHttp.NewAuthHttp(authService)
	http.HandleFunc("/v1/user/login", aHttp.Login)

	httpServer.Start(":8080")
}