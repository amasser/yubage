// Code generated by MockGen. DO NOT EDIT.
// Source: eagain.net/go/yubage/internal/pivcard (interfaces: Opener,Card)

// Package mock_pivcard is a generated GoMock package.
package mock_pivcard

import (
	ecdsa "crypto/ecdsa"
	pivcard "eagain.net/go/yubage/internal/pivcard"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockOpener is a mock of Opener interface
type MockOpener struct {
	ctrl     *gomock.Controller
	recorder *MockOpenerMockRecorder
}

// MockOpenerMockRecorder is the mock recorder for MockOpener
type MockOpenerMockRecorder struct {
	mock *MockOpener
}

// NewMockOpener creates a new mock instance
func NewMockOpener(ctrl *gomock.Controller) *MockOpener {
	mock := &MockOpener{ctrl: ctrl}
	mock.recorder = &MockOpenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOpener) EXPECT() *MockOpenerMockRecorder {
	return m.recorder
}

// Open mocks base method
func (m *MockOpener) Open(arg0 uint32, arg1 byte) (pivcard.Card, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", arg0, arg1)
	ret0, _ := ret[0].(pivcard.Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open
func (mr *MockOpenerMockRecorder) Open(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockOpener)(nil).Open), arg0, arg1)
}

// MockCard is a mock of Card interface
type MockCard struct {
	ctrl     *gomock.Controller
	recorder *MockCardMockRecorder
}

// MockCardMockRecorder is the mock recorder for MockCard
type MockCardMockRecorder struct {
	mock *MockCard
}

// NewMockCard creates a new mock instance
func NewMockCard(ctrl *gomock.Controller) *MockCard {
	mock := &MockCard{ctrl: ctrl}
	mock.recorder = &MockCardMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCard) EXPECT() *MockCardMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockCard) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockCardMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCard)(nil).Close))
}

// Public mocks base method
func (m *MockCard) Public() *ecdsa.PublicKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Public")
	ret0, _ := ret[0].(*ecdsa.PublicKey)
	return ret0
}

// Public indicates an expected call of Public
func (mr *MockCardMockRecorder) Public() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Public", reflect.TypeOf((*MockCard)(nil).Public))
}

// SharedKey mocks base method
func (m *MockCard) SharedKey(arg0 *ecdsa.PublicKey, arg1 pivcard.Prompter) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SharedKey", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SharedKey indicates an expected call of SharedKey
func (mr *MockCardMockRecorder) SharedKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SharedKey", reflect.TypeOf((*MockCard)(nil).SharedKey), arg0, arg1)
}
