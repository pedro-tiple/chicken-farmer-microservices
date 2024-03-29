// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package proto-old is a generated GoMock package.
package chicken_old

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAPI is a mock of API interface.
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI.
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance.
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// FeedChicken mocks base method.
func (m *MockAPI) FeedChicken(ctx context.Context, chickenID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FeedChicken", ctx, chickenID)
	ret0, _ := ret[0].(error)
	return ret0
}

// FeedChicken indicates an expected call of FeedChicken.
func (mr *MockAPIMockRecorder) FeedChicken(ctx, chickenID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FeedChicken", reflect.TypeOf((*MockAPI)(nil).FeedChicken), ctx, chickenID)
}

// GetChicken mocks base method.
func (m *MockAPI) GetChicken(ctx context.Context, chickenID string) (Chicken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChicken", ctx, chickenID)
	ret0, _ := ret[0].(Chicken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChicken indicates an expected call of GetChicken.
func (mr *MockAPIMockRecorder) GetChicken(ctx, chickenID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChicken", reflect.TypeOf((*MockAPI)(nil).GetChicken), ctx, chickenID)
}

// NewChicken mocks base method.
func (m *MockAPI) NewChicken(ctx context.Context, farmID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewChicken", ctx, farmID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewChicken indicates an expected call of NewChicken.
func (mr *MockAPIMockRecorder) NewChicken(ctx, farmID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewChicken", reflect.TypeOf((*MockAPI)(nil).NewChicken), ctx, farmID)
}
