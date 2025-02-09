// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

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

// CreateStay provides a mock function with given fields: _a0, _a1
func (_m *StaysRepo) CreateStay(_a0 context.Context, _a1 *stays.StayEntity) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *stays.StayEntity) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateStayImage provides a mock function with given fields: ctx, fileName, isMain, stayID
func (_m *StaysRepo) CreateStayImage(ctx context.Context, fileName string, isMain bool, stayID uuid.UUID) error {
	ret := _m.Called(ctx, fileName, isMain, stayID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, bool, uuid.UUID) error); ok {
		r0 = rf(ctx, fileName, isMain, stayID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteStayByID provides a mock function with given fields: _a0, _a1
func (_m *StaysRepo) DeleteStayByID(_a0 context.Context, _a1 uuid.UUID) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteStayImage provides a mock function with given fields: _a0, _a1
func (_m *StaysRepo) DeleteStayImage(_a0 context.Context, _a1 uuid.UUID) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetImagesByStayID provides a mock function with given fields: _a0, _a1
func (_m *StaysRepo) GetImagesByStayID(_a0 context.Context, _a1 uuid.UUID) ([]stays.StayImage, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []stays.StayImage
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []stays.StayImage); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]stays.StayImage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMainImageByStayID provides a mock function with given fields: _a0, _a1
func (_m *StaysRepo) GetMainImageByStayID(_a0 context.Context, _a1 uuid.UUID) (stays.StayImage, error) {
	ret := _m.Called(_a0, _a1)

	var r0 stays.StayImage
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) stays.StayImage); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(stays.StayImage)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStayByID provides a mock function with given fields: _a0, _a1
func (_m *StaysRepo) GetStayByID(_a0 context.Context, _a1 uuid.UUID) (*stays.Stay, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *stays.Stay
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *stays.Stay); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stays.Stay)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStayImageByID provides a mock function with given fields: _a0, _a1
func (_m *StaysRepo) GetStayImageByID(_a0 context.Context, _a1 uuid.UUID) (stays.StayImage, error) {
	ret := _m.Called(_a0, _a1)

	var r0 stays.StayImage
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) stays.StayImage); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(stays.StayImage)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStays provides a mock function with given fields: _a0
func (_m *StaysRepo) GetStays(_a0 context.Context) ([]stays.StayResponse, error) {
	ret := _m.Called(_a0)


	var r0 []*stays.Stay
	if rf, ok := ret.Get(0).(func(context.Context) []*stays.Stay); ok {

		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]stays.StayResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStaysByLocationID provides a mock function with given fields: _a0, _a1
func (_m *StaysRepo) GetStaysByLocationID(_a0 context.Context, _a1 uuid.UUID) (*[]stays.Stay, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *[]stays.Stay
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *[]stays.Stay); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]stays.Stay)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStaysByUserID provides a mock function with given fields: _a0, _a1
func (_m *StaysRepo) GetStaysByUserID(_a0 context.Context, _a1 uuid.UUID) ([]*stays.Stay, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*stays.Stay
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []*stays.Stay); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*stays.Stay)
		}
	}

	var r1 error
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

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *stays.StayEntity, uuid.UUID) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewStaysRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewStaysRepo creates a new instance of StaysRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStaysRepo(t mockConstructorTestingTNewStaysRepo) *StaysRepo {
	mock := &StaysRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
