package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/RomaBiliak/lets-go-chat/internal/mock/mock_chat_handlers"
	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/RomaBiliak/lets-go-chat/pkg/middleware"
	"github.com/RomaBiliak/lets-go-chat/pkg/token"
	"github.com/golang/mock/gomock"
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
	defer func() {
		err := truncateUsers()
		assert.NoError(t, err)
		err = truncateMessages()
		assert.NoError(t, err)
	}()
	userId := createTestUser(t)

	tokenString, err := token.CreateToken(uint64(userId))
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

func TestUsersInChat(t *testing.T) {
	var users []models.User
	users = append(users, models.User{Name: userTest.UserName, Password: userTest.Password})

	testCases := []struct {
		name          string
		body          CreateUserRequest
		method        string
		create        func(mockChatService *mock_chat_handlers.MockchatService)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "Ok",
			method: http.MethodGet,
			create: func(mockChatService *mock_chat_handlers.MockchatService) {
				mockChatService.EXPECT().UsersInChat().Return(users)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusCreated, recorder.Code)

				usersResponse := &[]CreateUserResponse{}
				err := json.NewDecoder(recorder.Body).Decode(usersResponse)

				assert.NoError(t, err)
				assert.Equal(t, http.StatusCreated, recorder.Code)
				assert.NotEmpty(t, usersResponse)
			},
		},
		{
			name:   "StatusMethodNotAllowed",
			method: http.MethodPost,
			create: func(mockChatService *mock_chat_handlers.MockchatService) {},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusMethodNotAllowed, recorder.Code)
			},
		},
	}

	for i := range testCases {
		testCase := testCases[i]
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockChatService := mock_chat_handlers.NewMockchatService(ctrl)
			testCase.create(mockChatService)
			cHttp := NewChatHttp(mockChatService)

			req, err := http.NewRequest(testCase.method, "/v1/user/active", nil)
			assert.NoError(t, err)

			recorder := httptest.NewRecorder()
			handler := http.HandlerFunc(cHttp.UsersInChat)
			handler.ServeHTTP(recorder, req)

			testCase.checkResponse(recorder)
		})
	}
}
