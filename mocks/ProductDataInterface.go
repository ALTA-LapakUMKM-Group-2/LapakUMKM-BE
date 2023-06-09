// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	products "lapakUmkm/features/products"

	mock "github.com/stretchr/testify/mock"
)

// ProductDataInterface is an autogenerated mock type for the ProductDataInterface type
type ProductDataInterface struct {
	mock.Mock
}

// Destroy provides a mock function with given fields: id
func (_m *ProductDataInterface) Destroy(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Edit provides a mock function with given fields: productEntity, id
func (_m *ProductDataInterface) Edit(productEntity products.ProductEntity, id uint) error {
	ret := _m.Called(productEntity, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(products.ProductEntity, uint) error); ok {
		r0 = rf(productEntity, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectAll provides a mock function with given fields: productFilter
func (_m *ProductDataInterface) SelectAll(productFilter products.ProductFilter) ([]products.ProductEntity, error) {
	ret := _m.Called(productFilter)

	var r0 []products.ProductEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(products.ProductFilter) ([]products.ProductEntity, error)); ok {
		return rf(productFilter)
	}
	if rf, ok := ret.Get(0).(func(products.ProductFilter) []products.ProductEntity); ok {
		r0 = rf(productFilter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]products.ProductEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(products.ProductFilter) error); ok {
		r1 = rf(productFilter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectById provides a mock function with given fields: id
func (_m *ProductDataInterface) SelectById(id uint) (products.ProductEntity, error) {
	ret := _m.Called(id)

	var r0 products.ProductEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (products.ProductEntity, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) products.ProductEntity); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(products.ProductEntity)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: productEntity
func (_m *ProductDataInterface) Store(productEntity products.ProductEntity) (uint, error) {
	ret := _m.Called(productEntity)

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(products.ProductEntity) (uint, error)); ok {
		return rf(productEntity)
	}
	if rf, ok := ret.Get(0).(func(products.ProductEntity) uint); ok {
		r0 = rf(productEntity)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(products.ProductEntity) error); ok {
		r1 = rf(productEntity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewProductDataInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewProductDataInterface creates a new instance of ProductDataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductDataInterface(t mockConstructorTestingTNewProductDataInterface) *ProductDataInterface {
	mock := &ProductDataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
