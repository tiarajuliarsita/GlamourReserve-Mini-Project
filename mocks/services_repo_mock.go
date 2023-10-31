// Code generated by MockGen. DO NOT EDIT.
// Source: glamour_reserve/features/repositories (interfaces: ServiceRepoInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	core "glamour_reserve/entity/core"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockServiceRepoInterface is a mock of ServiceRepoInterface interface.
type MockServiceRepoInterface struct {
	ctrl     *gomock.Controller
	recorder *MockServiceRepoInterfaceMockRecorder
}

// MockServiceRepoInterfaceMockRecorder is the mock recorder for MockServiceRepoInterface.
type MockServiceRepoInterfaceMockRecorder struct {
	mock *MockServiceRepoInterface
}

// NewMockServiceRepoInterface creates a new mock instance.
func NewMockServiceRepoInterface(ctrl *gomock.Controller) *MockServiceRepoInterface {
	mock := &MockServiceRepoInterface{ctrl: ctrl}
	mock.recorder = &MockServiceRepoInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceRepoInterface) EXPECT() *MockServiceRepoInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockServiceRepoInterface) Create(arg0 core.ServiceCore) (core.ServiceCore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(core.ServiceCore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockServiceRepoInterfaceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockServiceRepoInterface)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockServiceRepoInterface) Delete(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceRepoInterfaceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServiceRepoInterface)(nil).Delete), arg0)
}

// FindAll mocks base method.
func (m *MockServiceRepoInterface) FindAll(arg0 string, arg1, arg2 int) ([]core.ServiceCore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", arg0, arg1, arg2)
	ret0, _ := ret[0].([]core.ServiceCore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockServiceRepoInterfaceMockRecorder) FindAll(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockServiceRepoInterface)(nil).FindAll), arg0, arg1, arg2)
}

// FindById mocks base method.
func (m *MockServiceRepoInterface) FindById(arg0 string) (core.ServiceCore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0)
	ret0, _ := ret[0].(core.ServiceCore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockServiceRepoInterfaceMockRecorder) FindById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockServiceRepoInterface)(nil).FindById), arg0)
}

// Update mocks base method.
func (m *MockServiceRepoInterface) Update(arg0 string, arg1 core.ServiceCore) (core.ServiceCore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(core.ServiceCore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockServiceRepoInterfaceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockServiceRepoInterface)(nil).Update), arg0, arg1)
}
