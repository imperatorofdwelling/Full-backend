// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// PaymentHandler is an autogenerated mock type for the PaymentHandler type
type PaymentHandler struct {
	mock.Mock
}

// MakePayment provides a mock function with given fields: w, r
func (_m *PaymentHandler) MakePayment(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// NewPaymentHandler creates a new instance of PaymentHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPaymentHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *PaymentHandler {
	mock := &PaymentHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
