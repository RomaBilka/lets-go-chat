package middleware

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/RomaBiliak/lets-go-chat/pkg/token"
	"github.com/stretchr/testify/assert"
)

func TestLogRequest(t *testing.T) {
	server := httptest.NewServer(LogRequest(testLog, func(w http.ResponseWriter, r *http.Request) {
	}))
	defer server.Close()

	_, err := http.Get(server.URL)

	assert.NoError(t, err)
	assert.Equal(t, "Request", testLog.GetName())
	assert.Equal(t, http.MethodGet, testLog.GetMessage("method"))

}

func TestLogError(t *testing.T) {
	server := httptest.NewServer(LogError(testLog, func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "", http.StatusBadRequest)
	}))
	defer server.Close()

	_, err := http.Get(server.URL)

	assert.NoError(t, err)
	assert.Equal(t, "Error", testLog.GetName())
	assert.Equal(t, http.MethodGet, testLog.GetMessage("method"))
	assert.Equal(t, strconv.Itoa(http.StatusBadRequest), testLog.GetMessage("status"))
}

func TestLogPanic(t *testing.T) {
	testPanicStr := "Test Panic"
	server := httptest.NewServer(LogPanic(testLog, func(w http.ResponseWriter, r *http.Request) {
		panic(testPanicStr)
	}))
	defer server.Close()

	_, err := http.Get(server.URL)

	assert.NoError(t, err)
	assert.Equal(t, "Panic", testLog.GetName())
	assert.Equal(t, testPanicStr, testLog.GetMessage("message"))
}

func TestAuthentication(t *testing.T) {
	server := httptest.NewServer(Authentication(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	newToken, err := token.CreateToken(uint64(1))
	assert.NoError(t, err)

	res, err := http.Get(server.URL + "?token=" + newToken)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestAuthenticationFail(t *testing.T) {
	server := httptest.NewServer(Authentication(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	res, err := http.Get(server.URL + "?token=")

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
}