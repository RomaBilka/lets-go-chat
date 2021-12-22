package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/RomaBiliak/lets-go-chat/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUserAPIV2(t *testing.T) {
	testCases := []struct {
		name          string
		body          CreateUserRequest
		method        string
		create        func(mockUserService *mocks.UserService)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "Ok",
			body:   userTest,
			method: http.MethodPost,
			create: func(mockUserService *mocks.UserService) {
				//mockUserService.On("CreateUser", models.User{Name: userTest.UserName, Password: userTest.Password}).Return(models.User{Id: 1, Name: userTest.UserName, Password: userTest.Password}, nil)
				mockUserService.On("CreateUser", mock.Anything).Return(models.User{Id: 1, Name: userTest.UserName, Password: userTest.Password}, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusCreated, recorder.Code)

				userResponse := &CreateUserResponse{}
				err := json.NewDecoder(recorder.Body).Decode(userResponse)
				assert.NoError(t, err)
				assert.Equal(t, userTest.UserName, userResponse.UserName)
				assert.Equal(t, uint64(1), userResponse.Id)
			},
		},
		{
			name:   "StatusMethodNotAllowed",
			body:   userTest,
			method: http.MethodGet,
			create: func(mockUserService *mocks.UserService) {
				//mockUserService.On("CreateUser", models.User{}).Times(0)
				mockUserService.On("CreateUser", mock.Anything).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusMethodNotAllowed, recorder.Code)
			},
		},
	}

	for i := range testCases {
		testCase := testCases[i]
		t.Run(testCase.name, func(t *testing.T) {

			mockUserService := new(mocks.UserService)
			testCase.create(mockUserService)
			uHttp := NewUserHttp(mockUserService)

			b := new(bytes.Buffer)
			err := json.NewEncoder(b).Encode(testCase.body)
			assert.NoError(t, err)

			req, err := http.NewRequest(testCase.method, "/v1/user", b)
			assert.NoError(t, err)

			recorder := httptest.NewRecorder()
			handler := http.HandlerFunc(uHttp.CreateUser)
			handler.ServeHTTP(recorder, req)

			testCase.checkResponse(recorder)
		})
	}
}
