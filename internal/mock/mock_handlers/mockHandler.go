// Code generated by MockGen. DO NOT EDIT.
// Source: userHandler.go

// Package mock_handlers is a generated GoMock package.
package mock_handlers

import (
	reflect "reflect"

	models "github.com/RomaBiliak/lets-go-chat/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockuserService is a mock of userService interface.
type MockuserService struct {
	ctrl     *gomock.Controller
	recorder *MockuserServiceMockRecorder
}

// MockuserServiceMockRecorder is the mock recorder for MockuserService.
type MockuserServiceMockRecorder struct {
	mock *MockuserService
}

// NewMockuserService creates a new mock instance.
func NewMockuserService(ctrl *gomock.Controller) *MockuserService {
	mock := &MockuserService{ctrl: ctrl}
	mock.recorder = &MockuserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockuserService) EXPECT() *MockuserServiceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockuserService) CreateUser(user models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockuserServiceMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockuserService)(nil).CreateUser), user)
}
