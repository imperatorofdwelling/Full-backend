// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	context "context"

	message "github.com/imperatorofdwelling/Full-backend/internal/domain/models/message"

	mock "github.com/stretchr/testify/mock"
)

// MessageService is an autogenerated mock type for the MessageService type
type MessageService struct {
	mock.Mock
}

// DeleteMessageByID provides a mock function with given fields: ctx, messageId
func (_m *MessageService) DeleteMessageByID(ctx context.Context, messageId string) error {
	ret := _m.Called(ctx, messageId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteMessageByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, messageId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetMessageByMessageID provides a mock function with given fields: ctx, messageId
func (_m *MessageService) GetMessageByMessageID(ctx context.Context, messageId string) (*message.Message, error) {
	ret := _m.Called(ctx, messageId)

	if len(ret) == 0 {
		panic("no return value specified for GetMessageByMessageID")
	}

	var r0 *message.Message
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*message.Message, error)); ok {
		return rf(ctx, messageId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *message.Message); ok {
		r0 = rf(ctx, messageId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*message.Message)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, messageId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMessagesByUserID provides a mock function with given fields: ctx, userId
func (_m *MessageService) GetMessagesByUserID(ctx context.Context, userId string) ([]*message.Entity, error) {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for GetMessagesByUserID")
	}

	var r0 []*message.Entity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*message.Entity, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*message.Entity); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*message.Entity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMessageByID provides a mock function with given fields: ctx, messageId, msg
func (_m *MessageService) UpdateMessageByID(ctx context.Context, messageId string, msg message.Entity) (*message.Message, error) {
	ret := _m.Called(ctx, messageId, msg)

	if len(ret) == 0 {
		panic("no return value specified for UpdateMessageByID")
	}

	var r0 *message.Message
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, message.Entity) (*message.Message, error)); ok {
		return rf(ctx, messageId, msg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, message.Entity) *message.Message); ok {
		r0 = rf(ctx, messageId, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*message.Message)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, message.Entity) error); ok {
		r1 = rf(ctx, messageId, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMessageService creates a new instance of MessageService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMessageService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MessageService {
	mock := &MessageService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
