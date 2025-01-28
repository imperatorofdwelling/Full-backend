// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	newPassword "github.com/imperatorofdwelling/Full-backend/internal/domain/models/newPassword"

	user "github.com/imperatorofdwelling/Full-backend/internal/domain/models/user"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// CheckUserEmail provides a mock function with given fields: ctx, userID, newEmail
func (_m *UserService) CheckUserEmail(ctx context.Context, userID string, newEmail string) error {
	ret := _m.Called(ctx, userID, newEmail)

	if len(ret) == 0 {
		panic("no return value specified for CheckUserEmail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, userID, newEmail)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckUserPassword provides a mock function with given fields: ctx, newPass
func (_m *UserService) CheckUserPassword(ctx context.Context, newPass newPassword.NewPassword) error {
	ret := _m.Called(ctx, newPass)

	if len(ret) == 0 {
		panic("no return value specified for CheckUserPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, newPassword.NewPassword) error); ok {
		r0 = rf(ctx, newPass)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUserPfp provides a mock function with given fields: ctx, userId, image
func (_m *UserService) CreateUserPfp(ctx context.Context, userId string, image []byte) error {
	ret := _m.Called(ctx, userId, image)

	if len(ret) == 0 {
		panic("no return value specified for CreateUserPfp")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte) error); ok {
		r0 = rf(ctx, userId, image)
	} else {
		r0 = ret.Error(0)
	}

	return r0
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

// GetUserPfp provides a mock function with given fields: ctx, userId
func (_m *UserService) GetUserPfp(ctx context.Context, userId string) (string, error) {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for GetUserPfp")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, userId)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userId)
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

// UpdateUserPasswordByEmail provides a mock function with given fields: ctx, newPass
func (_m *UserService) UpdateUserPasswordByEmail(ctx context.Context, newPass newPassword.NewPassword) error {
	ret := _m.Called(ctx, newPass)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUserPasswordByEmail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, newPassword.NewPassword) error); ok {
		r0 = rf(ctx, newPass)
	} else {
		r0 = ret.Error(0)
	}

	return r0
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
