// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	carts "lapakUmkm/features/carts"

	mock "github.com/stretchr/testify/mock"
)

// CartData is an autogenerated mock type for the CartData type
type CartData struct {
	mock.Mock
}

// Add provides a mock function with given fields: newCart
func (_m *CartData) Add(newCart carts.Core) (carts.Core, error) {
	ret := _m.Called(newCart)

	var r0 carts.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(carts.Core) (carts.Core, error)); ok {
		return rf(newCart)
	}
	if rf, ok := ret.Get(0).(func(carts.Core) carts.Core); ok {
		r0 = rf(newCart)
	} else {
		r0 = ret.Get(0).(carts.Core)
	}

	if rf, ok := ret.Get(1).(func(carts.Core) error); ok {
		r1 = rf(newCart)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: userID, cartID
func (_m *CartData) Delete(userID uint, cartID uint) error {
	ret := _m.Called(userID, cartID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(userID, cartID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MyCart provides a mock function with given fields: userID
func (_m *CartData) MyCart(userID uint) ([]carts.Core, error) {
	ret := _m.Called(userID)

	var r0 []carts.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]carts.Core, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint) []carts.Core); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]carts.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: updateCart
func (_m *CartData) Update(updateCart carts.Core) (carts.Core, error) {
	ret := _m.Called(updateCart)

	var r0 carts.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(carts.Core) (carts.Core, error)); ok {
		return rf(updateCart)
	}
	if rf, ok := ret.Get(0).(func(carts.Core) carts.Core); ok {
		r0 = rf(updateCart)
	} else {
		r0 = ret.Get(0).(carts.Core)
	}

	if rf, ok := ret.Get(1).(func(carts.Core) error); ok {
		r1 = rf(updateCart)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCartData interface {
	mock.TestingT
	Cleanup(func())
}

// NewCartData creates a new instance of CartData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCartData(t mockConstructorTestingTNewCartData) *CartData {
	mock := &CartData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
