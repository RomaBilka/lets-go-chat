package handlers

import (
	"database/sql"
	"errors"
	"os"
	"testing"

	"github.com/RomaBiliak/lets-go-chat/internal/repositories"
	"github.com/RomaBiliak/lets-go-chat/internal/services"
	"github.com/RomaBiliak/lets-go-chat/pkg/database/postgres"
	"github.com/joho/godotenv"
)

var testUserRepository *repositories.UserRepository
var uHttp *userHTTP
var aHttp *authHTTP
var cHttp *chatHTTP

type errorResponse struct {
	Error string `json:"error"`
}

var db *sql.DB

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}

	pgUser, ok := os.LookupEnv("PG_USER")
	if !ok {
		panic(errors.New("PG_USER is empty"))
	}
	pgPassword, ok := os.LookupEnv("PG_PASSWORD")
	if !ok {
		panic(errors.New("PG_PASSWORD is empty"))
	}
	pgDatabase, ok := os.LookupEnv("PG_TEST_DATABASE")
	if !ok {
		panic(errors.New("PG_TEST_DATABASE is empty"))
	}

	dbConfig := postgres.Config{
		pgUser,
		pgPassword,
		pgDatabase,
	}
	db = postgres.Run(dbConfig)

	defer db.Close()
	testUserRepository = repositories.NewPostgreUserRepository(db)

	userService := services.NewUserService(testUserRepository)
	uHttp = NewUserHttp(userService)

	authService := services.NewAuthService(testUserRepository)
	aHttp = NewAuthHttp(authService)

	cService := services.NewChatService(testUserRepository)
	cHttp = NewChatHttp(cService)

	os.Exit(m.Run())
}

func truncateUsers() error {
	_, err := db.Query("TRUNCATE users")
	return err
}
