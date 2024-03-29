// Code generated by MockGen. DO NOT EDIT.
// Source: chatHandler.go

// Package mock_handlers is a generated GoMock package.
package mock_chat_handlers

import (
	reflect "reflect"

	models "github.com/RomaBiliak/lets-go-chat/internal/models"
	gomock "github.com/golang/mock/gomock"
	websocket "github.com/gorilla/websocket"
)

// MockchatService is a mock of chatService interface.
type MockchatService struct {
	ctrl     *gomock.Controller
	recorder *MockchatServiceMockRecorder
}

// MockchatServiceMockRecorder is the mock recorder for MockchatService.
type MockchatServiceMockRecorder struct {
	mock *MockchatService
}

// NewMockchatService creates a new mock instance.
func NewMockchatService(ctrl *gomock.Controller) *MockchatService {
	mock := &MockchatService{ctrl: ctrl}
	mock.recorder = &MockchatServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockchatService) EXPECT() *MockchatServiceMockRecorder {
	return m.recorder
}

// AddUserToChat mocks base method.
func (m *MockchatService) AddUserToChat(user models.User, connect *websocket.Conn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUserToChat", user, connect)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUserToChat indicates an expected call of AddUserToChat.
func (mr *MockchatServiceMockRecorder) AddUserToChat(user, connect interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserToChat", reflect.TypeOf((*MockchatService)(nil).AddUserToChat), user, connect)
}

// GetUserById mocks base method.
func (m *MockchatService) GetUserById(id models.UserId) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", id)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockchatServiceMockRecorder) GetUserById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockchatService)(nil).GetUserById), id)
}

// Upgrader mocks base method.
func (m *MockchatService) Upgrader() websocket.Upgrader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upgrader")
	ret0, _ := ret[0].(websocket.Upgrader)
	return ret0
}

// Upgrader indicates an expected call of Upgrader.
func (mr *MockchatServiceMockRecorder) Upgrader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upgrader", reflect.TypeOf((*MockchatService)(nil).Upgrader))
}

// UsersInChat mocks base method.
func (m *MockchatService) UsersInChat() []models.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UsersInChat")
	ret0, _ := ret[0].([]models.User)
	return ret0
}

// UsersInChat indicates an expected call of UsersInChat.
func (mr *MockchatServiceMockRecorder) UsersInChat() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UsersInChat", reflect.TypeOf((*MockchatService)(nil).UsersInChat))
}
