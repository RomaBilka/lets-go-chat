package repositories

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}


func createUser(t *testing.T) models.UserId {
	id, err := testUserRepository.CreateUser(models.User{Name: user.Name, Password: user.Password})

	assert.NotNil(t, id)
	assert.NoError(t, err)
	return id
}

func TestCreateUser(t *testing.T) {
	createUser(t)
}

func TestGetUserById(t *testing.T) {
	id := createUser(t)

	u, err := testUserRepository.GetUserById(id)
	assert.Equal(t, u.Id, id)
	assert.NoError(t, err)
}

func TestGetUserByName(t *testing.T) {
	createUser(t)

	u, err := testUserRepository.GetUserByName(user.Name)

	assert.Equal(t, u.Name, user.Name)
	assert.NoError(t, err)
}

func TestCheckUserExistsTrue(t *testing.T) {
	createUser(t)

	ok, err := testUserRepository.CheckUserExists(user.Name)

	assert.True(t, ok)
	assert.NoError(t, err)
}

func TestCheckUserExistsFalse(t *testing.T) {
	createUser(t)

	ok, err := testUserRepository.CheckUserExists("")

	assert.False(t, ok)
	assert.NoError(t, err)
}
