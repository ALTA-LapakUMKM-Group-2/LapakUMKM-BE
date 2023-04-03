// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	dashboards "lapakUmkm/features/dashboards"

	mock "github.com/stretchr/testify/mock"
)

// DashboardDataInterface is an autogenerated mock type for the DashboardDataInterface type
type DashboardDataInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: userId
func (_m *DashboardDataInterface) Create(userId uint) error {
	ret := _m.Called(userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectByUserId provides a mock function with given fields: id
func (_m *DashboardDataInterface) SelectByUserId(id uint) (dashboards.DashboardEntity, error) {
	ret := _m.Called(id)

	var r0 dashboards.DashboardEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (dashboards.DashboardEntity, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) dashboards.DashboardEntity); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(dashboards.DashboardEntity)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDashboardDataInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewDashboardDataInterface creates a new instance of DashboardDataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDashboardDataInterface(t mockConstructorTestingTNewDashboardDataInterface) *DashboardDataInterface {
	mock := &DashboardDataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
