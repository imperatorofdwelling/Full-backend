// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	context "context"

	auth "github.com/imperatorofdwelling/Full-backend/internal/domain/models/auth"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/gofrs/uuid"
)

// AuthService is an autogenerated mock type for the AuthService type
type AuthService struct {
	mock.Mock
}

// CheckEmailChangeOTP provides a mock function with given fields: ctx, userID, otp
func (_m *AuthService) CheckEmailChangeOTP(ctx context.Context, userID string, otp string) error {
	ret := _m.Called(ctx, userID, otp)

	if len(ret) == 0 {
		panic("no return value specified for CheckEmailChangeOTP")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, userID, otp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckEmailOTP provides a mock function with given fields: ctx, userID, otp
func (_m *AuthService) CheckEmailOTP(ctx context.Context, userID string, otp string) error {
	ret := _m.Called(ctx, userID, otp)

	if len(ret) == 0 {
		panic("no return value specified for CheckEmailOTP")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, userID, otp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckPasswordOTP provides a mock function with given fields: ctx, email, otp
func (_m *AuthService) CheckPasswordOTP(ctx context.Context, email string, otp string) error {
	ret := _m.Called(ctx, email, otp)

	if len(ret) == 0 {
		panic("no return value specified for CheckPasswordOTP")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, email, otp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Login provides a mock function with given fields: ctx, user
func (_m *AuthService) Login(ctx context.Context, user auth.Login) (uuid.UUID, int, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 uuid.UUID
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, auth.Login) (uuid.UUID, int, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, auth.Login) uuid.UUID); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, auth.Login) int); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, auth.Login) error); ok {
		r2 = rf(ctx, user)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Register provides a mock function with given fields: ctx, user
func (_m *AuthService) Register(ctx context.Context, user auth.Registration) (uuid.UUID, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, auth.Registration) (uuid.UUID, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, auth.Registration) uuid.UUID); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, auth.Registration) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthService creates a new instance of AuthService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthService(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthService {
	mock := &AuthService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
