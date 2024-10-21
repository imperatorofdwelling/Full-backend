// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	user "github.com/imperatorofdwelling/Full-backend/internal/domain/models/user"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// DeleteUserByID provides a mock function with given fields: ctx, idStr
func (_m *UserService) DeleteUserByID(ctx context.Context, idStr string) error {
	ret := _m.Called(ctx, idStr)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUserByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, idStr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserByID provides a mock function with given fields: ctx, idStr
func (_m *UserService) GetUserByID(ctx context.Context, idStr string) (user.User, error) {
	ret := _m.Called(ctx, idStr)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByID")
	}

	var r0 user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (user.User, error)); ok {
		return rf(ctx, idStr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) user.User); ok {
		r0 = rf(ctx, idStr)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, idStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserByID provides a mock function with given fields: ctx, idStr, _a2
func (_m *UserService) UpdateUserByID(ctx context.Context, idStr string, _a2 user.User) (user.User, error) {
	ret := _m.Called(ctx, idStr, _a2)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUserByID")
	}

	var r0 user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, user.User) (user.User, error)); ok {
		return rf(ctx, idStr, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, user.User) user.User); ok {
		r0 = rf(ctx, idStr, _a2)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, user.User) error); ok {
		r1 = rf(ctx, idStr, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserService(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
