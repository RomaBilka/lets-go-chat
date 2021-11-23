package handlers

import (
	"errors"
	"os"
	"testing"

	"github.com/RomaBiliak/lets-go-chat/internal/repositories"
	"github.com/RomaBiliak/lets-go-chat/internal/services"
	"github.com/RomaBiliak/lets-go-chat/pkg/database/postgres"
	"github.com/joho/godotenv"
)


var testUserRepository * repositories.UserRepository
var uHttp *UserHTTP
type  errorResponse struct {
	Error string `json:"error"`
}

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
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
	pgDatabase, ok:=os.LookupEnv("PG_TEST_DATABASE")
	if !ok {
		panic(errors.New("PG_TEST_DATABASE is empty"))
	}


	dbConfig := postgres.Config{
		pgUser,
		pgPassword,
		pgDatabase,
	}
	db:= postgres.Run(dbConfig)
	defer db.Close()
	testUserRepository = repositories.NewPostgreUserRepository(db)

	userService := services.NewUserService(testUserRepository)
	uHttp = NewUserHttp(userService)

	os.Exit(m.Run())
}
