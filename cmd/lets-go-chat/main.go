//Home task (retraining program)
package main

import (
	"net/http"
	"os"
	"errors"

	"github.com/RomaBiliak/lets-go-chat/internal/auth"
	authHttp "github.com/RomaBiliak/lets-go-chat/internal/auth/http"
	"github.com/RomaBiliak/lets-go-chat/internal/repositories/postgres"
	"github.com/RomaBiliak/lets-go-chat/internal/user"
	userHttp "github.com/RomaBiliak/lets-go-chat/internal/user/http"
	httpServer "github.com/RomaBiliak/lets-go-chat/pkg/http"
	"github.com/joho/godotenv"
	 database "github.com/RomaBiliak/lets-go-chat/pkg/database/postgres"
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

	userService := user.NewService(userRepository)
	uHttp := userHttp.NewUserHttp(userService)
	http.HandleFunc("/v1/user", uHttp.CreateUser)

	authService := auth.NewService(userRepository)
	aHttp := authHttp.NewAuthHttp(authService)
	http.HandleFunc("/v1/user/login", aHttp.Login)

	httpServer.Start(":8080")
}
