// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/sclevine/cflocal/cf (interfaces: Help)

package mocks

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of Help interface
type MockHelp struct {
	ctrl     *gomock.Controller
	recorder *_MockHelpRecorder
}

// Recorder for MockHelp (not exported)
type _MockHelpRecorder struct {
	mock *MockHelp
}

func NewMockHelp(ctrl *gomock.Controller) *MockHelp {
	mock := &MockHelp{ctrl: ctrl}
	mock.recorder = &_MockHelpRecorder{mock}
	return mock
}

func (_m *MockHelp) EXPECT() *_MockHelpRecorder {
	return _m.recorder
}

func (_m *MockHelp) Long() {
	_m.ctrl.Call(_m, "Long")
}

func (_mr *_MockHelpRecorder) Long() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Long")
}
