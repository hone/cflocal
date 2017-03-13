// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/sclevine/cflocal/cf/cmd (interfaces: App)

package mocks

import (
	gomock "github.com/golang/mock/gomock"
	remote "github.com/sclevine/cflocal/remote"
	io "io"
)

// Mock of App interface
type MockApp struct {
	ctrl     *gomock.Controller
	recorder *_MockAppRecorder
}

// Recorder for MockApp (not exported)
type _MockAppRecorder struct {
	mock *MockApp
}

func NewMockApp(ctrl *gomock.Controller) *MockApp {
	mock := &MockApp{ctrl: ctrl}
	mock.recorder = &_MockAppRecorder{mock}
	return mock
}

func (_m *MockApp) EXPECT() *_MockAppRecorder {
	return _m.recorder
}

func (_m *MockApp) Command(_param0 string) (string, error) {
	ret := _m.ctrl.Call(_m, "Command", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAppRecorder) Command(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Command", arg0)
}

func (_m *MockApp) Droplet(_param0 string) (io.ReadCloser, int64, error) {
	ret := _m.ctrl.Call(_m, "Droplet", _param0)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockAppRecorder) Droplet(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Droplet", arg0)
}

func (_m *MockApp) Env(_param0 string) (*remote.AppEnv, error) {
	ret := _m.ctrl.Call(_m, "Env", _param0)
	ret0, _ := ret[0].(*remote.AppEnv)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAppRecorder) Env(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Env", arg0)
}

func (_m *MockApp) Forward(_param0 string, _param1 remote.Services) (remote.Services, string, error) {
	ret := _m.ctrl.Call(_m, "Forward", _param0, _param1)
	ret0, _ := ret[0].(remote.Services)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockAppRecorder) Forward(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Forward", arg0, arg1)
}

func (_m *MockApp) Services(_param0 string) (remote.Services, error) {
	ret := _m.ctrl.Call(_m, "Services", _param0)
	ret0, _ := ret[0].(remote.Services)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAppRecorder) Services(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Services", arg0)
}