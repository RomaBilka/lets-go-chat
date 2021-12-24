package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RomaBiliak/lets-go-chat/internal/mock/mock_chat_handlers"
	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

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
