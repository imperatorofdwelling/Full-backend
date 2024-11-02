// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	usersreports "github.com/imperatorofdwelling/Full-backend/internal/domain/models/usersreports"
)

// UsersReportsRepo is an autogenerated mock type for the UsersReportsRepo type
type UsersReportsRepo struct {
	mock.Mock
}

// CreateUsersReports provides a mock function with given fields: ctx, userId, toBlameId, title, description
func (_m *UsersReportsRepo) CreateUsersReports(ctx context.Context, userId string, toBlameId string, title string, description string) error {
	ret := _m.Called(ctx, userId, toBlameId, title, description)

	if len(ret) == 0 {
		panic("no return value specified for CreateUsersReports")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) error); ok {
		r0 = rf(ctx, userId, toBlameId, title, description)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUsersReports provides a mock function with given fields: ctx, userId, reportId
func (_m *UsersReportsRepo) DeleteUsersReports(ctx context.Context, userId string, reportId string) error {
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
func (_m *UsersReportsRepo) GetAllUsersReports(ctx context.Context, userId string) ([]usersreports.UsersReportEntity, error) {
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

// UpdateUsersReports provides a mock function with given fields: ctx, userId, toBlameId, title, description
func (_m *UsersReportsRepo) UpdateUsersReports(ctx context.Context, userId string, toBlameId string, title string, description string) (*usersreports.UsersReportEntity, error) {
	ret := _m.Called(ctx, userId, toBlameId, title, description)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUsersReports")
	}

	var r0 *usersreports.UsersReportEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) (*usersreports.UsersReportEntity, error)); ok {
		return rf(ctx, userId, toBlameId, title, description)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) *usersreports.UsersReportEntity); ok {
		r0 = rf(ctx, userId, toBlameId, title, description)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usersreports.UsersReportEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string) error); ok {
		r1 = rf(ctx, userId, toBlameId, title, description)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUsersReportsRepo creates a new instance of UsersReportsRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUsersReportsRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *UsersReportsRepo {
	mock := &UsersReportsRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
