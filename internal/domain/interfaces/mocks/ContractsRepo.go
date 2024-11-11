// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	contracts "github.com/imperatorofdwelling/Full-backend/internal/domain/models/contracts"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// ContractsRepo is an autogenerated mock type for the ContractsRepo type
type ContractsRepo struct {
	mock.Mock
}

// AddContract provides a mock function with given fields: ctx, userId, stayId, dateStart, dateEnd
func (_m *ContractsRepo) AddContract(ctx context.Context, userId string, stayId string, dateStart time.Time, dateEnd time.Time) error {
	ret := _m.Called(ctx, userId, stayId, dateStart, dateEnd)

	if len(ret) == 0 {
		panic("no return value specified for AddContract")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, time.Time, time.Time) error); ok {
		r0 = rf(ctx, userId, stayId, dateStart, dateEnd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllContracts provides a mock function with given fields: ctx, userId
func (_m *ContractsRepo) GetAllContracts(ctx context.Context, userId string) ([]contracts.ContractEntity, error) {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for GetAllContracts")
	}

	var r0 []contracts.ContractEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]contracts.ContractEntity, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []contracts.ContractEntity); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]contracts.ContractEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateContract provides a mock function with given fields: ctx, userId, stayId, dateStart, dateEnd
func (_m *ContractsRepo) UpdateContract(ctx context.Context, userId string, stayId string, dateStart time.Time, dateEnd time.Time) (*contracts.ContractEntity, error) {
	ret := _m.Called(ctx, userId, stayId, dateStart, dateEnd)

	if len(ret) == 0 {
		panic("no return value specified for UpdateContract")
	}

	var r0 *contracts.ContractEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, time.Time, time.Time) (*contracts.ContractEntity, error)); ok {
		return rf(ctx, userId, stayId, dateStart, dateEnd)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, time.Time, time.Time) *contracts.ContractEntity); ok {
		r0 = rf(ctx, userId, stayId, dateStart, dateEnd)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contracts.ContractEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, time.Time, time.Time) error); ok {
		r1 = rf(ctx, userId, stayId, dateStart, dateEnd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewContractsRepo creates a new instance of ContractsRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewContractsRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *ContractsRepo {
	mock := &ContractsRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}