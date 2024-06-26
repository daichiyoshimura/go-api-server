// Code generated by MockGen. DO NOT EDIT.
// Source: internal/module/account/iRepo.go
//
// Generated by this command:
//
//	mockgen -source=internal/module/account/iRepo.go -destination=internal/module/account/Repo_mock.go -package=account -self_package=awsomeapp/internal/module/account
//
// Package account is a generated GoMock package.
package account

import (
	domain "awsomeapp/internal/module/account/internal/domain"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockiAccountRepository is a mock of iAccountRepository interface.
type MockiAccountRepository struct {
	ctrl     *gomock.Controller
	recorder *MockiAccountRepositoryMockRecorder
}

// MockiAccountRepositoryMockRecorder is the mock recorder for MockiAccountRepository.
type MockiAccountRepositoryMockRecorder struct {
	mock *MockiAccountRepository
}

// NewMockiAccountRepository creates a new mock instance.
func NewMockiAccountRepository(ctrl *gomock.Controller) *MockiAccountRepository {
	mock := &MockiAccountRepository{ctrl: ctrl}
	mock.recorder = &MockiAccountRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockiAccountRepository) EXPECT() *MockiAccountRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockiAccountRepository) Create(in *domain.AccountUnspecifiedDTO) (*domain.AccountDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", in)
	ret0, _ := ret[0].(*domain.AccountDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockiAccountRepositoryMockRecorder) Create(in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockiAccountRepository)(nil).Create), in)
}

// Delete mocks base method.
func (m *MockiAccountRepository) Delete(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockiAccountRepositoryMockRecorder) Delete(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockiAccountRepository)(nil).Delete), id)
}

// Get mocks base method.
func (m *MockiAccountRepository) Get(id string) (*domain.AccountDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*domain.AccountDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockiAccountRepositoryMockRecorder) Get(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockiAccountRepository)(nil).Get), id)
}

// Update mocks base method.
func (m *MockiAccountRepository) Update(in *domain.AccountDTO) (*domain.AccountDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", in)
	ret0, _ := ret[0].(*domain.AccountDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockiAccountRepositoryMockRecorder) Update(in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockiAccountRepository)(nil).Update), in)
}
