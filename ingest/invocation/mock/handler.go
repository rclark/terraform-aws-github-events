// Code generated by MockGen. DO NOT EDIT.
// Source: ./handler.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	cloudwatchevents "github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockLogger is a mock of Logger interface
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Clear mocks base method
func (m *MockLogger) Clear() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Clear")
}

// Clear indicates an expected call of Clear
func (mr *MockLoggerMockRecorder) Clear() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockLogger)(nil).Clear))
}

// Set mocks base method
func (m *MockLogger) Set(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", arg0, arg1)
}

// Set indicates an expected call of Set
func (mr *MockLoggerMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockLogger)(nil).Set), arg0, arg1)
}

// Print mocks base method
func (m *MockLogger) Print() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Print")
}

// Print indicates an expected call of Print
func (mr *MockLoggerMockRecorder) Print() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Print", reflect.TypeOf((*MockLogger)(nil).Print))
}

// MockCanPutEvents is a mock of CanPutEvents interface
type MockCanPutEvents struct {
	ctrl     *gomock.Controller
	recorder *MockCanPutEventsMockRecorder
}

// MockCanPutEventsMockRecorder is the mock recorder for MockCanPutEvents
type MockCanPutEventsMockRecorder struct {
	mock *MockCanPutEvents
}

// NewMockCanPutEvents creates a new mock instance
func NewMockCanPutEvents(ctrl *gomock.Controller) *MockCanPutEvents {
	mock := &MockCanPutEvents{ctrl: ctrl}
	mock.recorder = &MockCanPutEventsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCanPutEvents) EXPECT() *MockCanPutEventsMockRecorder {
	return m.recorder
}

// PutEvents mocks base method
func (m *MockCanPutEvents) PutEvents(ctx context.Context, params *cloudwatchevents.PutEventsInput, optFns ...func(*cloudwatchevents.Options)) (*cloudwatchevents.PutEventsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutEvents", varargs...)
	ret0, _ := ret[0].(*cloudwatchevents.PutEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutEvents indicates an expected call of PutEvents
func (mr *MockCanPutEventsMockRecorder) PutEvents(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutEvents", reflect.TypeOf((*MockCanPutEvents)(nil).PutEvents), varargs...)
}
