// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -source=interface.go -destination=mock_interface.go -package=ftp
//

// Package ftp is a generated GoMock package.
package ftp

import (
	context "context"
	io "io"
	reflect "reflect"
	time "time"

	ftp "github.com/jlaffaye/ftp"
	gomock "go.uber.org/mock/gomock"
)

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Debug mocks base method.
func (m *MockLogger) Debug(args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debug", varargs...)
}

// Debug indicates an expected call of Debug.
func (mr *MockLoggerMockRecorder) Debug(args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockLogger)(nil).Debug), args...)
}

// Errorf mocks base method.
func (m *MockLogger) Errorf(pattern string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{pattern}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockLoggerMockRecorder) Errorf(pattern any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pattern}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockLogger)(nil).Errorf), varargs...)
}

// Logf mocks base method.
func (m *MockLogger) Logf(pattern string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{pattern}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Logf", varargs...)
}

// Logf indicates an expected call of Logf.
func (mr *MockLoggerMockRecorder) Logf(pattern any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pattern}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logf", reflect.TypeOf((*MockLogger)(nil).Logf), varargs...)
}

// MockMetrics is a mock of Metrics interface.
type MockMetrics struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsMockRecorder
}

// MockMetricsMockRecorder is the mock recorder for MockMetrics.
type MockMetricsMockRecorder struct {
	mock *MockMetrics
}

// NewMockMetrics creates a new mock instance.
func NewMockMetrics(ctrl *gomock.Controller) *MockMetrics {
	mock := &MockMetrics{ctrl: ctrl}
	mock.recorder = &MockMetricsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetrics) EXPECT() *MockMetricsMockRecorder {
	return m.recorder
}

// NewHistogram mocks base method.
func (m *MockMetrics) NewHistogram(name, desc string, buckets ...float64) {
	m.ctrl.T.Helper()
	varargs := []any{name, desc}
	for _, a := range buckets {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "NewHistogram", varargs...)
}

// NewHistogram indicates an expected call of NewHistogram.
func (mr *MockMetricsMockRecorder) NewHistogram(name, desc any, buckets ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name, desc}, buckets...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewHistogram", reflect.TypeOf((*MockMetrics)(nil).NewHistogram), varargs...)
}

// RecordHistogram mocks base method.
func (m *MockMetrics) RecordHistogram(ctx context.Context, name string, value float64, labels ...string) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, name, value}
	for _, a := range labels {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "RecordHistogram", varargs...)
}

// RecordHistogram indicates an expected call of RecordHistogram.
func (mr *MockMetricsMockRecorder) RecordHistogram(ctx, name, value any, labels ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, name, value}, labels...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordHistogram", reflect.TypeOf((*MockMetrics)(nil).RecordHistogram), varargs...)
}

// MockserverConn is a mock of serverConn interface.
type MockserverConn struct {
	ctrl     *gomock.Controller
	recorder *MockserverConnMockRecorder
}

// MockserverConnMockRecorder is the mock recorder for MockserverConn.
type MockserverConnMockRecorder struct {
	mock *MockserverConn
}

// NewMockserverConn creates a new mock instance.
func NewMockserverConn(ctrl *gomock.Controller) *MockserverConn {
	mock := &MockserverConn{ctrl: ctrl}
	mock.recorder = &MockserverConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockserverConn) EXPECT() *MockserverConnMockRecorder {
	return m.recorder
}

// ChangeDir mocks base method.
func (m *MockserverConn) ChangeDir(path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeDir", path)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeDir indicates an expected call of ChangeDir.
func (mr *MockserverConnMockRecorder) ChangeDir(path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeDir", reflect.TypeOf((*MockserverConn)(nil).ChangeDir), path)
}

// CurrentDir mocks base method.
func (m *MockserverConn) CurrentDir() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentDir")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrentDir indicates an expected call of CurrentDir.
func (mr *MockserverConnMockRecorder) CurrentDir() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentDir", reflect.TypeOf((*MockserverConn)(nil).CurrentDir))
}

// Delete mocks base method.
func (m *MockserverConn) Delete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockserverConnMockRecorder) Delete(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockserverConn)(nil).Delete), arg0)
}

// FileSize mocks base method.
func (m *MockserverConn) FileSize(name string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FileSize", name)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FileSize indicates an expected call of FileSize.
func (mr *MockserverConnMockRecorder) FileSize(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileSize", reflect.TypeOf((*MockserverConn)(nil).FileSize), name)
}

// GetTime mocks base method.
func (m *MockserverConn) GetTime(path string) (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTime", path)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTime indicates an expected call of GetTime.
func (mr *MockserverConnMockRecorder) GetTime(path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTime", reflect.TypeOf((*MockserverConn)(nil).GetTime), path)
}

// List mocks base method.
func (m *MockserverConn) List(arg0 string) ([]*ftp.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]*ftp.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockserverConnMockRecorder) List(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockserverConn)(nil).List), arg0)
}

// Login mocks base method.
func (m *MockserverConn) Login(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Login indicates an expected call of Login.
func (mr *MockserverConnMockRecorder) Login(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockserverConn)(nil).Login), arg0, arg1)
}

// MakeDir mocks base method.
func (m *MockserverConn) MakeDir(path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeDir", path)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeDir indicates an expected call of MakeDir.
func (mr *MockserverConnMockRecorder) MakeDir(path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeDir", reflect.TypeOf((*MockserverConn)(nil).MakeDir), path)
}

// Quit mocks base method.
func (m *MockserverConn) Quit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Quit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Quit indicates an expected call of Quit.
func (mr *MockserverConnMockRecorder) Quit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Quit", reflect.TypeOf((*MockserverConn)(nil).Quit))
}

// RemoveDir mocks base method.
func (m *MockserverConn) RemoveDir(path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveDir", path)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveDir indicates an expected call of RemoveDir.
func (mr *MockserverConnMockRecorder) RemoveDir(path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDir", reflect.TypeOf((*MockserverConn)(nil).RemoveDir), path)
}

// RemoveDirRecur mocks base method.
func (m *MockserverConn) RemoveDirRecur(path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveDirRecur", path)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveDirRecur indicates an expected call of RemoveDirRecur.
func (mr *MockserverConnMockRecorder) RemoveDirRecur(path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDirRecur", reflect.TypeOf((*MockserverConn)(nil).RemoveDirRecur), path)
}

// Rename mocks base method.
func (m *MockserverConn) Rename(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rename", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rename indicates an expected call of Rename.
func (mr *MockserverConnMockRecorder) Rename(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rename", reflect.TypeOf((*MockserverConn)(nil).Rename), arg0, arg1)
}

// Retr mocks base method.
func (m *MockserverConn) Retr(arg0 string) (ftpResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Retr", arg0)
	ret0, _ := ret[0].(ftpResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Retr indicates an expected call of Retr.
func (mr *MockserverConnMockRecorder) Retr(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Retr", reflect.TypeOf((*MockserverConn)(nil).Retr), arg0)
}

// RetrFrom mocks base method.
func (m *MockserverConn) RetrFrom(arg0 string, arg1 uint64) (ftpResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrFrom", arg0, arg1)
	ret0, _ := ret[0].(ftpResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrFrom indicates an expected call of RetrFrom.
func (mr *MockserverConnMockRecorder) RetrFrom(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrFrom", reflect.TypeOf((*MockserverConn)(nil).RetrFrom), arg0, arg1)
}

// Stor mocks base method.
func (m *MockserverConn) Stor(arg0 string, arg1 io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stor", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stor indicates an expected call of Stor.
func (mr *MockserverConnMockRecorder) Stor(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stor", reflect.TypeOf((*MockserverConn)(nil).Stor), arg0, arg1)
}

// StorFrom mocks base method.
func (m *MockserverConn) StorFrom(arg0 string, arg1 io.Reader, arg2 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorFrom", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// StorFrom indicates an expected call of StorFrom.
func (mr *MockserverConnMockRecorder) StorFrom(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorFrom", reflect.TypeOf((*MockserverConn)(nil).StorFrom), arg0, arg1, arg2)
}

// MockftpResponse is a mock of ftpResponse interface.
type MockftpResponse struct {
	ctrl     *gomock.Controller
	recorder *MockftpResponseMockRecorder
}

// MockftpResponseMockRecorder is the mock recorder for MockftpResponse.
type MockftpResponseMockRecorder struct {
	mock *MockftpResponse
}

// NewMockftpResponse creates a new mock instance.
func NewMockftpResponse(ctrl *gomock.Controller) *MockftpResponse {
	mock := &MockftpResponse{ctrl: ctrl}
	mock.recorder = &MockftpResponseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockftpResponse) EXPECT() *MockftpResponseMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockftpResponse) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockftpResponseMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockftpResponse)(nil).Close))
}

// Read mocks base method.
func (m *MockftpResponse) Read(buf []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", buf)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockftpResponseMockRecorder) Read(buf any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockftpResponse)(nil).Read), buf)
}

// SetDeadline mocks base method.
func (m *MockftpResponse) SetDeadline(t time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDeadline", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDeadline indicates an expected call of SetDeadline.
func (mr *MockftpResponseMockRecorder) SetDeadline(t any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeadline", reflect.TypeOf((*MockftpResponse)(nil).SetDeadline), t)
}
