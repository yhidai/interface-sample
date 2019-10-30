// Code generated by MockGen. DO NOT EDIT.
// Source: ./sample3/main3.go

// Package mock_main is a generated GoMock package.
package main

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserTableInterface is a mock of UserTableInterface interface
type MockUserTableInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserTableInterfaceMockRecorder
}

// MockUserTableInterfaceMockRecorder is the mock recorder for MockUserTableInterface
type MockUserTableInterfaceMockRecorder struct {
	mock *MockUserTableInterface
}

// NewMockUserTableInterface creates a new mock instance
func NewMockUserTableInterface(ctrl *gomock.Controller) *MockUserTableInterface {
	mock := &MockUserTableInterface{ctrl: ctrl}
	mock.recorder = &MockUserTableInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserTableInterface) EXPECT() *MockUserTableInterfaceMockRecorder {
	return m.recorder
}

// Count mocks base method
func (m *MockUserTableInterface) Count() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count")
	ret0, _ := ret[0].(int)
	return ret0
}

// Count indicates an expected call of Count
func (mr *MockUserTableInterfaceMockRecorder) Count() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockUserTableInterface)(nil).Count))
}
