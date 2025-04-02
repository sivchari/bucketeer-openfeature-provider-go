// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bucketeer-io/go-server-sdk/pkg/bucketeer (interfaces: SDK)
//
// Generated by this command:
//
//	mockgen -destination=./mock_provider.go -package=bucketeeropenfeatureprovidergo github.com/bucketeer-io/go-server-sdk/pkg/bucketeer SDK
//

// Package bucketeeropenfeatureprovidergo is a generated GoMock package.
package bucketeeropenfeatureprovidergo

import (
	context "context"
	reflect "reflect"

	model "github.com/bucketeer-io/go-server-sdk/pkg/bucketeer/model"
	user "github.com/bucketeer-io/go-server-sdk/pkg/bucketeer/user"
	gomock "go.uber.org/mock/gomock"
)

// MockSDK is a mock of SDK interface.
type MockSDK struct {
	ctrl     *gomock.Controller
	recorder *MockSDKMockRecorder
	isgomock struct{}
}

// MockSDKMockRecorder is the mock recorder for MockSDK.
type MockSDKMockRecorder struct {
	mock *MockSDK
}

// NewMockSDK creates a new mock instance.
func NewMockSDK(ctrl *gomock.Controller) *MockSDK {
	mock := &MockSDK{ctrl: ctrl}
	mock.recorder = &MockSDKMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSDK) EXPECT() *MockSDKMockRecorder {
	return m.recorder
}

// BoolVariation mocks base method.
func (m *MockSDK) BoolVariation(ctx context.Context, user *user.User, featureID string, defaultValue bool) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BoolVariation", ctx, user, featureID, defaultValue)
	ret0, _ := ret[0].(bool)
	return ret0
}

// BoolVariation indicates an expected call of BoolVariation.
func (mr *MockSDKMockRecorder) BoolVariation(ctx, user, featureID, defaultValue any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BoolVariation", reflect.TypeOf((*MockSDK)(nil).BoolVariation), ctx, user, featureID, defaultValue)
}

// BoolVariationDetails mocks base method.
func (m *MockSDK) BoolVariationDetails(ctx context.Context, user *user.User, featureID string, defaultValue bool) model.BKTEvaluationDetails[bool] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BoolVariationDetails", ctx, user, featureID, defaultValue)
	ret0, _ := ret[0].(model.BKTEvaluationDetails[bool])
	return ret0
}

// BoolVariationDetails indicates an expected call of BoolVariationDetails.
func (mr *MockSDKMockRecorder) BoolVariationDetails(ctx, user, featureID, defaultValue any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BoolVariationDetails", reflect.TypeOf((*MockSDK)(nil).BoolVariationDetails), ctx, user, featureID, defaultValue)
}

// Close mocks base method.
func (m *MockSDK) Close(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockSDKMockRecorder) Close(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSDK)(nil).Close), ctx)
}

// Float64Variation mocks base method.
func (m *MockSDK) Float64Variation(ctx context.Context, user *user.User, featureID string, defaultValue float64) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Float64Variation", ctx, user, featureID, defaultValue)
	ret0, _ := ret[0].(float64)
	return ret0
}

// Float64Variation indicates an expected call of Float64Variation.
func (mr *MockSDKMockRecorder) Float64Variation(ctx, user, featureID, defaultValue any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Float64Variation", reflect.TypeOf((*MockSDK)(nil).Float64Variation), ctx, user, featureID, defaultValue)
}

// Float64VariationDetails mocks base method.
func (m *MockSDK) Float64VariationDetails(ctx context.Context, user *user.User, featureID string, defaultValue float64) model.BKTEvaluationDetails[float64] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Float64VariationDetails", ctx, user, featureID, defaultValue)
	ret0, _ := ret[0].(model.BKTEvaluationDetails[float64])
	return ret0
}

// Float64VariationDetails indicates an expected call of Float64VariationDetails.
func (mr *MockSDKMockRecorder) Float64VariationDetails(ctx, user, featureID, defaultValue any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Float64VariationDetails", reflect.TypeOf((*MockSDK)(nil).Float64VariationDetails), ctx, user, featureID, defaultValue)
}

// Int64Variation mocks base method.
func (m *MockSDK) Int64Variation(ctx context.Context, user *user.User, featureID string, defaultValue int64) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Int64Variation", ctx, user, featureID, defaultValue)
	ret0, _ := ret[0].(int64)
	return ret0
}

// Int64Variation indicates an expected call of Int64Variation.
func (mr *MockSDKMockRecorder) Int64Variation(ctx, user, featureID, defaultValue any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int64Variation", reflect.TypeOf((*MockSDK)(nil).Int64Variation), ctx, user, featureID, defaultValue)
}

// Int64VariationDetails mocks base method.
func (m *MockSDK) Int64VariationDetails(ctx context.Context, user *user.User, featureID string, defaultValue int64) model.BKTEvaluationDetails[int64] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Int64VariationDetails", ctx, user, featureID, defaultValue)
	ret0, _ := ret[0].(model.BKTEvaluationDetails[int64])
	return ret0
}

// Int64VariationDetails indicates an expected call of Int64VariationDetails.
func (mr *MockSDKMockRecorder) Int64VariationDetails(ctx, user, featureID, defaultValue any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int64VariationDetails", reflect.TypeOf((*MockSDK)(nil).Int64VariationDetails), ctx, user, featureID, defaultValue)
}

// IntVariation mocks base method.
func (m *MockSDK) IntVariation(ctx context.Context, user *user.User, featureID string, defaultValue int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IntVariation", ctx, user, featureID, defaultValue)
	ret0, _ := ret[0].(int)
	return ret0
}

// IntVariation indicates an expected call of IntVariation.
func (mr *MockSDKMockRecorder) IntVariation(ctx, user, featureID, defaultValue any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IntVariation", reflect.TypeOf((*MockSDK)(nil).IntVariation), ctx, user, featureID, defaultValue)
}

// IntVariationDetails mocks base method.
func (m *MockSDK) IntVariationDetails(ctx context.Context, user *user.User, featureID string, defaultValue int) model.BKTEvaluationDetails[int] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IntVariationDetails", ctx, user, featureID, defaultValue)
	ret0, _ := ret[0].(model.BKTEvaluationDetails[int])
	return ret0
}

// IntVariationDetails indicates an expected call of IntVariationDetails.
func (mr *MockSDKMockRecorder) IntVariationDetails(ctx, user, featureID, defaultValue any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IntVariationDetails", reflect.TypeOf((*MockSDK)(nil).IntVariationDetails), ctx, user, featureID, defaultValue)
}

// JSONVariation mocks base method.
func (m *MockSDK) JSONVariation(ctx context.Context, user *user.User, featureID string, dst any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "JSONVariation", ctx, user, featureID, dst)
}

// JSONVariation indicates an expected call of JSONVariation.
func (mr *MockSDKMockRecorder) JSONVariation(ctx, user, featureID, dst any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JSONVariation", reflect.TypeOf((*MockSDK)(nil).JSONVariation), ctx, user, featureID, dst)
}

// ObjectVariation mocks base method.
func (m *MockSDK) ObjectVariation(ctx context.Context, user *user.User, featureID string, defaultValue any) any {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectVariation", ctx, user, featureID, defaultValue)
	ret0, _ := ret[0].(any)
	return ret0
}

// ObjectVariation indicates an expected call of ObjectVariation.
func (mr *MockSDKMockRecorder) ObjectVariation(ctx, user, featureID, defaultValue any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectVariation", reflect.TypeOf((*MockSDK)(nil).ObjectVariation), ctx, user, featureID, defaultValue)
}

// ObjectVariationDetails mocks base method.
func (m *MockSDK) ObjectVariationDetails(ctx context.Context, user *user.User, featureID string, defaultValue any) model.BKTEvaluationDetails[any] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectVariationDetails", ctx, user, featureID, defaultValue)
	ret0, _ := ret[0].(model.BKTEvaluationDetails[any])
	return ret0
}

// ObjectVariationDetails indicates an expected call of ObjectVariationDetails.
func (mr *MockSDKMockRecorder) ObjectVariationDetails(ctx, user, featureID, defaultValue any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectVariationDetails", reflect.TypeOf((*MockSDK)(nil).ObjectVariationDetails), ctx, user, featureID, defaultValue)
}

// StringVariation mocks base method.
func (m *MockSDK) StringVariation(ctx context.Context, user *user.User, featureID, defaultValue string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StringVariation", ctx, user, featureID, defaultValue)
	ret0, _ := ret[0].(string)
	return ret0
}

// StringVariation indicates an expected call of StringVariation.
func (mr *MockSDKMockRecorder) StringVariation(ctx, user, featureID, defaultValue any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StringVariation", reflect.TypeOf((*MockSDK)(nil).StringVariation), ctx, user, featureID, defaultValue)
}

// StringVariationDetails mocks base method.
func (m *MockSDK) StringVariationDetails(ctx context.Context, user *user.User, featureID, defaultValue string) model.BKTEvaluationDetails[string] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StringVariationDetails", ctx, user, featureID, defaultValue)
	ret0, _ := ret[0].(model.BKTEvaluationDetails[string])
	return ret0
}

// StringVariationDetails indicates an expected call of StringVariationDetails.
func (mr *MockSDKMockRecorder) StringVariationDetails(ctx, user, featureID, defaultValue any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StringVariationDetails", reflect.TypeOf((*MockSDK)(nil).StringVariationDetails), ctx, user, featureID, defaultValue)
}

// Track mocks base method.
func (m *MockSDK) Track(ctx context.Context, user *user.User, GoalID string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Track", ctx, user, GoalID)
}

// Track indicates an expected call of Track.
func (mr *MockSDKMockRecorder) Track(ctx, user, GoalID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Track", reflect.TypeOf((*MockSDK)(nil).Track), ctx, user, GoalID)
}

// TrackValue mocks base method.
func (m *MockSDK) TrackValue(ctx context.Context, user *user.User, GoalID string, value float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "TrackValue", ctx, user, GoalID, value)
}

// TrackValue indicates an expected call of TrackValue.
func (mr *MockSDKMockRecorder) TrackValue(ctx, user, GoalID, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrackValue", reflect.TypeOf((*MockSDK)(nil).TrackValue), ctx, user, GoalID, value)
}
