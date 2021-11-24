package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/stretchr/testify/assert"
)

var user = CreateUserRequest{
	"test_name1",
	"test_password",
}

func createUser(t *testing.T, user CreateUserRequest) *httptest.ResponseRecorder {

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(user)

	req, err := http.NewRequest(http.MethodPost, "/v1/user", b)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(uHttp.CreateUser)
	handler.ServeHTTP(recorder, req)
	return recorder
}

func TestCreateUser(t *testing.T) {
	defer func() {
		err := truncateUsers()
		assert.NoError(t, err)
	}()

	recorder := createUser(t, user)
	assert.Equal(t, http.StatusCreated, recorder.Code)

	userResponse := &CreateUserResponse{}
	err := json.NewDecoder(recorder.Body).Decode(userResponse)

	assert.NoError(t, err)

	userInDb, err := testUserRepository.GetUserById(models.UserId(userResponse.Id))

	assert.NoError(t, err)
	assert.Equal(t, user.UserName, userInDb.Name)
}

func TestCreateSecondUser(t *testing.T) {
	defer func() {
		err := truncateUsers()
		assert.NoError(t, err)
	}()

	createUser(t, user)
	recorder := createUser(t, user)

	errorResponse := &errorResponse{}
	err := json.NewDecoder(recorder.Body).Decode(errorResponse)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.NotEmpty(t, errorResponse)
	assert.Equal(t, "User with that name already exists", errorResponse.Error)
}

func TestCreateUserShortNamePassword(t *testing.T) {
	recorder := createUser(t, CreateUserRequest{})

	errorResponse := &errorResponse{}
	err := json.NewDecoder(recorder.Body).Decode(errorResponse)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.NotEmpty(t, errorResponse)
	assert.Equal(t, "Bad request, short user name or password", errorResponse.Error)
}
