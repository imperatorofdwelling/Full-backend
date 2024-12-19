// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	context "context"

	advantage "github.com/imperatorofdwelling/Full-backend/internal/domain/models/advantage"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/gofrs/uuid"
)

// AdvantageRepo is an autogenerated mock type for the AdvantageRepo type
type AdvantageRepo struct {
	mock.Mock
}

// CheckAdvantageIfExists provides a mock function with given fields: ctx, advName
func (_m *AdvantageRepo) CheckAdvantageIfExists(ctx context.Context, advName string) (bool, error) {
	ret := _m.Called(ctx, advName)

	if len(ret) == 0 {
		panic("no return value specified for CheckAdvantageIfExists")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, advName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, advName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, advName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateAdvantage provides a mock function with given fields: ctx, advTitle, imgPath
func (_m *AdvantageRepo) CreateAdvantage(ctx context.Context, advTitle string, imgPath string) error {
	ret := _m.Called(ctx, advTitle, imgPath)

	if len(ret) == 0 {
		panic("no return value specified for CreateAdvantage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, advTitle, imgPath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAdvantageByID provides a mock function with given fields: ctx, id
func (_m *AdvantageRepo) FindAdvantageByID(ctx context.Context, id uuid.UUID) (*advantage.Advantage, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindAdvantageByID")
	}

	var r0 *advantage.Advantage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*advantage.Advantage, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *advantage.Advantage); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*advantage.Advantage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllAdvantages provides a mock function with given fields: _a0
func (_m *AdvantageRepo) GetAllAdvantages(_a0 context.Context) ([]advantage.Advantage, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetAllAdvantages")
	}

	var r0 []advantage.Advantage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]advantage.Advantage, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []advantage.Advantage); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]advantage.Advantage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveAdvantage provides a mock function with given fields: ctx, id
func (_m *AdvantageRepo) RemoveAdvantage(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for RemoveAdvantage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateAdvantageByID provides a mock function with given fields: ctx, id, adv
func (_m *AdvantageRepo) UpdateAdvantageByID(ctx context.Context, id uuid.UUID, adv *advantage.Advantage) error {
	ret := _m.Called(ctx, id, adv)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAdvantageByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, *advantage.Advantage) error); ok {
		r0 = rf(ctx, id, adv)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewAdvantageRepo creates a new instance of AdvantageRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAdvantageRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *AdvantageRepo {
	mock := &AdvantageRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
