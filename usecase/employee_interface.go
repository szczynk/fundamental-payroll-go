package usecase

import "fundamental-payroll-go/model"

type EmployeeUsecase interface {
	List() []model.Employee
	Add(req *model.EmployeeRequest) (*model.Employee, error)
	Detail(id int64) (*model.Employee, error)
}
