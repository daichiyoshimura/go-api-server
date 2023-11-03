// Code generated by MockGen. DO NOT EDIT.
// Source: internal/module/account/internal/domain/iAccountRepo.go
//
// Generated by this command:
//
//	mockgen -source=internal/module/account/internal/domain/iAccountRepo.go -destination=internal/module/account/internal/domain/mock/iAccountRepo_mock.go -package=mock
//
// Package mock is a generated GoMock package.
package mock

import (
	domain "awsomeapp/internal/module/account/internal/domain"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockIAccountRepository is a mock of IAccountRepository interface.
type MockIAccountRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIAccountRepositoryMockRecorder
}

// MockIAccountRepositoryMockRecorder is the mock recorder for MockIAccountRepository.
type MockIAccountRepositoryMockRecorder struct {
	mock *MockIAccountRepository
}

// NewMockIAccountRepository creates a new mock instance.
func NewMockIAccountRepository(ctrl *gomock.Controller) *MockIAccountRepository {
	mock := &MockIAccountRepository{ctrl: ctrl}
	mock.recorder = &MockIAccountRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAccountRepository) EXPECT() *MockIAccountRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIAccountRepository) Create(in *domain.AccountDTO) (*domain.AccountDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", in)
	ret0, _ := ret[0].(*domain.AccountDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIAccountRepositoryMockRecorder) Create(in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIAccountRepository)(nil).Create), in)
}

// Delete mocks base method.
func (m *MockIAccountRepository) Delete(id domain.AccountID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIAccountRepositoryMockRecorder) Delete(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIAccountRepository)(nil).Delete), id)
}

// Get mocks base method.
func (m *MockIAccountRepository) Get(id domain.AccountID) (*domain.AccountDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*domain.AccountDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIAccountRepositoryMockRecorder) Get(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIAccountRepository)(nil).Get), id)
}

// Update mocks base method.
func (m *MockIAccountRepository) Update(in *domain.AccountDTO) (*domain.AccountDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", in)
	ret0, _ := ret[0].(*domain.AccountDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIAccountRepositoryMockRecorder) Update(in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIAccountRepository)(nil).Update), in)
}