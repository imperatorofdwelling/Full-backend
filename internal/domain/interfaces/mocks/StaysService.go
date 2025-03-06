// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	multipart "mime/multipart"

	stays "github.com/imperatorofdwelling/Full-backend/internal/domain/models/stays"

	uuid "github.com/gofrs/uuid"
)

// StaysService is an autogenerated mock type for the StaysService type
type StaysService struct {
	mock.Mock
}

// CreateImages provides a mock function with given fields: _a0, _a1, _a2
func (_m *StaysService) CreateImages(_a0 context.Context, _a1 []*multipart.FileHeader, _a2 uuid.UUID) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for CreateImages")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*multipart.FileHeader, uuid.UUID) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateMainImage provides a mock function with given fields: _a0, _a1, _a2
func (_m *StaysService) CreateMainImage(_a0 context.Context, _a1 *multipart.FileHeader, _a2 uuid.UUID) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for CreateMainImage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *multipart.FileHeader, uuid.UUID) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
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

// DeleteStayByID provides a mock function with given fields: _a0, _a1
func (_m *StaysService) DeleteStayByID(_a0 context.Context, _a1 uuid.UUID) error {
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

// DeleteStayImage provides a mock function with given fields: _a0, _a1
func (_m *StaysService) DeleteStayImage(_a0 context.Context, _a1 uuid.UUID) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DeleteStayImage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetImagesByStayID provides a mock function with given fields: _a0, _a1
func (_m *StaysService) GetImagesByStayID(_a0 context.Context, _a1 uuid.UUID) ([]stays.StayImage, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetImagesByStayID")
	}

	var r0 []stays.StayImage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]stays.StayImage, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []stays.StayImage); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]stays.StayImage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMainImageByStayID provides a mock function with given fields: _a0, _a1
func (_m *StaysService) GetMainImageByStayID(_a0 context.Context, _a1 uuid.UUID) (stays.StayImage, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetMainImageByStayID")
	}

	var r0 stays.StayImage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (stays.StayImage, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) stays.StayImage); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(stays.StayImage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
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
func (_m *StaysService) GetStays(_a0 context.Context) ([]stays.StayResponse, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetStays")
	}

	var r0 []stays.StayResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]stays.StayResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []stays.StayResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]stays.StayResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStaysByLocationID provides a mock function with given fields: _a0, _a1
func (_m *StaysService) GetStaysByLocationID(_a0 context.Context, _a1 uuid.UUID) (*[]stays.Stay, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetStaysByLocationID")
	}

	var r0 *[]stays.Stay
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*[]stays.Stay, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *[]stays.Stay); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]stays.Stay)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStaysByUserID provides a mock function with given fields: _a0, _a1
func (_m *StaysService) GetStaysByUserID(_a0 context.Context, _a1 uuid.UUID) ([]*stays.Stay, error) {
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

// Filtration provides a mock function with given fields: ctx, search
func (_m *StaysService) Filtration(ctx context.Context, search stays.Filtration) ([]stays.Stay, error) {
	ret := _m.Called(ctx, search)

	if len(ret) == 0 {
		panic("no return value specified for Filtration")
	}

	var r0 []stays.Stay
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, stays.Filtration) ([]stays.Stay, error)); ok {
		return rf(ctx, search)
	}
	if rf, ok := ret.Get(0).(func(context.Context, stays.Filtration) []stays.Stay); ok {
		r0 = rf(ctx, search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]stays.Stay)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, stays.Filtration) error); ok {
		r1 = rf(ctx, search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStayByID provides a mock function with given fields: _a0, _a1, _a2
func (_m *StaysService) UpdateStayByID(_a0 context.Context, _a1 *stays.StayEntity, _a2 uuid.UUID) (*stays.Stay, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for UpdateStayByID")
	}

	var r0 *stays.Stay
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *stays.StayEntity, uuid.UUID) (*stays.Stay, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *stays.StayEntity, uuid.UUID) *stays.Stay); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stays.Stay)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *stays.StayEntity, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1, _a2)
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
