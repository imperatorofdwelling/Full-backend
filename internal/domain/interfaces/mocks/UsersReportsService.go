// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	usersreports "github.com/imperatorofdwelling/Full-backend/internal/domain/models/usersreports"
)

// UsersReportsService is an autogenerated mock type for the UsersReportsService type
type UsersReportsService struct {
	mock.Mock
}

// CreateUsersReports provides a mock function with given fields: ctx, userId, toBlameId, title, description, image
func (_m *UsersReportsService) CreateUsersReports(ctx context.Context, userId string, toBlameId string, title string, description string, image []byte) error {
	ret := _m.Called(ctx, userId, toBlameId, title, description, image)

	if len(ret) == 0 {
		panic("no return value specified for CreateUsersReports")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, []byte) error); ok {
		r0 = rf(ctx, userId, toBlameId, title, description, image)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUsersReports provides a mock function with given fields: ctx, userId, reportId
func (_m *UsersReportsService) DeleteUsersReports(ctx context.Context, userId string, reportId string) error {
	ret := _m.Called(ctx, userId, reportId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUsersReports")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, userId, reportId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllUsersReports provides a mock function with given fields: ctx, userId
func (_m *UsersReportsService) GetAllUsersReports(ctx context.Context, userId string) ([]usersreports.UsersReportEntity, error) {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for GetAllUsersReports")
	}

	var r0 []usersreports.UsersReportEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]usersreports.UsersReportEntity, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []usersreports.UsersReportEntity); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]usersreports.UsersReportEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsersReportById provides a mock function with given fields: ctx, userId, id
func (_m *UsersReportsService) GetUsersReportById(ctx context.Context, userId string, id string) (*usersreports.UsersReport, error) {
	ret := _m.Called(ctx, userId, id)

	if len(ret) == 0 {
		panic("no return value specified for GetUsersReportById")
	}

	var r0 *usersreports.UsersReport
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*usersreports.UsersReport, error)); ok {
		return rf(ctx, userId, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *usersreports.UsersReport); ok {
		r0 = rf(ctx, userId, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usersreports.UsersReport)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, userId, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUsersReports provides a mock function with given fields: ctx, userId, reportId, title, description, imageData
func (_m *UsersReportsService) UpdateUsersReports(ctx context.Context, userId string, reportId string, title string, description string, imageData []byte) (*usersreports.UsersReportEntity, error) {
	ret := _m.Called(ctx, userId, reportId, title, description, imageData)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUsersReports")
	}

	var r0 *usersreports.UsersReportEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, []byte) (*usersreports.UsersReportEntity, error)); ok {
		return rf(ctx, userId, reportId, title, description, imageData)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, []byte) *usersreports.UsersReportEntity); ok {
		r0 = rf(ctx, userId, reportId, title, description, imageData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usersreports.UsersReportEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string, []byte) error); ok {
		r1 = rf(ctx, userId, reportId, title, description, imageData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUsersReportsService creates a new instance of UsersReportsService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUsersReportsService(t interface {
	mock.TestingT
	Cleanup(func())
}) *UsersReportsService {
	mock := &UsersReportsService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
