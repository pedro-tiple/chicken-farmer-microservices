// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/farm/service_grpc.go

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
func (m *MockIController) BuyBarn(ctx context.Context, farmerID, farmID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuyBarn", ctx, farmerID, farmID)
	ret0, _ := ret[0].(error)
	return ret0
}

// BuyBarn indicates an expected call of BuyBarn.
func (mr *MockIControllerMockRecorder) BuyBarn(ctx, farmerID, farmID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuyBarn", reflect.TypeOf((*MockIController)(nil).BuyBarn), ctx, farmerID, farmID)
}

// BuyChicken mocks base method.
func (m *MockIController) BuyChicken(ctx context.Context, farmerID, barnID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuyChicken", ctx, farmerID, barnID)
	ret0, _ := ret[0].(error)
	return ret0
}

// BuyChicken indicates an expected call of BuyChicken.
func (mr *MockIControllerMockRecorder) BuyChicken(ctx, farmerID, barnID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuyChicken", reflect.TypeOf((*MockIController)(nil).BuyChicken), ctx, farmerID, barnID)
}

// BuyFeedBags mocks base method.
func (m *MockIController) BuyFeedBags(ctx context.Context, farmerID, barnID uuid.UUID, amount uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuyFeedBags", ctx, farmerID, barnID, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// BuyFeedBags indicates an expected call of BuyFeedBags.
func (mr *MockIControllerMockRecorder) BuyFeedBags(ctx, farmerID, barnID, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuyFeedBags", reflect.TypeOf((*MockIController)(nil).BuyFeedBags), ctx, farmerID, barnID, amount)
}

// FarmDetails mocks base method.
func (m *MockIController) FarmDetails(ctx context.Context, farmerID, farmID uuid.UUID) (FarmDetailsResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FarmDetails", ctx, farmerID, farmID)
	ret0, _ := ret[0].(FarmDetailsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FarmDetails indicates an expected call of FarmDetails.
func (mr *MockIControllerMockRecorder) FarmDetails(ctx, farmerID, farmID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FarmDetails", reflect.TypeOf((*MockIController)(nil).FarmDetails), ctx, farmerID, farmID)
}

// FeedChicken mocks base method.
func (m *MockIController) FeedChicken(ctx context.Context, farmerID, chickenID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FeedChicken", ctx, farmerID, chickenID)
	ret0, _ := ret[0].(error)
	return ret0
}

// FeedChicken indicates an expected call of FeedChicken.
func (mr *MockIControllerMockRecorder) FeedChicken(ctx, farmerID, chickenID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FeedChicken", reflect.TypeOf((*MockIController)(nil).FeedChicken), ctx, farmerID, chickenID)
}

// NewFarm mocks base method.
func (m *MockIController) NewFarm(ctx context.Context, farmerID uuid.UUID, name string) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewFarm", ctx, farmerID, name)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewFarm indicates an expected call of NewFarm.
func (mr *MockIControllerMockRecorder) NewFarm(ctx, farmerID, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewFarm", reflect.TypeOf((*MockIController)(nil).NewFarm), ctx, farmerID, name)
}

// SellChicken mocks base method.
func (m *MockIController) SellChicken(ctx context.Context, farmerID, chickenID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SellChicken", ctx, farmerID, chickenID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SellChicken indicates an expected call of SellChicken.
func (mr *MockIControllerMockRecorder) SellChicken(ctx, farmerID, chickenID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SellChicken", reflect.TypeOf((*MockIController)(nil).SellChicken), ctx, farmerID, chickenID)
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
