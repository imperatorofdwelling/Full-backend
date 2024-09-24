// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	stays "github.com/imperatorofdwelling/Full-backend/internal/domain/models/stays"

	uuid "github.com/gofrs/uuid"
)

// StaysService is an autogenerated mock type for the StaysService type
type StaysService struct {
	mock.Mock
}

// CreateStay provides a mock function with given fields: _a0, _a1
func (_m *StaysService) CreateStay(_a0 context.Context, _a1 *stays.StayEntity) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateStay")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *stays.StayEntity) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetStayByID provides a mock function with given fields: _a0, _a1
func (_m *StaysService) GetStayByID(_a0 context.Context, _a1 uuid.UUID) (*stays.Stay, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetStayByID")
	}

	var r0 *stays.Stay
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*stays.Stay, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *stays.Stay); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stays.Stay)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStays provides a mock function with given fields: _a0
func (_m *StaysService) GetStays(_a0 context.Context) ([]*stays.Stay, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetStays")
	}

	var r0 []*stays.Stay
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*stays.Stay, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*stays.Stay); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*stays.Stay)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewStaysService creates a new instance of StaysService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStaysService(t interface {
	mock.TestingT
	Cleanup(func())
}) *StaysService {
	mock := &StaysService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
