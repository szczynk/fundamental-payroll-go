package repository

import (
	"fundamental-payroll-go/helper/apperrors"
	"fundamental-payroll-go/model"
)

type employeeRepository struct{}

func NewEmployeeRepository() EmployeeRepository {
	return new(employeeRepository)
}

func (repo *employeeRepository) List() ([]model.Employee, error) {
	return model.Employees, nil
}

func (repo *employeeRepository) getLastID() int64 {
	employees, _ := repo.List()

	var tempID int64
	for _, employee := range employees {
		if tempID < employee.ID {
			tempID = employee.ID
		}
	}
	return tempID
}

func (repo *employeeRepository) Add(employee *model.Employee) (*model.Employee, error) {
	id := repo.getLastID()

	newEmployee := *employee
	newEmployee.ID = id + 1

	model.Employees = append(model.Employees, newEmployee)

	return &newEmployee, nil
}

func (repo *employeeRepository) getIndexByID(id int64) (int, error) {
	employees, _ := repo.List()

	for i, employee := range employees {
		if id == employee.ID {
			return i, nil
		}
	}

	return -1, apperrors.New(apperrors.ErrIdNotFound)
}

func (repo *employeeRepository) Detail(id int64) (*model.Employee, error) {
	employees, _ := repo.List()

	index, err := repo.getIndexByID(id)
	if err != nil {
		return nil, err
	}

	employee := employees[index]

	return &employee, nil
}
