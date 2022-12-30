// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/farm/service.go

// Package farm is a generated GoMock package.
package farm

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockIController is a mock of IController interface.
type MockIController struct {
	ctrl     *gomock.Controller
	recorder *MockIControllerMockRecorder
}

// MockIControllerMockRecorder is the mock recorder for MockIController.
type MockIControllerMockRecorder struct {
	mock *MockIController
}

// NewMockIController creates a new mock instance.
func NewMockIController(ctrl *gomock.Controller) *MockIController {
	mock := &MockIController{ctrl: ctrl}
	mock.recorder = &MockIControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIController) EXPECT() *MockIControllerMockRecorder {
	return m.recorder
}

// BuyBarn mocks base method.
func (m *MockIController) BuyBarn(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuyBarn", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// BuyBarn indicates an expected call of BuyBarn.
func (mr *MockIControllerMockRecorder) BuyBarn(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuyBarn", reflect.TypeOf((*MockIController)(nil).BuyBarn), ctx)
}

// BuyChicken mocks base method.
func (m *MockIController) BuyChicken(ctx context.Context, barnID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuyChicken", ctx, barnID)
	ret0, _ := ret[0].(error)
	return ret0
}

// BuyChicken indicates an expected call of BuyChicken.
func (mr *MockIControllerMockRecorder) BuyChicken(ctx, barnID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuyChicken", reflect.TypeOf((*MockIController)(nil).BuyChicken), ctx, barnID)
}

// BuyFeed mocks base method.
func (m *MockIController) BuyFeed(ctx context.Context, barnID uuid.UUID, amount uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuyFeed", ctx, barnID, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// BuyFeed indicates an expected call of BuyFeed.
func (mr *MockIControllerMockRecorder) BuyFeed(ctx, barnID, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuyFeed", reflect.TypeOf((*MockIController)(nil).BuyFeed), ctx, barnID, amount)
}

// FeedChicken mocks base method.
func (m *MockIController) FeedChicken(ctx context.Context, chickenID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FeedChicken", ctx, chickenID)
	ret0, _ := ret[0].(error)
	return ret0
}

// FeedChicken indicates an expected call of FeedChicken.
func (mr *MockIControllerMockRecorder) FeedChicken(ctx, chickenID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FeedChicken", reflect.TypeOf((*MockIController)(nil).FeedChicken), ctx, chickenID)
}

// FeedChickensOfBarn mocks base method.
func (m *MockIController) FeedChickensOfBarn(ctx context.Context, barnID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FeedChickensOfBarn", ctx, barnID)
	ret0, _ := ret[0].(error)
	return ret0
}

// FeedChickensOfBarn indicates an expected call of FeedChickensOfBarn.
func (mr *MockIControllerMockRecorder) FeedChickensOfBarn(ctx, barnID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FeedChickensOfBarn", reflect.TypeOf((*MockIController)(nil).FeedChickensOfBarn), ctx, barnID)
}

// GetFarm mocks base method.
func (m *MockIController) GetFarm(ctx context.Context) (GetFarmResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFarm", ctx)
	ret0, _ := ret[0].(GetFarmResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFarm indicates an expected call of GetFarm.
func (mr *MockIControllerMockRecorder) GetFarm(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFarm", reflect.TypeOf((*MockIController)(nil).GetFarm), ctx)
}

// SetDay mocks base method.
func (m *MockIController) SetDay(ctx context.Context, day uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDay", ctx, day)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDay indicates an expected call of SetDay.
func (mr *MockIControllerMockRecorder) SetDay(ctx, day interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDay", reflect.TypeOf((*MockIController)(nil).SetDay), ctx, day)
}
