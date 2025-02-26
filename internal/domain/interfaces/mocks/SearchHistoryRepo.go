// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	searchhistory "github.com/imperatorofdwelling/Full-backend/internal/domain/models/searchhistory"
)

// SearchHistoryRepo is an autogenerated mock type for the SearchHistoryRepo type
type SearchHistoryRepo struct {
	mock.Mock
}

// AddHistory provides a mock function with given fields: ctx, userId, name
func (_m *SearchHistoryRepo) AddHistory(ctx context.Context, userId string, name string) error {
	ret := _m.Called(ctx, userId, name)

	if len(ret) == 0 {
		panic("no return value specified for AddHistory")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, userId, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllHistoryByUserId provides a mock function with given fields: ctx, userId
func (_m *SearchHistoryRepo) GetAllHistoryByUserId(ctx context.Context, userId string) ([]searchhistory.SearchHistory, error) {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for GetAllHistoryByUserId")
	}

	var r0 []searchhistory.SearchHistory
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]searchhistory.SearchHistory, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []searchhistory.SearchHistory); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]searchhistory.SearchHistory)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSearchHistoryRepo creates a new instance of SearchHistoryRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSearchHistoryRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *SearchHistoryRepo {
	mock := &SearchHistoryRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
