// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	reservation "github.com/imperatorofdwelling/Full-backend/internal/domain/models/reservation"

	uuid "github.com/gofrs/uuid"
)

// ReservationRepo is an autogenerated mock type for the ReservationRepo type
type ReservationRepo struct {
	mock.Mock
}

// CheckReservationIfExists provides a mock function with given fields: _a0, _a1
func (_m *ReservationRepo) CheckReservationIfExists(_a0 context.Context, _a1 uuid.UUID) (bool, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CheckReservationIfExists")
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

// CreateReservation provides a mock function with given fields: _a0, _a1
func (_m *ReservationRepo) CreateReservation(_a0 context.Context, _a1 *reservation.ReservationEntity) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateReservation")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *reservation.ReservationEntity) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteReservationByID provides a mock function with given fields: _a0, _a1
func (_m *ReservationRepo) DeleteReservationByID(_a0 context.Context, _a1 uuid.UUID) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DeleteReservationByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllReservationsByUserID provides a mock function with given fields: _a0, _a1
func (_m *ReservationRepo) GetAllReservationsByUserID(_a0 context.Context, _a1 uuid.UUID) (*[]reservation.Reservation, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetAllReservationsByUserID")
	}

	var r0 *[]reservation.Reservation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*[]reservation.Reservation, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *[]reservation.Reservation); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]reservation.Reservation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReservationByID provides a mock function with given fields: _a0, _a1
func (_m *ReservationRepo) GetReservationByID(_a0 context.Context, _a1 uuid.UUID) (*reservation.Reservation, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetReservationByID")
	}

	var r0 *reservation.Reservation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*reservation.Reservation, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *reservation.Reservation); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*reservation.Reservation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateReservationByID provides a mock function with given fields: _a0, _a1
func (_m *ReservationRepo) UpdateReservationByID(_a0 context.Context, _a1 *reservation.ReservationUpdateEntity) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateReservationByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *reservation.ReservationUpdateEntity) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewReservationRepo creates a new instance of ReservationRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReservationRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *ReservationRepo {
	mock := &ReservationRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
