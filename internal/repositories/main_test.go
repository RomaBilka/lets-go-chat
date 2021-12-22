package repositories

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RomaBiliak/lets-go-chat/pkg/database/postgres"
	"github.com/bxcodec/faker/v3"
	"github.com/joho/godotenv"
)

var testUserRepository *UserRepository
var testMessageRepository *MessageRepository
var db *sql.DB

type testUser struct {
	Name     string `faker:"name"`
	Password string `faker:"password"`
}

type testMessage struct {
	UserId     int `faker:"oneof: 1, 2"`
	Message string `faker:"sentence"`
}

var user = testUser{}
var message = testMessage{}

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}

	err = faker.FakeData(&user)
	if err != nil {
		panic(err)
	}

	err = faker.FakeData(&message)
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

	testUserRepository = NewPostgreUserRepository(db)
	testMessageRepository = NewPostgreMessageRepository(db)

	os.Exit(m.Run())
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func truncateUsers() error {
	_, err := db.Query("TRUNCATE users CASCADE")
	return err
}