// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ConfirmEmailService is an autogenerated mock type for the ConfirmEmailService type
type ConfirmEmailService struct {
	mock.Mock
}

// CreateOTP provides a mock function with given fields: ctx, userID
func (_m *ConfirmEmailService) CreateOTP(ctx context.Context, userID string) error {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for CreateOTP")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewConfirmEmailService creates a new instance of ConfirmEmailService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConfirmEmailService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ConfirmEmailService {
	mock := &ConfirmEmailService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
