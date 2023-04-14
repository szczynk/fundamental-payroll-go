package repository

import (
	"encoding/json"
	"fundamental-payroll-go/helper/apperrors"
	"fundamental-payroll-go/model"
	"os"
)

const employeeJsonFile = "data/employee.json"

type employeeJsonRepository struct{}

func NewEmployeeJsonRepository() EmployeeRepository {
	return new(employeeJsonRepository)
}

func (repo *employeeJsonRepository) encodeJSON(path string, employees *[]model.Employee) error {
	writer, err := os.Create(path)
	if err != nil {
		return err
	}
	defer writer.Close()

	encoder := json.NewEncoder(writer)
	err = encoder.Encode(&employees)
	if err != nil {
		return err
	}
	return nil
}

func (repo *employeeJsonRepository) decodeJSON(path string, employees *[]model.Employee) error {
	reader, err := os.Open(path)
	if err != nil {
		return err
	}
	defer reader.Close()

	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&employees)
	if err != nil {
		return err
	}
	return nil
}

func (repo *employeeJsonRepository) List() ([]model.Employee, error) {
	err := repo.decodeJSON(employeeJsonFile, &model.Employees)
	if err != nil {
		return []model.Employee{}, err
	}

	return model.Employees, nil
}

func (repo *employeeJsonRepository) getLastID() (int64, error) {
	employees, err := repo.List()

	var tempID int64
	for _, employee := range employees {
		if tempID < employee.ID {
			tempID = employee.ID
		}
	}
	return tempID, err
}

func (repo *employeeJsonRepository) Add(employee *model.Employee) (*model.Employee, error) {
	id, err := repo.getLastID()
	if err != nil {
		return nil, err
	}

	newEmployee := employee
	newEmployee.ID = id + 1

	model.Employees = append(model.Employees, *newEmployee)

	err = repo.encodeJSON(employeeJsonFile, &model.Employees)
	if err != nil {
		return nil, err
	}

	return newEmployee, nil
}

func (repo *employeeJsonRepository) getIndexByID(id int64) (int, error) {
	employees, err := repo.List()
	if err != nil {
		return -1, err
	}

	for i, employee := range employees {
		if id == employee.ID {
			return i, nil
		}
	}

	return -1, apperrors.New(apperrors.ErrIdNotFound)
}

func (repo *employeeJsonRepository) Detail(id int64) (*model.Employee, error) {
	employees, err := repo.List()
	if err != nil {
		return nil, err
	}

	index, err := repo.getIndexByID(id)
	if err != nil {
		return nil, err
	}

	employee := employees[index]

	return &employee, nil
}
