package repositories

import (
	"os"
	"testing"
	"errors"

	"github.com/RomaBiliak/lets-go-chat/pkg/database/postgres"
	"github.com/joho/godotenv"
)


var testUserRepository * UserRepository

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
	testUserRepository = NewPostgreUserRepository(db)

	os.Exit(m.Run())
}
