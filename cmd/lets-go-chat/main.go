//Home task (retraining program)
package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/RomaBiliak/lets-go-chat/internal/chat"
	"github.com/RomaBiliak/lets-go-chat/internal/handlers"
	"github.com/RomaBiliak/lets-go-chat/internal/repositories"
	"github.com/RomaBiliak/lets-go-chat/internal/services"
	"github.com/RomaBiliak/lets-go-chat/pkg/database/postgres"
	httpServer "github.com/RomaBiliak/lets-go-chat/pkg/http"
	"github.com/RomaBiliak/lets-go-chat/pkg/log"
	"github.com/RomaBiliak/lets-go-chat/pkg/middleware"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	pgUser, ok := os.LookupEnv("PG_USER")
	if !ok {
		panic(errors.New("PG_USER is empty"))
	}
	pgPassword, ok := os.LookupEnv("PG_PASSWORD")
	if !ok {
		panic(errors.New("PG_PASSWORD is empty"))
	}
	pgHost, ok := os.LookupEnv("PG_HOST")
	if !ok {
		panic(errors.New("Db Host is empty"))
	}
	pgDatabase, ok := os.LookupEnv("PG_DATABASE")
	if !ok {
		panic(errors.New("PG_DATABASE is empty"))
	}

	dbConfig := postgres.Config{
		pgUser,
		pgPassword,
		pgDatabase,
		pgHost,
	}
	db := postgres.Run(dbConfig)
	defer db.Close()

	userRepository := repositories.NewPostgreUserRepository(db)
	messageRepository := repositories.NewPostgreMessageRepository(db)

	mux := http.NewServeMux()

	logStdout := log.NewLogStdout()

	userService := services.NewUserService(userRepository)
	uHttp := handlers.NewUserHttp(userService)
	mux.Handle("/v1/user", middleware.LogRequest(logStdout, middleware.LogError(logStdout, middleware.LogPanic(logStdout, uHttp.CreateUser))))

	authService := services.NewAuthService(userRepository)
	aHttp := handlers.NewAuthHttp(authService)
	mux.Handle("/v1/user/login", middleware.LogRequest(logStdout, middleware.LogError(logStdout, middleware.LogPanic(logStdout, aHttp.Login))))

	newChat := chat.NewChat(messageRepository)
	go newChat.Run()
	chatService := services.NewChatService(userRepository, newChat)
	cHttp := handlers.NewChatHttp(chatService)
	mux.Handle("/v1/ws", middleware.LogRequest(logStdout, middleware.Authentication(cHttp.Chat)))

	mux.Handle("/v1/user/active", middleware.LogRequest(logStdout, middleware.LogError(logStdout, middleware.LogPanic(logStdout, cHttp.UsersInChat))))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Index")
	})

	httpServer.Start(":8080", mux)
}
