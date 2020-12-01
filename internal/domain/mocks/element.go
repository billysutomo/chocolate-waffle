// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/billysutomo/chocolate-waffle/internal/domain (interfaces: ElementRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	domain "github.com/billysutomo/chocolate-waffle/internal/domain"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockElementRepository is a mock of ElementRepository interface
type MockElementRepository struct {
	ctrl     *gomock.Controller
	recorder *MockElementRepositoryMockRecorder
}

// MockElementRepositoryMockRecorder is the mock recorder for MockElementRepository
type MockElementRepositoryMockRecorder struct {
	mock *MockElementRepository
}

// NewMockElementRepository creates a new mock instance
func NewMockElementRepository(ctrl *gomock.Controller) *MockElementRepository {
	mock := &MockElementRepository{ctrl: ctrl}
	mock.recorder = &MockElementRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockElementRepository) EXPECT() *MockElementRepositoryMockRecorder {
	return m.recorder
}

// CreateElement mocks base method
func (m *MockElementRepository) CreateElement(arg0 context.Context, arg1 domain.ElementModel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateElement", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateElement indicates an expected call of CreateElement
func (mr *MockElementRepositoryMockRecorder) CreateElement(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateElement", reflect.TypeOf((*MockElementRepository)(nil).CreateElement), arg0, arg1)
}
