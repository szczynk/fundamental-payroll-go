// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	model "fundamental-payroll-go/model"

	mock "github.com/stretchr/testify/mock"
)

// PayrollRepository is an autogenerated mock type for the PayrollRepository type
type PayrollRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: payroll
func (_m *PayrollRepository) Add(payroll *model.Payroll) (*model.Payroll, error) {
	ret := _m.Called(payroll)

	var r0 *model.Payroll
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.Payroll) (*model.Payroll, error)); ok {
		return rf(payroll)
	}
	if rf, ok := ret.Get(0).(func(*model.Payroll) *model.Payroll); ok {
		r0 = rf(payroll)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Payroll)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.Payroll) error); ok {
		r1 = rf(payroll)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Detail provides a mock function with given fields: id
func (_m *PayrollRepository) Detail(id int64) (*model.Payroll, error) {
	ret := _m.Called(id)

	var r0 *model.Payroll
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (*model.Payroll, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int64) *model.Payroll); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Payroll)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields:
func (_m *PayrollRepository) List() ([]model.Payroll, error) {
	ret := _m.Called()

	var r0 []model.Payroll
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.Payroll, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.Payroll); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Payroll)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPayrollRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewPayrollRepository creates a new instance of PayrollRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPayrollRepository(t mockConstructorTestingTNewPayrollRepository) *PayrollRepository {
	mock := &PayrollRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}