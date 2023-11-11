// Code generated by MockGen. DO NOT EDIT.
// Source: internal/handler/iAccountUsecase.go
//
// Generated by this command:
//
//	mockgen -source=internal/handler/iAccountUsecase.go -destination=internal/handler/accountUsecase_mock.go -package=handler -self_package=awsomeapp/internal/handler
//
// Package handler is a generated GoMock package.
package handler

import (
	account "awsomeapp/internal/module/account"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockiAccountUsecase is a mock of iAccountUsecase interface.
type MockiAccountUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockiAccountUsecaseMockRecorder
}

// MockiAccountUsecaseMockRecorder is the mock recorder for MockiAccountUsecase.
type MockiAccountUsecaseMockRecorder struct {
	mock *MockiAccountUsecase
}

// NewMockiAccountUsecase creates a new mock instance.
func NewMockiAccountUsecase(ctrl *gomock.Controller) *MockiAccountUsecase {
	mock := &MockiAccountUsecase{ctrl: ctrl}
	mock.recorder = &MockiAccountUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockiAccountUsecase) EXPECT() *MockiAccountUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockiAccountUsecase) Create(in *account.AccountCreateInput) (*account.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", in)
	ret0, _ := ret[0].(*account.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockiAccountUsecaseMockRecorder) Create(in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockiAccountUsecase)(nil).Create), in)
}

// Delete mocks base method.
func (m *MockiAccountUsecase) Delete(in *account.AccountDeleteInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", in)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockiAccountUsecaseMockRecorder) Delete(in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockiAccountUsecase)(nil).Delete), in)
}

// Get mocks base method.
func (m *MockiAccountUsecase) Get(in *account.AccountGetInput) (*account.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", in)
	ret0, _ := ret[0].(*account.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockiAccountUsecaseMockRecorder) Get(in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockiAccountUsecase)(nil).Get), in)
}

// Update mocks base method.
func (m *MockiAccountUsecase) Update(in *account.AccountUpdateInput) (*account.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", in)
	ret0, _ := ret[0].(*account.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockiAccountUsecaseMockRecorder) Update(in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockiAccountUsecase)(nil).Update), in)
}
