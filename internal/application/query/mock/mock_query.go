// Code generated by MockGen. DO NOT EDIT.
// Source: active_evse_location_query.go
//
// Generated by this command:
//
//	mockgen -source=active_evse_location_query.go -destination=./mock/mock_query.go
//

// Package mock_query is a generated GoMock package.
package mock

import (
	context "context"
	domain "go-challenge/internal/domain"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockActiveEVSELocationQueryService is a mock of ActiveEVSELocationQueryService interface.
type MockActiveEVSELocationQueryService struct {
	ctrl     *gomock.Controller
	recorder *MockActiveEVSELocationQueryServiceMockRecorder
}

// MockActiveEVSELocationQueryServiceMockRecorder is the mock recorder for MockActiveEVSELocationQueryService.
type MockActiveEVSELocationQueryServiceMockRecorder struct {
	mock *MockActiveEVSELocationQueryService
}

// NewMockActiveEVSELocationQueryService creates a new mock instance.
func NewMockActiveEVSELocationQueryService(ctrl *gomock.Controller) *MockActiveEVSELocationQueryService {
	mock := &MockActiveEVSELocationQueryService{ctrl: ctrl}
	mock.recorder = &MockActiveEVSELocationQueryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActiveEVSELocationQueryService) EXPECT() *MockActiveEVSELocationQueryServiceMockRecorder {
	return m.recorder
}

// FindLocationsWithActiveEVSE mocks base method.
func (m *MockActiveEVSELocationQueryService) FindLocationsWithActiveEVSE(ctx context.Context, latitude, longitude float64, radius int) ([]domain.AvailableEVSELocation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindLocationsWithActiveEVSE", ctx, latitude, longitude, radius)
	ret0, _ := ret[0].([]domain.AvailableEVSELocation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindLocationsWithActiveEVSE indicates an expected call of FindLocationsWithActiveEVSE.
func (mr *MockActiveEVSELocationQueryServiceMockRecorder) FindLocationsWithActiveEVSE(ctx, latitude, longitude, radius any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindLocationsWithActiveEVSE", reflect.TypeOf((*MockActiveEVSELocationQueryService)(nil).FindLocationsWithActiveEVSE), ctx, latitude, longitude, radius)
}
