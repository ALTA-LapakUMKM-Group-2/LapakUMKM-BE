// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	productTransactionDetails "lapakUmkm/features/productTransactionDetails"

	mock "github.com/stretchr/testify/mock"
)

// ProductTransactionDetailServiceInterface is an autogenerated mock type for the ProductTransactionDetailServiceInterface type
type ProductTransactionDetailServiceInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: productTransactionDetailEntity
func (_m *ProductTransactionDetailServiceInterface) Create(productTransactionDetailEntity productTransactionDetails.ProductTransactionDetailEntity) (productTransactionDetails.ProductTransactionDetailEntity, error) {
	ret := _m.Called(productTransactionDetailEntity)

	var r0 productTransactionDetails.ProductTransactionDetailEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(productTransactionDetails.ProductTransactionDetailEntity) (productTransactionDetails.ProductTransactionDetailEntity, error)); ok {
		return rf(productTransactionDetailEntity)
	}
	if rf, ok := ret.Get(0).(func(productTransactionDetails.ProductTransactionDetailEntity) productTransactionDetails.ProductTransactionDetailEntity); ok {
		r0 = rf(productTransactionDetailEntity)
	} else {
		r0 = ret.Get(0).(productTransactionDetails.ProductTransactionDetailEntity)
	}

	if rf, ok := ret.Get(1).(func(productTransactionDetails.ProductTransactionDetailEntity) error); ok {
		r1 = rf(productTransactionDetailEntity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *ProductTransactionDetailServiceInterface) GetById(id uint) (productTransactionDetails.ProductTransactionDetailEntity, error) {
	ret := _m.Called(id)

	var r0 productTransactionDetails.ProductTransactionDetailEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (productTransactionDetails.ProductTransactionDetailEntity, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) productTransactionDetails.ProductTransactionDetailEntity); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(productTransactionDetails.ProductTransactionDetailEntity)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByTransaksiId provides a mock function with given fields: transaksiId
func (_m *ProductTransactionDetailServiceInterface) GetByTransaksiId(transaksiId uint) ([]productTransactionDetails.ProductTransactionDetailEntity, error) {
	ret := _m.Called(transaksiId)

	var r0 []productTransactionDetails.ProductTransactionDetailEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]productTransactionDetails.ProductTransactionDetailEntity, error)); ok {
		return rf(transaksiId)
	}
	if rf, ok := ret.Get(0).(func(uint) []productTransactionDetails.ProductTransactionDetailEntity); ok {
		r0 = rf(transaksiId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]productTransactionDetails.ProductTransactionDetailEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(transaksiId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewProductTransactionDetailServiceInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewProductTransactionDetailServiceInterface creates a new instance of ProductTransactionDetailServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductTransactionDetailServiceInterface(t mockConstructorTestingTNewProductTransactionDetailServiceInterface) *ProductTransactionDetailServiceInterface {
	mock := &ProductTransactionDetailServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
