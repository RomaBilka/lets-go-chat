//Home task (retraining program)
package main

import (
	"fmt"
	"net/http"
	"os"

	httpServer "github.com/RomaBiliak/lets-go-chat/pkg/http"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	pgUser, _ := os.LookupEnv("PG_USER")
	/*if !ok {
		panic(errors.New("PG_USER is empty"))
	}
	pgPassword, ok := os.LookupEnv("PG_PASSWORD")
	if !ok {
		panic(errors.New("PG_PASSWORD is empty"))
	}
	pgDatabase, ok := os.LookupEnv("PG_DATABASE")
	if !ok {
		panic(errors.New("PG_DATABASE is empty"))
	}*/
/*
	dbConfig := postgres.Config{
		pgUser,
		pgPassword,
		pgDatabase,
	}
	db := postgres.Run(dbConfig)
	defer db.Close()

	userRepository := repositories.NewPostgreUserRepository(db)
*/
	mux := http.NewServeMux()
/*
	logStdout := log.NewLogStdout()

	userService := services.NewUserService(userRepository)
	uHttp := handlers.NewUserHttp(userService)
	mux.Handle("/v1/user", middleware.LogRequest(logStdout, middleware.LogError(logStdout, middleware.LogPanic(logStdout, uHttp.CreateUser))))

	authService := services.NewAuthService(userRepository)
	aHttp := handlers.NewAuthHttp(authService)
	mux.Handle("/v1/user/login", middleware.LogRequest(logStdout, middleware.LogError(logStdout, middleware.LogPanic(logStdout, aHttp.Login))))

	chatService := services.NewChatService(userRepository)
	cHttp := handlers.NewChatHttp(chatService)
	mux.Handle("/v1/ws", middleware.LogRequest(logStdout, middleware.Authentication(cHttp.Chat)))

	mux.Handle("/v1/user/active", middleware.LogRequest(logStdout, middleware.LogError(logStdout, middleware.LogPanic(logStdout, cHttp.UsersInChat))))*/
	mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, pgUser)
		fmt.Fprintf(w, "Test 3")
	})

	httpServer.Start(":8080", mux)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "home.html")
}
