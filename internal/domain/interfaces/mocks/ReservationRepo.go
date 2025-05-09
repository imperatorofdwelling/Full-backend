// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	reservation "github.com/imperatorofdwelling/Full-backend/internal/domain/models/reservation"

	stays "github.com/imperatorofdwelling/Full-backend/internal/domain/models/stays"

	uuid "github.com/gofrs/uuid"
)

// ReservationRepo is an autogenerated mock type for the ReservationRepo type
type ReservationRepo struct {
	mock.Mock
}

// CheckIfArrivalIsCorrect provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *ReservationRepo) CheckIfArrivalIsCorrect(_a0 context.Context, _a1 uuid.UUID, _a2 uuid.UUID, _a3 reservation.ReservationCheckInEntity) (bool, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	if len(ret) == 0 {
		panic("no return value specified for CheckIfArrivalIsCorrect")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID, reservation.ReservationCheckInEntity) (bool, error)); ok {
		return rf(_a0, _a1, _a2, _a3)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID, reservation.ReservationCheckInEntity) bool); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, uuid.UUID, reservation.ReservationCheckInEntity) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckIfReservationExists provides a mock function with given fields: _a0, _a1
func (_m *ReservationRepo) CheckIfReservationExists(_a0 context.Context, _a1 uuid.UUID) (bool, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CheckIfReservationExists")
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

// CheckIfReservationExistsByStayID provides a mock function with given fields: _a0, _a1
func (_m *ReservationRepo) CheckIfReservationExistsByStayID(_a0 context.Context, _a1 uuid.UUID) (bool, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CheckIfReservationExistsByStayID")
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

// CheckIfUserIsOwner provides a mock function with given fields: _a0, _a1, _a2
func (_m *ReservationRepo) CheckIfUserIsOwner(_a0 context.Context, _a1 uuid.UUID, _a2 uuid.UUID) (bool, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for CheckIfUserIsOwner")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID) (bool, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID) bool); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckInApproval provides a mock function with given fields: _a0, _a1, _a2
func (_m *ReservationRepo) CheckInApproval(_a0 context.Context, _a1 uuid.UUID, _a2 reservation.ReservationCheckInEntity) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for CheckInApproval")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, reservation.ReservationCheckInEntity) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckReservationIfExistsByUserId provides a mock function with given fields: _a0, _a1
func (_m *ReservationRepo) CheckReservationIfExistsByUserId(_a0 context.Context, _a1 uuid.UUID) (bool, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CheckReservationIfExistsByUserId")
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

// CheckReservationIsFree provides a mock function with given fields: _a0, _a1
func (_m *ReservationRepo) CheckReservationIsFree(_a0 context.Context, _a1 *reservation.ReservationEntity) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CheckReservationIsFree")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *reservation.ReservationEntity) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckTimeForCheckOutReservation provides a mock function with given fields: _a0, _a1, _a2
func (_m *ReservationRepo) CheckTimeForCheckOutReservation(_a0 context.Context, _a1 string, _a2 string) (bool, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for CheckTimeForCheckOutReservation")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (bool, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConfirmCheckOutReservation provides a mock function with given fields: _a0, _a1, _a2
func (_m *ReservationRepo) ConfirmCheckOutReservation(_a0 context.Context, _a1 string, _a2 string) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for ConfirmCheckOutReservation")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateReservation provides a mock function with given fields: _a0, _a1, _a2
func (_m *ReservationRepo) CreateReservation(_a0 context.Context, _a1 *reservation.ReservationEntity, _a2 string) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for CreateReservation")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *reservation.ReservationEntity, string) error); ok {
		r0 = rf(_a0, _a1, _a2)
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

// GetFreeReservationsByUserID provides a mock function with given fields: ctx, id
func (_m *ReservationRepo) GetFreeReservationsByUserID(ctx context.Context, id uuid.UUID) (*[]stays.Stay, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetFreeReservationsByUserID")
	}

	var r0 *[]stays.Stay
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*[]stays.Stay, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *[]stays.Stay); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]stays.Stay)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOccupiedReservationsByUserID provides a mock function with given fields: ctx, id
func (_m *ReservationRepo) GetOccupiedReservationsByUserID(ctx context.Context, id uuid.UUID) (*[]stays.StayOccupied, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetOccupiedReservationsByUserID")
	}

	var r0 *[]stays.StayOccupied
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*[]stays.StayOccupied, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *[]stays.StayOccupied); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]stays.StayOccupied)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
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
