// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	staysreviews "github.com/imperatorofdwelling/Full-backend/internal/domain/models/staysreviews"

	uuid "github.com/gofrs/uuid"
)

// StaysReviewsService is an autogenerated mock type for the StaysReviewsService type
type StaysReviewsService struct {
	mock.Mock
}

// CreateStaysReview provides a mock function with given fields: _a0, _a1
func (_m *StaysReviewsService) CreateStaysReview(_a0 context.Context, _a1 *staysreviews.StaysReviewEntity) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateStaysReview")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *staysreviews.StaysReviewEntity) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteStaysReview provides a mock function with given fields: _a0, _a1
func (_m *StaysReviewsService) DeleteStaysReview(_a0 context.Context, _a1 uuid.UUID) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DeleteStaysReview")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAllStaysReviews provides a mock function with given fields: _a0
func (_m *StaysReviewsService) FindAllStaysReviews(_a0 context.Context) ([]staysreviews.StaysReview, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for FindAllStaysReviews")
	}

	var r0 []staysreviews.StaysReview
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]staysreviews.StaysReview, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []staysreviews.StaysReview); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]staysreviews.StaysReview)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOneStaysReview provides a mock function with given fields: _a0, _a1
func (_m *StaysReviewsService) FindOneStaysReview(_a0 context.Context, _a1 uuid.UUID) (*staysreviews.StaysReview, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for FindOneStaysReview")
	}

	var r0 *staysreviews.StaysReview
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*staysreviews.StaysReview, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *staysreviews.StaysReview); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*staysreviews.StaysReview)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStaysReview provides a mock function with given fields: _a0, _a1, _a2
func (_m *StaysReviewsService) UpdateStaysReview(_a0 context.Context, _a1 *staysreviews.StaysReviewEntity, _a2 uuid.UUID) (*staysreviews.StaysReview, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for UpdateStaysReview")
	}

	var r0 *staysreviews.StaysReview
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *staysreviews.StaysReviewEntity, uuid.UUID) (*staysreviews.StaysReview, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *staysreviews.StaysReviewEntity, uuid.UUID) *staysreviews.StaysReview); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*staysreviews.StaysReview)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *staysreviews.StaysReviewEntity, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewStaysReviewsService creates a new instance of StaysReviewsService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStaysReviewsService(t interface {
	mock.TestingT
	Cleanup(func())
}) *StaysReviewsService {
	mock := &StaysReviewsService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
