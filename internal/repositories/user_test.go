package repositories

import (
	"database/sql"
	"log"
	"regexp"
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

func TestGetUserByIdV2(t *testing.T) {
	db, mock := NewMock()
	repo := NewPostgreUserRepository(db)

	query := regexp.QuoteMeta("SELECT * FROM users WHERE id=$1")
	rows := sqlmock.NewRows([]string{"id", "name", "email"}).AddRow(1, user.Name, user.Password)
	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

	user, err := repo.GetUserById(1)

	assert.NotNil(t, user)
	assert.NoError(t, err)
}

func TestGetUserByNameV2(t *testing.T) {
	db, mock := NewMock()
	repo := NewPostgreUserRepository(db)

	query := regexp.QuoteMeta("SELECT * FROM users WHERE name=$1")
	rows := sqlmock.NewRows([]string{"id", "name", "email"}).AddRow(1, user.Name, user.Password)
	mock.ExpectQuery(query).WithArgs(user.Name).WillReturnRows(rows)

	user, err := repo.GetUserByName(user.Name)

	assert.NotNil(t, user)
	assert.NoError(t, err)
}

func TestCheckUserExistsV2(t *testing.T) {
	db, mock := NewMock()
	repo := NewPostgreUserRepository(db)

	query := regexp.QuoteMeta("SELECT id FROM users WHERE name=$1")
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectQuery(query).WithArgs(user.Name).WillReturnRows(rows)

	ok, err := repo.CheckUserExists(user.Name)

	assert.True(t, ok)
	assert.NoError(t, err)
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
