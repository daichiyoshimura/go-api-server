// Code generated by MockGen. DO NOT EDIT.
// Source: internal/module/account/iAccountUsecase.go
//
// Generated by this command:
//
//	mockgen -source internal/module/account/iAccountUsecase.go -destination internal/module/account/mock/usecase_mock.go -package mock
//
// Package mock is a generated GoMock package.
package mock

import (
	account "awsomeapp/internal/module/account"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockIAccountUsecase is a mock of IAccountUsecase interface.
type MockIAccountUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockIAccountUsecaseMockRecorder
}

// MockIAccountUsecaseMockRecorder is the mock recorder for MockIAccountUsecase.
type MockIAccountUsecaseMockRecorder struct {
	mock *MockIAccountUsecase
}

// NewMockIAccountUsecase creates a new mock instance.
func NewMockIAccountUsecase(ctrl *gomock.Controller) *MockIAccountUsecase {
	mock := &MockIAccountUsecase{ctrl: ctrl}
	mock.recorder = &MockIAccountUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAccountUsecase) EXPECT() *MockIAccountUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIAccountUsecase) Create(in *account.AccountCreateInput) (*account.AccountCreateOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", in)
	ret0, _ := ret[0].(*account.AccountCreateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIAccountUsecaseMockRecorder) Create(in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIAccountUsecase)(nil).Create), in)
}

// Delete mocks base method.
func (m *MockIAccountUsecase) Delete(in *account.AccountDeleteInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", in)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIAccountUsecaseMockRecorder) Delete(in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIAccountUsecase)(nil).Delete), in)
}

// Get mocks base method.
func (m *MockIAccountUsecase) Get(in *account.AccountGetInput) (*account.AccountGetOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", in)
	ret0, _ := ret[0].(*account.AccountGetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIAccountUsecaseMockRecorder) Get(in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIAccountUsecase)(nil).Get), in)
}

// Update mocks base method.
func (m *MockIAccountUsecase) Update(in *account.AccountUpdateInput) (*account.AccountUpdateOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", in)
	ret0, _ := ret[0].(*account.AccountUpdateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIAccountUsecaseMockRecorder) Update(in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIAccountUsecase)(nil).Update), in)
}