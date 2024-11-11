// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	staysreports "github.com/imperatorofdwelling/Full-backend/internal/domain/models/staysreports"
)

// StaysReportsRepo is an autogenerated mock type for the StaysReportsRepo type
type StaysReportsRepo struct {
	mock.Mock
}

// CreateStaysReports provides a mock function with given fields: ctx, userId, stayId, title, description
func (_m *StaysReportsRepo) CreateStaysReports(ctx context.Context, userId string, stayId string, title string, description string) error {
	ret := _m.Called(ctx, userId, stayId, title, description)

	if len(ret) == 0 {
		panic("no return value specified for CreateStaysReports")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) error); ok {
		r0 = rf(ctx, userId, stayId, title, description)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteStaysReports provides a mock function with given fields: ctx, userId, reportId
func (_m *StaysReportsRepo) DeleteStaysReports(ctx context.Context, userId string, reportId string) error {
	ret := _m.Called(ctx, userId, reportId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteStaysReports")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, userId, reportId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllStaysReports provides a mock function with given fields: ctx, userId
func (_m *StaysReportsRepo) GetAllStaysReports(ctx context.Context, userId string) ([]staysreports.StaysReportEntity, error) {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for GetAllStaysReports")
	}

	var r0 []staysreports.StaysReportEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]staysreports.StaysReportEntity, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []staysreports.StaysReportEntity); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]staysreports.StaysReportEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStaysReports provides a mock function with given fields: ctx, userId, stayId, title, description
func (_m *StaysReportsRepo) UpdateStaysReports(ctx context.Context, userId string, stayId string, title string, description string) (*staysreports.StaysReportEntity, error) {
	ret := _m.Called(ctx, userId, stayId, title, description)

	if len(ret) == 0 {
		panic("no return value specified for UpdateStaysReports")
	}

	var r0 *staysreports.StaysReportEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) (*staysreports.StaysReportEntity, error)); ok {
		return rf(ctx, userId, stayId, title, description)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) *staysreports.StaysReportEntity); ok {
		r0 = rf(ctx, userId, stayId, title, description)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*staysreports.StaysReportEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string) error); ok {
		r1 = rf(ctx, userId, stayId, title, description)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewStaysReportsRepo creates a new instance of StaysReportsRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStaysReportsRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *StaysReportsRepo {
	mock := &StaysReportsRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}