// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	location "github.com/imperatorofdwelling/Full-backend/internal/domain/models/location"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/gofrs/uuid"
)

// LocationRepo is an autogenerated mock type for the LocationRepo type
type LocationRepo struct {
	mock.Mock
}

// DeleteByID provides a mock function with given fields: ctx, id
func (_m *LocationRepo) DeleteByID(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByNameMatch provides a mock function with given fields: ctx, match
func (_m *LocationRepo) FindByNameMatch(ctx context.Context, match string) (*[]location.Location, error) {
	ret := _m.Called(ctx, match)

	var r0 *[]location.Location
	if rf, ok := ret.Get(0).(func(context.Context, string) *[]location.Location); ok {
		r0 = rf(ctx, match)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]location.Location)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, match)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx
func (_m *LocationRepo) GetAll(ctx context.Context) (*[]location.Location, error) {
	ret := _m.Called(ctx)

	var r0 *[]location.Location
	if rf, ok := ret.Get(0).(func(context.Context) *[]location.Location); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]location.Location)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *LocationRepo) GetByID(ctx context.Context, id uuid.UUID) (*location.Location, error) {
	ret := _m.Called(ctx, id)

	var r0 *location.Location
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *location.Location); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*location.Location)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateByID provides a mock function with given fields: ctx, id, _a2
func (_m *LocationRepo) UpdateByID(ctx context.Context, id uuid.UUID, _a2 location.LocationEntity) error {
	ret := _m.Called(ctx, id, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, location.LocationEntity) error); ok {
		r0 = rf(ctx, id, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewLocationRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewLocationRepo creates a new instance of LocationRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLocationRepo(t mockConstructorTestingTNewLocationRepo) *LocationRepo {
	mock := &LocationRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
