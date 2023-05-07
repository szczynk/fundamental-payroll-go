//go:generate mockery --output=../mocks --name EmployeeRepository
package repository

import "fundamental-payroll-go/model"

type EmployeeRepository interface {
	List() ([]model.Employee, error)
	Add(req *model.Employee) (*model.Employee, error)
	Detail(id int64) (*model.Employee, error)
}
