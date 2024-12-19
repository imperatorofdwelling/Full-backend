// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	user "github.com/imperatorofdwelling/Full-backend/internal/domain/models/user"

	uuid "github.com/gofrs/uuid"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// CheckUserExists provides a mock function with given fields: ctx, email
func (_m *UserRepository) CheckUserExists(ctx context.Context, email string) (bool, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for CheckUserExists")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUserByID provides a mock function with given fields: ctx, id
func (_m *UserRepository) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUserByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindUserByID provides a mock function with given fields: ctx, id
func (_m *UserRepository) FindUserByID(ctx context.Context, id uuid.UUID) (user.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindUserByID")
	}

	var r0 user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (user.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) user.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserByID provides a mock function with given fields: ctx, id, _a2
func (_m *UserRepository) UpdateUserByID(ctx context.Context, id uuid.UUID, _a2 user.User) error {
	ret := _m.Called(ctx, id, _a2)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUserByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, user.User) error); ok {
		r0 = rf(ctx, id, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
