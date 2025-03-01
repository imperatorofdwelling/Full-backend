// Code generated by mockery v2.46.3. DO NOT EDIT.

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

// CreateUserPfp provides a mock function with given fields: ctx, userId, imagePath
func (_m *UserRepository) CreateUserPfp(ctx context.Context, userId string, imagePath string) error {
	ret := _m.Called(ctx, userId, imagePath)

	if len(ret) == 0 {
		panic("no return value specified for CreateUserPfp")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, userId, imagePath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
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

// DeleteUserPfp provides a mock function with given fields: ctx, userId
func (_m *UserRepository) DeleteUserPfp(ctx context.Context, userId uuid.UUID) error {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUserPfp")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, userId)
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

// GetUserIDByEmail provides a mock function with given fields: ctx, email
func (_m *UserRepository) GetUserIDByEmail(ctx context.Context, email string) (string, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserIDByEmail")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserPasswordByEmail provides a mock function with given fields: ctx, email
func (_m *UserRepository) GetUserPasswordByEmail(ctx context.Context, email string) (string, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserPasswordByEmail")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserPfp provides a mock function with given fields: ctx, userId
func (_m *UserRepository) GetUserPfp(ctx context.Context, userId string) (string, error) {
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

// UpdateUserEmailByID provides a mock function with given fields: ctx, id, newEmail
func (_m *UserRepository) UpdateUserEmailByID(ctx context.Context, id uuid.UUID, newEmail string) error {
	ret := _m.Called(ctx, id, newEmail)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUserEmailByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string) error); ok {
		r0 = rf(ctx, id, newEmail)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUserPasswordByID provides a mock function with given fields: ctx, id, newPassword
func (_m *UserRepository) UpdateUserPasswordByID(ctx context.Context, id uuid.UUID, newPassword string) error {
	ret := _m.Called(ctx, id, newPassword)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUserPasswordByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string) error); ok {
		r0 = rf(ctx, id, newPassword)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUserPfp provides a mock function with given fields: ctx, userId, imagePath
func (_m *UserRepository) UpdateUserPfp(ctx context.Context, userId uuid.UUID, imagePath string) error {
	ret := _m.Called(ctx, userId, imagePath)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUserPfp")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string) error); ok {
		r0 = rf(ctx, userId, imagePath)
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
