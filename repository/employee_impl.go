package repository

import (
	"errors"
	"fundamental-payroll-go/model"
)

type employeeRepository struct{}

func NewEmployeeRepository() EmployeeRepository {
	return new(employeeRepository)
}

func (repo *employeeRepository) List() ([]model.Employee, error) {
	return model.Employees, nil
}

func (repo *employeeRepository) getLastID() (int64, error) {
	employees, err := repo.List()

	var tempID int64
	for _, employee := range employees {
		if tempID < employee.ID {
			tempID = employee.ID
		}
	}
	return tempID, err
}

func (repo *employeeRepository) Add(employee *model.Employee) (*model.Employee, error) {
	id, err := repo.getLastID()
	if err != nil {
		return &model.Employee{}, err
	}

	newEmployee := *employee
	newEmployee.ID = id + 1

	model.Employees = append(model.Employees, newEmployee)

	return &newEmployee, nil
}

func (repo *employeeRepository) getIndexByID(id int64) (int, error) {
	employees, err := repo.List()
	if err != nil {
		return -1, err
	}

	for i, employee := range employees {
		if id == employee.ID {
			return i, nil
		}
	}

	return -1, errors.New("ID tidak ditemukan")
}

func (repo *employeeRepository) Detail(id int64) (*model.Employee, error) {
	employees, err := repo.List()
	if err != nil {
		return &model.Employee{}, err
	}

	index, err := repo.getIndexByID(id)
	if err != nil {
		return &model.Employee{}, err
	}

	employee := employees[index]

	return &employee, nil
}
