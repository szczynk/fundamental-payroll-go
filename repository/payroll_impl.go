package repository

import (
	"errors"
	"fundamental-payroll-go/model"
)

type payrollRepository struct{}

func NewPayrollRepository() PayrollRepository {
	return new(payrollRepository)
}

func (repo *payrollRepository) List() ([]model.Payroll, error) {
	return model.Payrolls, nil
}

func (repo *payrollRepository) getLastID() (int64, error) {
	payrolls, err := repo.List()

	var tempID int64
	for _, payroll := range payrolls {
		if tempID < payroll.ID {
			tempID = payroll.ID
		}
	}
	return tempID, err
}

func (repo *payrollRepository) Add(payroll *model.Payroll) (*model.Payroll, error) {
	id, err := repo.getLastID()
	if err != nil {
		return &model.Payroll{}, nil
	}

	newPayroll := payroll
	newPayroll.ID = id + 1

	model.Payrolls = append(model.Payrolls, *newPayroll)

	return newPayroll, nil
}

func (repo *payrollRepository) getIndexByID(id int64) (int, error) {
	payrolls, err := repo.List()
	if err != nil {
		return -1, err
	}

	for i, payroll := range payrolls {
		if id == payroll.ID {
			return i, nil
		}
	}

	return -1, errors.New("ID tidak ditemukan")
}

func (repo *payrollRepository) Detail(id int64) (*model.Payroll, error) {
	payrolls, err := repo.List()
	if err != nil {
		return &model.Payroll{}, nil
	}

	index, err := repo.getIndexByID(id)
	if err != nil {
		return &model.Payroll{}, err
	}

	payroll := payrolls[index]

	return &payroll, nil
}
