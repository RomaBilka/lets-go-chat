package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var login = loginRequest{
	"test_name",
	"test_password",
}

type responseToken struct {
	url string `json:"url"`
}

func loginTest(t *testing.T, login loginRequest) *httptest.ResponseRecorder {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(login)

	req, err := http.NewRequest(http.MethodPost, "/v1/user/login", b)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(aHttp.Login)
	handler.ServeHTTP(recorder, req)
	return recorder
}

func TestLogin(t *testing.T) {
	defer func() {
		err := truncateUsers()
		assert.NoError(t, err)
	}()

	createTestUser(t)

	recorder := loginTest(t, login)
	responseToken := &responseToken{}

	err := json.NewDecoder(recorder.Body).Decode(responseToken)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, recorder.Code)
}

func TestLoginWrongPassword(t *testing.T) {
	defer func() {
		err := truncateUsers()
		assert.NoError(t, err)
	}()

	createTestUser(t)

	l := login
	l.Password = "error"
	recorder := loginTest(t, l)

	errorResponse := &errorResponse{}
	err := json.NewDecoder(recorder.Body).Decode(errorResponse)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, "Invalid password", errorResponse.Error)
}

func TestLoginUserFound(t *testing.T) {
	defer func() {
		err := truncateUsers()
		assert.NoError(t, err)
	}()

	createTestUser(t)

	l := login
	l.UserName = "error"
	recorder := loginTest(t, l)

	errorResponse := &errorResponse{}
	err := json.NewDecoder(recorder.Body).Decode(errorResponse)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, "Bad request, user not found", errorResponse.Error)
}

func TestLoginEmptyLoginData(t *testing.T) {
	defer func() {
		err := truncateUsers()
		assert.NoError(t, err)
	}()

	createTestUser(t)

	l := login
	l.UserName = ""
	l.Password = ""
	recorder := loginTest(t, l)

	errorResponse := &errorResponse{}
	err := json.NewDecoder(recorder.Body).Decode(errorResponse)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, "Bad request, empty user name or password", errorResponse.Error)
}
