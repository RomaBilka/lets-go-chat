package repositories

import (
	"errors"
	"os"
	"testing"

	"github.com/RomaBiliak/lets-go-chat/pkg/database/postgres"
	"github.com/bxcodec/faker/v3"
	"github.com/joho/godotenv"
)

var testUserRepository *UserRepository

type testUser struct {
	Name string  `faker:"name"`
	Password string `faker:"password"`
}

var user = testUser{}


func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}

	err = faker.FakeData(&user)
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
	db := postgres.Run(dbConfig)
	defer db.Close()
	testUserRepository = NewPostgreUserRepository(db)

	os.Exit(m.Run())
}
