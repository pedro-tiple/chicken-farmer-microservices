// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/farm/controller.go

// Package farm is a generated GoMock package.
package farm

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockIDataSource is a mock of IDataSource interface.
type MockIDataSource struct {
	ctrl     *gomock.Controller
	recorder *MockIDataSourceMockRecorder
}

// MockIDataSourceMockRecorder is the mock recorder for MockIDataSource.
type MockIDataSourceMockRecorder struct {
	mock *MockIDataSource
}

// NewMockIDataSource creates a new mock instance.
func NewMockIDataSource(ctrl *gomock.Controller) *MockIDataSource {
	mock := &MockIDataSource{ctrl: ctrl}
	mock.recorder = &MockIDataSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDataSource) EXPECT() *MockIDataSourceMockRecorder {
	return m.recorder
}

// DecrementBarnFeedGreaterEqualThan mocks base method.
func (m *MockIDataSource) DecrementBarnFeedGreaterEqualThan(ctx context.Context, barnID uuid.UUID, amount uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecrementBarnFeedGreaterEqualThan", ctx, barnID, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// DecrementBarnFeedGreaterEqualThan indicates an expected call of DecrementBarnFeedGreaterEqualThan.
func (mr *MockIDataSourceMockRecorder) DecrementBarnFeedGreaterEqualThan(ctx, barnID, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecrementBarnFeedGreaterEqualThan", reflect.TypeOf((*MockIDataSource)(nil).DecrementBarnFeedGreaterEqualThan), ctx, barnID, amount)
}

// DeleteChicken mocks base method.
func (m *MockIDataSource) DeleteChicken(ctx context.Context, chickenID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteChicken", ctx, chickenID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteChicken indicates an expected call of DeleteChicken.
func (mr *MockIDataSourceMockRecorder) DeleteChicken(ctx, chickenID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteChicken", reflect.TypeOf((*MockIDataSource)(nil).DeleteChicken), ctx, chickenID)
}

// GetBarn mocks base method.
func (m *MockIDataSource) GetBarn(ctx context.Context, barnID uuid.UUID) (Barn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBarn", ctx, barnID)
	ret0, _ := ret[0].(Barn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBarn indicates an expected call of GetBarn.
func (mr *MockIDataSourceMockRecorder) GetBarn(ctx, barnID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBarn", reflect.TypeOf((*MockIDataSource)(nil).GetBarn), ctx, barnID)
}

// GetBarnsOfFarm mocks base method.
func (m *MockIDataSource) GetBarnsOfFarm(ctx context.Context, farmID uuid.UUID) ([]Barn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBarnsOfFarm", ctx, farmID)
	ret0, _ := ret[0].([]Barn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBarnsOfFarm indicates an expected call of GetBarnsOfFarm.
func (mr *MockIDataSourceMockRecorder) GetBarnsOfFarm(ctx, farmID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBarnsOfFarm", reflect.TypeOf((*MockIDataSource)(nil).GetBarnsOfFarm), ctx, farmID)
}

// GetChicken mocks base method.
func (m *MockIDataSource) GetChicken(ctx context.Context, chickenID uuid.UUID) (Chicken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChicken", ctx, chickenID)
	ret0, _ := ret[0].(Chicken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChicken indicates an expected call of GetChicken.
func (mr *MockIDataSourceMockRecorder) GetChicken(ctx, chickenID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChicken", reflect.TypeOf((*MockIDataSource)(nil).GetChicken), ctx, chickenID)
}

// GetChickensOfBarn mocks base method.
func (m *MockIDataSource) GetChickensOfBarn(ctx context.Context, barnID uuid.UUID) ([]Chicken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChickensOfBarn", ctx, barnID)
	ret0, _ := ret[0].([]Chicken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChickensOfBarn indicates an expected call of GetChickensOfBarn.
func (mr *MockIDataSourceMockRecorder) GetChickensOfBarn(ctx, barnID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChickensOfBarn", reflect.TypeOf((*MockIDataSource)(nil).GetChickensOfBarn), ctx, barnID)
}

// GetFarm mocks base method.
func (m *MockIDataSource) GetFarm(ctx context.Context, farmID uuid.UUID) (Farm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFarm", ctx, farmID)
	ret0, _ := ret[0].(Farm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFarm indicates an expected call of GetFarm.
func (mr *MockIDataSourceMockRecorder) GetFarm(ctx, farmID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFarm", reflect.TypeOf((*MockIDataSource)(nil).GetFarm), ctx, farmID)
}

// IncrementBarnFeed mocks base method.
func (m *MockIDataSource) IncrementBarnFeed(ctx context.Context, barnID uuid.UUID, amount uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrementBarnFeed", ctx, barnID, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncrementBarnFeed indicates an expected call of IncrementBarnFeed.
func (mr *MockIDataSourceMockRecorder) IncrementBarnFeed(ctx, barnID, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrementBarnFeed", reflect.TypeOf((*MockIDataSource)(nil).IncrementBarnFeed), ctx, barnID, amount)
}

// IncrementChickenEggLayCount mocks base method.
func (m *MockIDataSource) IncrementChickenEggLayCount(ctx context.Context, chickenID uuid.UUID, normalEggCount, goldEggCount int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrementChickenEggLayCount", ctx, chickenID, normalEggCount, goldEggCount)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncrementChickenEggLayCount indicates an expected call of IncrementChickenEggLayCount.
func (mr *MockIDataSourceMockRecorder) IncrementChickenEggLayCount(ctx, chickenID, normalEggCount, goldEggCount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrementChickenEggLayCount", reflect.TypeOf((*MockIDataSource)(nil).IncrementChickenEggLayCount), ctx, chickenID, normalEggCount, goldEggCount)
}

// InsertBarn mocks base method.
func (m *MockIDataSource) InsertBarn(ctx context.Context, barn Barn) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertBarn", ctx, barn)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertBarn indicates an expected call of InsertBarn.
func (mr *MockIDataSourceMockRecorder) InsertBarn(ctx, barn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertBarn", reflect.TypeOf((*MockIDataSource)(nil).InsertBarn), ctx, barn)
}

// InsertChicken mocks base method.
func (m *MockIDataSource) InsertChicken(ctx context.Context, chicken Chicken) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertChicken", ctx, chicken)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertChicken indicates an expected call of InsertChicken.
func (mr *MockIDataSourceMockRecorder) InsertChicken(ctx, chicken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertChicken", reflect.TypeOf((*MockIDataSource)(nil).InsertChicken), ctx, chicken)
}

// InsertFarm mocks base method.
func (m *MockIDataSource) InsertFarm(ctx context.Context, farm Farm) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertFarm", ctx, farm)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertFarm indicates an expected call of InsertFarm.
func (mr *MockIDataSourceMockRecorder) InsertFarm(ctx, farm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertFarm", reflect.TypeOf((*MockIDataSource)(nil).InsertFarm), ctx, farm)
}

// UpdateChickenRestingUntil mocks base method.
func (m *MockIDataSource) UpdateChickenRestingUntil(ctx context.Context, chickenID uuid.UUID, day uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateChickenRestingUntil", ctx, chickenID, day)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateChickenRestingUntil indicates an expected call of UpdateChickenRestingUntil.
func (mr *MockIDataSourceMockRecorder) UpdateChickenRestingUntil(ctx, chickenID, day interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateChickenRestingUntil", reflect.TypeOf((*MockIDataSource)(nil).UpdateChickenRestingUntil), ctx, chickenID, day)
}

// MockIFarmerService is a mock of IFarmerService interface.
type MockIFarmerService struct {
	ctrl     *gomock.Controller
	recorder *MockIFarmerServiceMockRecorder
}

// MockIFarmerServiceMockRecorder is the mock recorder for MockIFarmerService.
type MockIFarmerServiceMockRecorder struct {
	mock *MockIFarmerService
}

// NewMockIFarmerService creates a new mock instance.
func NewMockIFarmerService(ctrl *gomock.Controller) *MockIFarmerService {
	mock := &MockIFarmerService{ctrl: ctrl}
	mock.recorder = &MockIFarmerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFarmerService) EXPECT() *MockIFarmerServiceMockRecorder {
	return m.recorder
}

// GetGoldEggs mocks base method.
func (m *MockIFarmerService) GetGoldEggs(ctx context.Context) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGoldEggs", ctx)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGoldEggs indicates an expected call of GetGoldEggs.
func (mr *MockIFarmerServiceMockRecorder) GetGoldEggs(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGoldEggs", reflect.TypeOf((*MockIFarmerService)(nil).GetGoldEggs), ctx)
}

// GrantGoldEggs mocks base method.
func (m *MockIFarmerService) GrantGoldEggs(ctx context.Context, amount uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GrantGoldEggs", ctx, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// GrantGoldEggs indicates an expected call of GrantGoldEggs.
func (mr *MockIFarmerServiceMockRecorder) GrantGoldEggs(ctx, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantGoldEggs", reflect.TypeOf((*MockIFarmerService)(nil).GrantGoldEggs), ctx, amount)
}

// SpendGoldEggs mocks base method.
func (m *MockIFarmerService) SpendGoldEggs(ctx context.Context, amount uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpendGoldEggs", ctx, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// SpendGoldEggs indicates an expected call of SpendGoldEggs.
func (mr *MockIFarmerServiceMockRecorder) SpendGoldEggs(ctx, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpendGoldEggs", reflect.TypeOf((*MockIFarmerService)(nil).SpendGoldEggs), ctx, amount)
}
