package repository

import "fundamental-payroll-go/model"

type EmployeeRepository interface {
	List() []model.Employee
	Add(req *model.Employee) (*model.Employee, error)
	Detail(id int64) (*model.Employee, error)
}
