package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/RomaBiliak/lets-go-chat/pkg/middleware"
	"github.com/RomaBiliak/lets-go-chat/pkg/token"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func newWSServer(t *testing.T, h http.Handler, tokenString string) (*httptest.Server, *websocket.Conn) {
	server := httptest.NewServer(h)

	wsURL, err := url.Parse(server.URL)
	assert.NoError(t, err)
	wsURL.Scheme = "ws"

	ws, _, err := websocket.DefaultDialer.Dial(wsURL.String()+"?token="+tokenString, nil)
	assert.NoError(t, err)

	return server, ws
}

func TestChat(t *testing.T) {
	tokenString, err := token.CreateToken(uint64(1))
	assert.NoError(t, err)

	server, ws := newWSServer(t, middleware.Authentication(cHttp.Chat), tokenString)
	defer server.Close()
	defer ws.Close()

	message := "test message"
	assert.NoError(t, err)

	err = ws.WriteMessage(websocket.BinaryMessage, []byte(message))
	assert.NoError(t, err)

	_, readMessage, err := ws.ReadMessage()
	assert.NoError(t, err)

	assert.Equal(t, message, string(readMessage))
}

//func TestUsersInChat(t *testing.T) {
//	defer func() {
//		err := truncateUsers()
//		assert.NoError(t, err)
//	}()
//	user, err := uHttp.userService.CreateUser(models.User{Name: login.UserName, Password: login.Password})
//	assert.NoError(t, err)
//
//	tokenString, err := token.CreateToken(uint64(user.Id))
//	assert.NoError(t, err)
//
//	server, ws := newWSServer(t, middleware.Authentication(cHttp.Chat), tokenString)
//	defer server.Close()
//	defer ws.Close()
//
//
//
//	req, err := http.NewRequest(http.MethodGet, "/v1/user/active", nil)
//	assert.NoError(t, err)
//
//	recorder := httptest.NewRecorder()
//	handler := http.HandlerFunc(uHttp.CreateUser)
//	handler.ServeHTTP(recorder, req)
//
//
//
//
//	//resp, err := http.Get(server.URL+"/v1/user/active")
//	//assert.NoError(t, err)
//	//defer resp.Body.Close()
//	//
//	//
//	//b, err := httputil.DumpResponse(recorder.Body, true)
//	//assert.NoError(t, err)
//	//
//	//usersResponse := make([]CreateUserResponse,0)
//	//json.Unmarshal(recorder.Body, &usersResponse)
//
//	fmt.Println(recorder.Body)
//}
