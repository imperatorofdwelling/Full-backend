// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	context "context"

	advantage "github.com/imperatorofdwelling/Full-backend/internal/domain/models/advantage"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/gofrs/uuid"
)

// AdvantageService is an autogenerated mock type for the AdvantageService type
type AdvantageService struct {
	mock.Mock
}

// CreateAdvantage provides a mock function with given fields: _a0, _a1
func (_m *AdvantageService) CreateAdvantage(_a0 context.Context, _a1 *advantage.AdvantageEntity) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateAdvantage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *advantage.AdvantageEntity) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAdvantageByID provides a mock function with given fields: _a0, _a1
func (_m *AdvantageService) GetAdvantageByID(_a0 context.Context, _a1 uuid.UUID) (*advantage.Advantage, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetAdvantageByID")
	}

	var r0 *advantage.Advantage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*advantage.Advantage, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *advantage.Advantage); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*advantage.Advantage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllAdvantages provides a mock function with given fields: _a0
func (_m *AdvantageService) GetAllAdvantages(_a0 context.Context) ([]advantage.Advantage, error) {
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

// RemoveAdvantage provides a mock function with given fields: _a0, _a1
func (_m *AdvantageService) RemoveAdvantage(_a0 context.Context, _a1 uuid.UUID) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for RemoveAdvantage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateAdvantageByID provides a mock function with given fields: _a0, _a1, _a2
func (_m *AdvantageService) UpdateAdvantageByID(_a0 context.Context, _a1 uuid.UUID, _a2 *advantage.AdvantageEntity) (advantage.Advantage, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAdvantageByID")
	}

	var r0 advantage.Advantage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, *advantage.AdvantageEntity) (advantage.Advantage, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, *advantage.AdvantageEntity) advantage.Advantage); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(advantage.Advantage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, *advantage.AdvantageEntity) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAdvantageService creates a new instance of AdvantageService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAdvantageService(t interface {
	mock.TestingT
	Cleanup(func())
}) *AdvantageService {
	mock := &AdvantageService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
