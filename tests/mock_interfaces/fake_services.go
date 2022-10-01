// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mohsin123321/cloud-project/dataservice (interfaces: DataserviceInterface)

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dto "github.com/mohsin123321/cloud-project/dto"
)

// MockDataserviceInterface is a mock of DataserviceInterface interface.
type MockDataserviceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDataserviceInterfaceMockRecorder
}

// MockDataserviceInterfaceMockRecorder is the mock recorder for MockDataserviceInterface.
type MockDataserviceInterfaceMockRecorder struct {
	mock *MockDataserviceInterface
}

// NewMockDataserviceInterface creates a new mock instance.
func NewMockDataserviceInterface(ctrl *gomock.Controller) *MockDataserviceInterface {
	mock := &MockDataserviceInterface{ctrl: ctrl}
	mock.recorder = &MockDataserviceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataserviceInterface) EXPECT() *MockDataserviceInterfaceMockRecorder {
	return m.recorder
}

// InsertData mocks base method.
func (m *MockDataserviceInterface) InsertData(arg0 dto.PostDataBody) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InsertData", arg0)
}

// InsertData indicates an expected call of InsertData.
func (mr *MockDataserviceInterfaceMockRecorder) InsertData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertData", reflect.TypeOf((*MockDataserviceInterface)(nil).InsertData), arg0)
}