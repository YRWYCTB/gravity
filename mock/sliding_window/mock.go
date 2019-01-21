// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/moiot/gravity/pkg/sliding_window (interfaces: WindowItem)

// Package mock_sliding_window is a generated GoMock package.
package mock_sliding_window

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockWindowItem is a mock of WindowItem interface
type MockWindowItem struct {
	ctrl     *gomock.Controller
	recorder *MockWindowItemMockRecorder
}

// MockWindowItemMockRecorder is the mock recorder for MockWindowItem
type MockWindowItemMockRecorder struct {
	mock *MockWindowItem
}

// NewMockWindowItem creates a new mock instance
func NewMockWindowItem(ctrl *gomock.Controller) *MockWindowItem {
	mock := &MockWindowItem{ctrl: ctrl}
	mock.recorder = &MockWindowItemMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWindowItem) EXPECT() *MockWindowItemMockRecorder {
	return m.recorder
}

// BeforeWindowMoveForward mocks base method
func (m *MockWindowItem) BeforeWindowMoveForward() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BeforeWindowMoveForward")
}

// BeforeWindowMoveForward indicates an expected call of BeforeWindowMoveForward
func (mr *MockWindowItemMockRecorder) BeforeWindowMoveForward() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeWindowMoveForward", reflect.TypeOf((*MockWindowItem)(nil).BeforeWindowMoveForward))
}

// EventTime mocks base method
func (m *MockWindowItem) EventTime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// EventTime indicates an expected call of EventTime
func (mr *MockWindowItemMockRecorder) EventTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventTime", reflect.TypeOf((*MockWindowItem)(nil).EventTime))
}

// SequenceNumber mocks base method
func (m *MockWindowItem) SequenceNumber() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SequenceNumber")
	ret0, _ := ret[0].(int64)
	return ret0
}

// SequenceNumber indicates an expected call of SequenceNumber
func (mr *MockWindowItemMockRecorder) SequenceNumber() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SequenceNumber", reflect.TypeOf((*MockWindowItem)(nil).SequenceNumber))
}