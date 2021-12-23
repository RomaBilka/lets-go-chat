package handlers

import (
	"database/sql"
	"errors"
	"os"
	"testing"

	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/RomaBiliak/lets-go-chat/internal/repositories"
	"github.com/RomaBiliak/lets-go-chat/internal/services"
	"github.com/RomaBiliak/lets-go-chat/pkg/database/postgres"
	"github.com/RomaBiliak/lets-go-chat/pkg/hasher"
	"github.com/bxcodec/faker/v3"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var testUserRepository *repositories.UserRepository
var testMessageRepository *repositories.MessageRepository
var uHttp *userHTTP
var aHttp *authHTTP
var cHttp *chatHTTP

type errorResponse struct {
	Error string `json:"error"`
}

var db *sql.DB

var userTest CreateUserRequest

var login loginRequest

type testData struct {
	UserName string `faker:"name"`
	Password string `faker:"password"`
}

var test = testData{}

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}

	err = faker.FakeData(&test)
	if err != nil {
		panic(err)
	}

	userTest = CreateUserRequest{test.UserName, test.Password}
	login = loginRequest{test.UserName, test.Password}

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
	pgHost, ok := os.LookupEnv("PG_HOST")
	if !ok {
		panic(errors.New("Db Host is empty"))
	}

	dbConfig := postgres.Config{
		pgUser,
		pgPassword,
		pgDatabase,
		pgHost,
	}
	db = postgres.Run(dbConfig)

	defer db.Close()
	testUserRepository = repositories.NewPostgreUserRepository(db)
	testMessageRepository = repositories.NewPostgreMessageRepository(db)

	userService := services.NewUserService(testUserRepository)
	uHttp = NewUserHttp(userService)

	authService := services.NewAuthService(testUserRepository)
	aHttp = NewAuthHttp(authService)

	//caht := chat.NewChat(testMessageRepository)
	//go caht.Run()
	//cService := services.NewChatService(testUserRepository, caht)
	//cHttp = NewChatHttp(cService)

	os.Exit(m.Run())
}

func truncateUsers() error {
	_, err := db.Query("TRUNCATE users CASCADE")
	return err
}

func createTestUser(t *testing.T) models.UserId {
	hashPassword, err := hasher.HashPassword(login.Password)
	assert.NoError(t, err)

	userId, err := testUserRepository.CreateUser(models.User{Name: login.UserName, Password: hashPassword})
	assert.NoError(t, err)
	return userId
}
