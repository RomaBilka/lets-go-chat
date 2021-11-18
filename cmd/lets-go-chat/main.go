//Home task (retraining program)
package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/RomaBiliak/lets-go-chat/internal/handlers"
	"github.com/RomaBiliak/lets-go-chat/internal/repositories"
	"github.com/RomaBiliak/lets-go-chat/internal/services"
	"github.com/RomaBiliak/lets-go-chat/pkg/database/postgres"
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


	dbConfig := postgres.Config{
		pgUser,
		pgPassword,
		pgDatabase,
	}
	db:= postgres.Run(dbConfig)
	defer db.Close()

	userRepository := repositories.NewPostgreUserRepository(db)

	mux := http.NewServeMux()

	userService := services.NewUserService(userRepository)
	uHttp := handlers.NewUserHttp(userService)
	mux.Handle("/v1/user", middleware.LogRequest(middleware.LogError(middleware.LogPanic(uHttp.CreateUser))))

	authService := services.NewAuthService(userRepository)
	aHttp := handlers.NewAuthHttp(authService)
	mux.Handle("/v1/user/login", middleware.LogRequest(middleware.LogError(middleware.LogPanic(aHttp.Login))))

	chatService := services.NewService(userRepository)
	cHttp := handlers.NewChatHttp(chatService)
	mux.Handle("/v1/ws", middleware.LogRequest(middleware.Authentication(cHttp.Chat)))

	mux.Handle("/v1/user/active", middleware.LogRequest(middleware.LogError(middleware.LogPanic(cHttp.UserInChat))))

	httpServer.Start(":8080", mux)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "home.html")
}
