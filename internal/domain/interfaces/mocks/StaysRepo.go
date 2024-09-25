// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	stays "github.com/imperatorofdwelling/Full-backend/internal/domain/models/stays"

	uuid "github.com/gofrs/uuid"
)

// StaysRepo is an autogenerated mock type for the StaysRepo type
type StaysRepo struct {
	mock.Mock
}

// CheckStayIfExistsByID provides a mock function with given fields: _a0, _a1
func (_m *StaysRepo) CheckStayIfExistsByID(_a0 context.Context, _a1 uuid.UUID) (bool, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CheckStayIfExistsByID")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (bool, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateStay provides a mock function with given fields: _a0, _a1
func (_m *StaysRepo) CreateStay(_a0 context.Context, _a1 *stays.StayEntity) error {
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

// DeleteStayByID provides a mock function with given fields: _a0, _a1
func (_m *StaysRepo) DeleteStayByID(_a0 context.Context, _a1 uuid.UUID) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DeleteStayByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetStayByID provides a mock function with given fields: _a0, _a1
func (_m *StaysRepo) GetStayByID(_a0 context.Context, _a1 uuid.UUID) (*stays.Stay, error) {
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
func (_m *StaysRepo) GetStays(_a0 context.Context) ([]*stays.Stay, error) {
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

// GetStaysByUserID provides a mock function with given fields: _a0, _a1
func (_m *StaysRepo) GetStaysByUserID(_a0 context.Context, _a1 uuid.UUID) ([]*stays.Stay, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetStaysByUserID")
	}

	var r0 []*stays.Stay
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]*stays.Stay, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []*stays.Stay); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*stays.Stay)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStayByID provides a mock function with given fields: _a0, _a1, _a2
func (_m *StaysRepo) UpdateStayByID(_a0 context.Context, _a1 *stays.StayEntity, _a2 uuid.UUID) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for UpdateStayByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *stays.StayEntity, uuid.UUID) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewStaysRepo creates a new instance of StaysRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStaysRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *StaysRepo {
	mock := &StaysRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}