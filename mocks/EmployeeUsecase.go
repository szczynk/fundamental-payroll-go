// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	model "fundamental-payroll-go/model"

	mock "github.com/stretchr/testify/mock"
)

// EmployeeUsecase is an autogenerated mock type for the EmployeeUsecase type
type EmployeeUsecase struct {
	mock.Mock
}

// Add provides a mock function with given fields: req
func (_m *EmployeeUsecase) Add(req *model.EmployeeRequest) (*model.Employee, error) {
	ret := _m.Called(req)

	var r0 *model.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.EmployeeRequest) (*model.Employee, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*model.EmployeeRequest) *model.Employee); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Employee)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.EmployeeRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Detail provides a mock function with given fields: id
func (_m *EmployeeUsecase) Detail(id int64) (*model.Employee, error) {
	ret := _m.Called(id)

	var r0 *model.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (*model.Employee, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int64) *model.Employee); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Employee)
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
func (_m *EmployeeUsecase) List() ([]model.Employee, error) {
	ret := _m.Called()

	var r0 []model.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.Employee, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.Employee); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Employee)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewEmployeeUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewEmployeeUsecase creates a new instance of EmployeeUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEmployeeUsecase(t mockConstructorTestingTNewEmployeeUsecase) *EmployeeUsecase {
	mock := &EmployeeUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
