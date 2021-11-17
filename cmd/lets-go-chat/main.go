//Home task (retraining program)
package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/RomaBiliak/lets-go-chat/internal/auth"
	authHttp "github.com/RomaBiliak/lets-go-chat/internal/auth/http"
	"github.com/RomaBiliak/lets-go-chat/internal/chat"
	"github.com/RomaBiliak/lets-go-chat/internal/repositories/postgres"
	"github.com/RomaBiliak/lets-go-chat/internal/user"
	userHttp "github.com/RomaBiliak/lets-go-chat/internal/user/http"
	database "github.com/RomaBiliak/lets-go-chat/pkg/database/postgres"
	httpServer "github.com/RomaBiliak/lets-go-chat/pkg/http"
	"github.com/RomaBiliak/lets-go-chat/pkg/middleware"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	pgUser, ok:=os.LookupEnv("PG_USER")
	if !ok {
		panic(errors.New("PG_USER is empty"))
	}
	pgPassword, ok:=os.LookupEnv("PG_PASSWORD")
	if !ok {
		panic(errors.New("PG_PASSWORD is empty"))
	}
	pgDatabase, ok:=os.LookupEnv("PG_DATABASE")
	if !ok {
		panic(errors.New("PG_DATABASE is empty"))
	}


	dbConfig := database.Config{
		pgUser,
		pgPassword,
		pgDatabase,
	}
	db:= database.Run(dbConfig)
	defer db.Close()



	userRepository := postgres.NewPostgreUserRepository(db)

	mux := http.NewServeMux()

	userService := user.NewService(userRepository)
	uHttp := userHttp.NewUserHttp(userService)
	mux.Handle("/v1/user", middleware.LogRequest(middleware.LogError(middleware.LogPanic(uHttp.CreateUser))))

	authService := auth.NewService(userRepository)
	aHttp := authHttp.NewAuthHttp(authService)
	mux.Handle("/v1/user/login", middleware.LogRequest(middleware.LogError(middleware.LogPanic(aHttp.Login))))

	mux.Handle("/", middleware.LogRequest(serveHome))


	hub := chat.NewHub()
	go hub.Run()

	mux.Handle("/ws", middleware.LogRequest(
		func(w http.ResponseWriter, r *http.Request) {
			chat.Chat(hub, w, r)
	}))

	mux.Handle("/ws1", middleware.LogRequest(chat.ChatTest))


	//http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	//	chat.Chat(hub, w, r)
	//})

	//chatService:= chat.NewChat()
	//mux.Handle("/ws", func(w http.ResponseWriter, r *http.Request) {
	//	serveWs(s, w, r)
	//})


	httpServer.Start(":8080", mux)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "home.html")
}
