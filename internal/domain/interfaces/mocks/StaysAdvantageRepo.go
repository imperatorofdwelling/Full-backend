// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	staysadvantage "github.com/imperatorofdwelling/Full-backend/internal/domain/models/staysadvantage"

	uuid "github.com/gofrs/uuid"
)

// StaysAdvantageRepo is an autogenerated mock type for the StaysAdvantageRepo type
type StaysAdvantageRepo struct {
	mock.Mock
}

// CheckStaysAdvantageIfExists provides a mock function with given fields: _a0, _a1
func (_m *StaysAdvantageRepo) CheckStaysAdvantageIfExists(_a0 context.Context, _a1 uuid.UUID) (bool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateStaysAdvantage provides a mock function with given fields: ctx, stayAdv
func (_m *StaysAdvantageRepo) CreateStaysAdvantage(ctx context.Context, stayAdv *staysadvantage.StayAdvantageEntity) error {
	ret := _m.Called(ctx, stayAdv)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *staysadvantage.StayAdvantageEntity) error); ok {
		r0 = rf(ctx, stayAdv)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteStaysAdvantageByID provides a mock function with given fields: _a0, _a1
func (_m *StaysAdvantageRepo) DeleteStaysAdvantageByID(_a0 context.Context, _a1 uuid.UUID) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewStaysAdvantageRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewStaysAdvantageRepo creates a new instance of StaysAdvantageRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStaysAdvantageRepo(t mockConstructorTestingTNewStaysAdvantageRepo) *StaysAdvantageRepo {
	mock := &StaysAdvantageRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
